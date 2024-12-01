Asset Preview

A project to help preview the assets for a Flutter/Dart project in one go.
Since Flutter or IDEs like Android Studio or Visual Studio Code do not provide a way to preview the assets, this project aims to provide a way to preview the assets.

It is a simple project built using Golang. It opens a browser window and displays the assets in a list with the file name and the image.

## How to use
1. Download the binary from the [releases page](https://github.com/nicks101/asset_preview/releases).
2. Run the binary in the root directory of the project. Or in the directory where the assets are located.

```bash
./asset-preview
```

3. Provide the path to the assets directory. If path is not provided, the program will look for the assets in the current directory.

```bash
./asset-preview -path="/path/to/assets"
```

4. An `preview_assets.html` file will be generated. It will open in the default browser.  
The file will be generated in the current directory.

**Note: Do not push this file to the version control.**


## How to build
1. Clone the repository.
2. Run the following command to build the binary.

```bash
go build
```

3. Run the binary in the root directory of the project. Or in the directory where the assets are located.

```bash
./asset-preview
```

## Contributing
1. Fork the repository.
2. Create a new branch.
3. Make changes.
4. Create a pull request.

All contributions are welcome! 🤗 

## TODO
- [ ] Test on Windows, Linus OS
- [ ] Add command line arguments to take in the file types to preview
