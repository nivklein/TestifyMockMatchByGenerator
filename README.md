# TestifyMockMatchByGenerator

When working with mocks, especially for methods that take a pointer to a struct, validating the input can be repetitive. I found myself writing mock.MatchedBy() repeatedly with similar logic for different structs. To simplify this process, I created a generator that dynamically creates mock.MatchedBy() matchers for you.

Key Features:
•	Works for primitive fields only (e.g., int, string, bool).
•	Can compare primitive fields across different struct types, as long as the field names and types match.

Check out the example below to see how it works! 😊