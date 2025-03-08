# Load and process French words dataset
import pandas as pd

def preprocess_word(word: str) -> str:
    translationTable = str.maketrans("ïëéàèùâêîôûç", "ieeaeuaeiouc")
    word = ''.join([(c.lower()).translate(translationTable) for c in word if '-' not in c])
    if not word.isalpha():
        raise ValueError(f"Word {word} is not alphabetic")
    return word

french_words = pd.read_csv("./liste_francais.txt")
print(french_words.head())
unique_words = french_words['word'].str.lower().drop_duplicates()
words_list = [preprocess_word(word) for word in unique_words if len(word) >= 4]

# Precompute masks for each word with exactly 7 unique letters
masks = []
candidates = []
for i, word in enumerate(words_list):
    if len(set(word)) == 7:
        candidates.append(i)
    mask = 0
    unique_letters = set(word)
    for c in unique_letters:
        pos = ord(c) - ord('a')
        if pos < 0:
            print(c)
        mask |= (1 << pos)
    masks.append(mask)
    
    if i % 10000 == 0:
        print(f"Processed {i} words")

result = []
min_count = len(words_list)
min_index = -1
min_list = []
for index, i in enumerate(candidates):
    current_mask = masks[i]
    count = 0
    matches = []
    for j in range(len(words_list)):
        if i == j:
            continue
        other_mask = masks[j]
        if (other_mask & current_mask) == other_mask:
            count += 1
            matches += [words_list[j]]
            
    if count < min_count:
        min_count = count
        min_index = i
        print(f"New min count: {words_list[i]}: {count}")
        min_list = matches
    if index % 100 == 0:
        print(f"Processed {index}/{len(candidates)} words")

    available_letters = []
    # Check there is a list of 12 words with a common letter
    for letter in set(words_list[i]):
        count_letter = 0
        for word in matches:
            if letter in word:
                count_letter += 1
        if count_letter >= 12:
            available_letters.append(letter)
            continue
    if len(available_letters) < 1:
        print(f"Skipping {words_list[i]}: {count}, {matches}")
        continue
            
    result.append((words_list[i], available_letters))

# Dump the result as json
import json
with open("candidates.json", "w") as f:
    json.dump(result, f)
with open("word_list.json", "w") as f:
    json.dump(words_list, f)
