# 🛠️ Go Utilities: Word Frequency & Palindrome Checker

A simple, interactive command-line tool written in Go that lets you:

- Count the frequency of words in a sentence
- Check if a word or phrase is a palindrome

---

## 🚀 Features

- **Word Frequency Counter:**
  - Ignores punctuation and case
  - Displays a table of each word and its count
- **Palindrome Checker:**
  - Ignores case and non-letter characters
  - Colorful output for results
- **User-friendly terminal formatting**
- **Input validation** for numbers and strings

---

## 🖥️ Usage

1. **Clone or download this project.**
2. **Run the program:**

```sh
go run main.go
```

3. **Follow the menu prompts:**

```
==============================
1. Count frequency of words in a sentence
2. Check if a word is a palindrome
3. Exit program
==============================
Enter choice: 1
Enter string: Go is fun. Go is powerful!

Result:
---------------------------
go              : 2
is              : 2
fun             : 1
powerful        : 1
---------------------------
```

---

## 🧩 Code Structure

- `freq_count(sentence string) map[string]int` — Counts word frequencies
- `is_palindrome(word string) string` — Checks if input is a palindrome
- `getStringInput`, `getIntInput` — Validated user input
- `prompt()` — Main menu loop

---

## 🛠 Built With

- Go 1.21+
- Standard libraries: `fmt`, `bufio`, `os`, `strings`, `strconv`, `unicode`

---

## 📄 License

Free to use and modify for learning or personal projects.

---

Feel free to suggest improvements or request new features!
