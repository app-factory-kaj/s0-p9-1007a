Feature: Greeting

  @story-1
  Rule: A request naming someone receives a greeting addressed to them

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Caller sends GET /hello?name=Alice
      Then the response is a JSON greeting addressed to "Alice"

  @story-2
  Rule: A request with no name still receives a default greeting

    Scenario: Greeting with no name supplied
      Given the greeter service is running
      When an API Caller sends GET /hello with no name parameter
      Then the response is a JSON greeting addressed to "World"

    Scenario: Greeting with an empty name
      Given the greeter service is running
      When an API Caller sends GET /hello?name=
      Then the response is a JSON greeting addressed to "World"
