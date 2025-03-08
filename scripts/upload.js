import fs from "fs";
import PocketBase from "pocketbase";

const pb = new PocketBase("http://127.0.0.1:8090");

await pb.collection('_superusers').authWithPassword("tykapl.breuil@gmail.com", "@9ibn6hXkXV%k^")

// Function to read JSON file synchronously
export function readJSONFile(filename) {
    try {
      const data = fs.readFileSync(filename, "utf8");
      const jsonData = JSON.parse(data);
      return jsonData;
    } catch (error) {
      console.error("Error reading JSON file:", error.message);
      return undefined;
    }
  }

const words = readJSONFile('./word_list.json');
const wordIds = {};
for (const wordIndex in words) {
    const word = words[wordIndex];
    const dbData = {
        "word": word,
    };
    try {
        const result = await pb.collection('words').create(dbData);
        wordIds[word] = result.id;
    } catch (error) {
        console.error('Error creating words: ' + word);
        throw error;
    }
}

const candidates = readJSONFile('./candidates.json');
for (const candidateIndex in candidates) {
    const candidate = candidates[candidateIndex];
    const dbData = {
        "available_letters": candidate[1],
        "word": wordIds[candidate[0]],
    };
    try {
        await pb.collection('candidates').create(dbData);
    } catch (error) {
        console.error('Error creating candidate: ' + candidate.name);
        throw error;
    }
}
