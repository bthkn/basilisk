# basilisk

Basilisk helps [Cockatrice][cockatrice-github-link] download cards in your language (if they have ever been printed in that language).

## Usage

Go to `Preferences` -> `Card Sources`, click green plus button (*Add New URL*), paste basilisk link and click OK. Then drag your new link to the top of the list. 

Link example:
```
http[s]://<host>[:<port>]/card?name=!name!&code=!setcode_lower!&number=!set:num!&lang=!sflang!
```

More info at [Cockatrice Wiki][cockatrice-wiki-link]

## Search Algorithm

1. `/cards/<code>/<number>/<lang>`
  - **Found**: redirect to image
  - **Not Found**: step 2
2. `/cards/search?q=<query>&order=released`
  - **Found**: redirect to image of the first found card
  - **Not Found**: step 3
3. redirect to `/cards/<code>/<number>?format=image`

## License

Distributed under the terms of the MIT license. See [LICENSE](LICENSE) for details.

[cockatrice-github-link]: https://github.com/Cockatrice/Cockatrice
[cockatrice-wiki-link]: https://github.com/Cockatrice/Cockatrice/wiki/Custom-Picture-Download-URLs
