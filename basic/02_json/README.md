# 02 JSON

Read from `https://api.dictionaryapi.dev/api/v2/entries/en/{word}`
using `net/http` package and it's `http.Get()` function.

Process the body using `io/ioutil` `ioutil.ReadAll` function.

Write the results to the console but only

- Word name
- Phonetics (just text)
- List of meanings

### Tip

When defining main struct for JSON unmarshal, mark it as an array of structs (`[]struct`) instead of simple `struct`

## Expected outcome

```
 $ go run main.go -word "hello"
hello
-------
Phonetics
həˈləʊ
hɛˈləʊ
-------
Meanings
Used as exclamation
Definition:
used as a greeting or to begin a phone conversation.
Example: hello there, Katie!

Used as noun
Definition: an utterance of ‘hello’; a greeting.
Example: she was getting polite nods and hellos from people

Used as verb
Definition: say or shout ‘hello’.
Example: I pressed the phone button and helloed
```

## JSON example response

```json
[
  {
    "word": "hello",
    "phonetics": [
      {
        "text": "/həˈloʊ/",
        "audio": "https://lex-audio.useremarkable.com/mp3/hello_us_1_rr.mp3"
      },
      {
        "text": "/hɛˈloʊ/",
        "audio": "https://lex-audio.useremarkable.com/mp3/hello_us_2_rr.mp3"
      }
    ],
    "meanings": [
      {
        "partOfSpeech": "exclamation",
        "definitions": [
          {
            "definition": "Used as a greeting or to begin a phone conversation.",
            "example": "hello there, Katie!"
          }
        ]
      },
      {
        "partOfSpeech": "noun",
        "definitions": [
          {
            "definition": "An utterance of “hello”; a greeting.",
            "example": "she was getting polite nods and hellos from people",
            "synonyms": [
              "greeting",
              "welcome",
              "salutation",
              "saluting",
              "hailing",
              "address",
              "hello",
              "hallo"
            ]
          }
        ]
      },
      {
        "partOfSpeech": "intransitive verb",
        "definitions": [
          {
            "definition": "Say or shout “hello”; greet someone.",
            "example": "I pressed the phone button and helloed"
          }
        ]
      }
    ]
  }
]
```
