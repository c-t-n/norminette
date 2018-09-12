# Contribution Guide

:tada: Welcome, curious one and thanks for taking the time to contribute !  :+1: :pray:

*This following contribution guide is mostly inspired from the Atom Contribution Guidelines*

Here lies a set of guidelines for contributing to Norminette.
These are mostly guidelines, not rules.
We recommend you to use your best judgement and logic, and PR are welcome if you want to propose changes on this document.

## Code of Conduct

This project and everyone participating in it is governed by this [Code of Conduct](https://github.com/c-t-n/norminette/wiki/Code-of-conduct).
By participating, you are expected to uphold this code.

Please report unacceptable behavior to a core contributor.

## You just want to ask a question ?

At the moment, we use the issue board as main way of communication for everyone.
Remember to scroll down the issue board before asking, but questions are still welcome.
We will edit an official FAQ with the most repeated issues :)

If you are a 42 Paris Student, you can reach us on Slack in `#norminette_opensource`.
Remember: Even if Slack is a chat service, sometimes it takes several hours for community members to respond -- please be patient :)

## Contributing

If you don't know yet, this project is written in Go and it should use all the standard tools used by the Go Community.

Anybody can contribute, even if this project is designed by and for the 42 Community.

You can contribute in many ways:
* Submiting a bug on our Issueboard
* Submitting an idea
* Discuss about a new feature on an issue
* Resolving an Issue by submitting your Pull request

## Usefull Resources

* [A tour of Go](https://tour.golang.org/)
* [Effective Go](https://golang.org/doc/effective_go.html)
* The 42 Norm (See the corresponding Wiki)

## Submiting a bug

Bugs are tracked as Github Issues. Bugs are defined by a wrong behaviour of the Norminette like a Norm Error not spotted or falsely spotted,
an unexpected behaviour that crashes the Norminette or missing tests.

Bugs will be tagged as :bug:`Bug` in the Issue Board

Explain the problem and include additional detials to help maintainers reproduce the problem:

* Be clear, concise and polite in your issue
* Use a clear a descriptive title and add the Rule number implied by the bug
* Provide a snippet of the tested C file that provokes the bug
* Provide the correct expected behaviour
* If the bug concern the core program and not the linting, describe the steps to reproduce the bug
* If you want, add a PR according to the development guidelines to your issue

Our Core Team will moderate your issues and validate your point (or not.............. because we can)

## Submiting an idea or an enhancement

You may want to enhance the Norminette with a new library, a better code coverage on some source code or improve the linting for a specific (set of) rule(s).
Your help is more than welcome, and your expertise is more than needed.

Improvements will be tagged as :racing_car:`Improvement`

Explain the problem and include additional detials to help maintainers to understand your solution:

* Be clear, concise and polite in your issue
* Use a clear a descriptive title
* Provide examples and/or simple benchmarks of your solution 
* If the bug concern the core program and not the linting, describe the steps to reproduce the bug
* Explain why your solution is important to consider, with humbleness and modesty

Our Core Team will moderate your issues and validate your point (or not.............. because we can)

## Submiting a Norm Rule implementation - Development guidelines

Rules need to be added to Norminette, and rules need to cover all possible usecases, even tricky ones.
You have to cover a maximum of use cases, but feel free to ask for more tricky use cases.

Rule implementation will be tagged with :100:`Rule Implementation`

If you implement a Norm rule:
* Your rule is in a specific file with the rule number and a distinguishable filename (Example: `01_01-Code_indentation.go`)
* A flag is added to the Cli in the Option Struct (cmd.go) to disable your rule.
* This flag has a distinguishable name (Example: BypassCodeIndentation) and a correct usage
* It's also added in the `init()` function block of cmd.go


## Submit your Pull Request

If you want to contribute in our code base, we expect that:

* Every feature is tested, and tests are provided for at least all use cases
* Tests are written and updated first, code goes after
* Your code should be maintainable by everyone, so make it simple and clear...
* ... and if you don't know how to make it clear, stick to our code design
* A clear and efficient documentation is provided with your code.
* Your contribution compiles with the rest of the code base (at least for Mac OS X)
* Your contribution will have the approbation of at least one of the core team member before merging
* Your contribution passes all the unit tests before merging



