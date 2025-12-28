# Chapter 1 — Go Basics

## Overview
This chapter introduces the very basics of the Go programming language.
The goal is to understand how a Go program is structured and how to write
simple programs using variables and basic data types.

---

## Learning Objectives
By the end of this chapter, you should be able to:
- Understand the structure of a Go program
- Write and run a simple Go application
- Declare variables and constants
- Use basic data types
- Print output to the terminal

---

## Topics Covered

### 1. Go Program Structure
- `package main`
- `func main()`
- `import`
- Entry point of a Go program

### 2. Printing Output
- `fmt.Println`
- `fmt.Printf`

### 3. Variables
- `var` keyword
- Short variable declaration (`:=`)
- Naming conventions

### 4. Constants
- `const`
- Difference between variable and constant

### 5. Basic Data Types
- `string`
- `int`
- `float64`
- `bool`

---

## Folder Structure

```

Chapter 1 - Go Basics/
├── README.md
├── examples/
│   ├── hello.go
│   ├── variables.go
│   └── constants.go
├── exercises/
│   └── basic_exercises.go
└── mini-project/
    └── simple-profile.go

```

---

## How to Run Examples

Navigate to the example file directory and run:

```bash
go run hello.go
```

---

## Mini Project

### Simple Profile CLI

Create a simple CLI program that:

* Stores your name, age, and profession in variables
* Prints them in a formatted way to the terminal

Example output:

```
Name: Arifian
Age: 24
Profession: Software Engineer
```

---

## Rules for This Chapter

* Do not use external libraries
* Do not use advanced features
* Focus on understanding, not optimization

---

## Self-Check Questions

Before moving to the next chapter, make sure you can answer:

* What does `package main` mean?
* Why is `main()` important?
* What is the difference between `var` and `:=`?
* When should I use `const`?

---

## Notes

Write down anything confusing you encounter.
Confusion is expected and part of the learning process.
