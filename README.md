# ASCII Art Generator

ASCII Art Generator is a program written in Go that converts input strings into graphic representations using ASCII characters.

## Features

- Converts strings into ASCII art
- Supports numbers, letters, spaces, special characters, and newline characters ('\n')
- Utilizes specific graphical templates for ASCII representation

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/khahussein/ascii-art
    ```

2. Navigate to the project directory:

    ```bash
    cd ascii-art
    ```

3. Build the project:

    ```bash
    go build .
    ```

## Usage

To generate ASCII art for a string, run the following command:

```bash
./ascii-art "input string"
```

Example:

```bash
./ascii-art "Hello\n" | cat -e
```

Output:

```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$                                                  
```

## File Formats

- `standard.txt`: Standard ASCII character set
- `shadow.txt`: Shadowed ASCII character set
- `thinkertoy.txt`: ASCII character set with thinkertoy style

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by  [Kherld Hussein](https://learn.zone01kisumu.ke/git/khahussein)