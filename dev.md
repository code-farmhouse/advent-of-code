# Developer Notes
These notes are used to track development brainstorming.

## Roadmap
1. Core Features
2. Past Event Solutions
3. Current Year Solutions

## Core Features

### User Interface
- CLI Tool
- User Journey
  - Welcome Message
  - Request Year
  - Request Question Number
  - Encouraging Message
  - Display Requested Question
  - Options:
    1. Check your answer
    2. Show answer
    3. Print code to terminal or file
- Types
  - Message
    - MessageType string
    - MessageText string
- Functions
  - 

### Questions
- Each question set can be retrieved from https://adventofcode.com/{x}/day/{y}, where {x} is the year and {y} is the question number
- Upon running the program, a cache should be checked to see if the programs have already been retrieved and stored. If the cache is empty for the year requested, retrieve the text and cache it
- Types
  - Question
    - QuestionYear int
    - QuestionNumber int
    - QuestionText string
- Functions
  - checkCacheForQuestions(year int) \*[]Question, err
  - cacheQuestions(year int) \*[]Question, err
  - 



  
