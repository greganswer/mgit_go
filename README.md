# mgit_go


- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Code of Conduct](#code-of-conduct)

[mGit](https://github.com/greganswer/mgit) rewritten in Golang

## Requirements

## Installation

## Usage

```bash
# Assuming you have a Jira ticket with the following info:
#       ID: JIR-642
#       Title: Update README file
mgit branch JIR-642
#=> jir-642-update-readme-file

mgit commit
#=> Message: JIR-642: Update Readme File

# If you make changes and run it again, it will make another commit with the
# same name
mgit commit
#=> Message: JIR-642: Update Readme File

mgit pull-request
#=> Title: JIR-642: Update Readme File

mgit tag
#=> Tag Title: my_project v2.1.7
```

## Development

## Testing

## Contributing

Bug reports and pull requests are welcome on GitHub. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to
the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The project is available as open source under the terms of the
[MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in this project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](/CODE_OF_CONDUCT.md).
