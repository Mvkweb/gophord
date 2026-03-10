#!/bin/bash
# scripts/generate_llms_txt.sh
# Generates llms-full.txt by concatenating llms.txt and all docs/*.md files.

set -e

# Change to repo root
cd "$(dirname "$0")/.."

OUTPUT="llms-full.txt"

# Start with llms.txt contents (omitting the Optional links at the bottom if desired, or just use the whole file)
cat llms.txt > "$OUTPUT"
echo "" >> "$OUTPUT"
echo "---" >> "$OUTPUT"
echo "" >> "$OUTPUT"

echo "Appending docs/*.md files..."

for file in docs/*.md; do
    if [ -f "$file" ]; then
        echo "Appending $file"
        cat "$file" >> "$OUTPUT"
        echo -e "\n\n---\n" >> "$OUTPUT"
    fi
done

echo "Successfully generated $OUTPUT ($(wc -l < "$OUTPUT") lines)"
