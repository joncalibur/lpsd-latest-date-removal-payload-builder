# Excel to JSON Converter

This utility reads data from an Excel file and generates a JSON file in the required format. It is designed for users who want to automate data transformation and JSON file generation.

## Features
- Reads an Excel file with `licensed_product_id` and `licensed_product_service_date_id` columns.
- Converts each row into a JSON object with the desired structure.
- Saves the JSON output to the same directory as the input Excel file.

## Prerequisites

### Option 1: Pre-built Binary (No Go Installation Required)
For users without Go installed:
1. Download the pre-built binary for your operating system from the [Releases](#) section of this repository.
2. Ensure the binary is executable (refer to the section [Running the Binary](#running-the-binary)).

### Option 2: Build From Source
For advanced users with Go installed:
- Requires Go 1.20 or higher.

## Installation

### Pre-built Binary
1. Download the appropriate binary for your operating system:
   - [MacOS](#)
   - [Windows](#)
   - [Linux](#)
2. Place the binary in a directory included in your system's PATH, or run it directly from the download location.

### Build From Source
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```
2. Build the project:
   ```bash
   go build -o excel-to-json
   ```
   This will create an executable file named `excel-to-json` in the project directory.

## Usage

### Input File Requirements
The input Excel file must:
- Contain at least two columns:
  - `licensed_product_id`
  - `licensed_product_service_date_id`
- Be saved in `.xlsx` format.

### Running the Binary
1. Place the input Excel file in your desired directory, e.g.:
   ```
   /Users/Tom/Personal Development/Go Lang/Payload Builders/lpsd.xlsx
   ```
2. Execute the utility:
   ```bash
   ./excel-to-json /path/to/your/excel/file.xlsx
   ```

   Example:
   ```bash
   ./excel-to-json "/Users/Tom/Personal Development/Go Lang/Payload Builders/lpsd.xlsx"
   ```

3. The output JSON file (`output.json`) will be saved in the same directory as the input Excel file.

### Example
#### Input Excel File
| licensed_product_id | licensed_product_service_date_id |
|---------------------|----------------------------------|
| 1001               | 20240101                        |
| 1002               | 20240102                        |

#### Generated JSON
```json
{
	"id": 1234,
	"products": {
		"items": [
			{
				"licensed_product_id": "1001",
				"service_dates": [
					{
						"service_date_id": "20240101",
						"action": "REMOVE"
					}
				]
			},
			{
				"licensed_product_id": "1002",
				"service_dates": [
					{
						"service_date_id": "20240102",
						"action": "REMOVE"
					}
				]
			}
		]
	}
}
```

## FAQ

### What if I don't have Go installed?
Use the pre-built binary for your operating system. No additional installation is needed.

### Can I specify a different output file location?
Currently, the output file is saved in the same directory as the input file. You can move the JSON file to another location after generation.

### How do I make the binary executable?
#### On MacOS/Linux:
1. Open a terminal and navigate to the binary's location.
2. Run:
   ```bash
   chmod +x ./excel-to-json
   ```
3. Now you can run it directly:
   ```bash
   ./excel-to-json
   ```

#### On Windows:
1. Double-click the `.exe` file, or run it via the command prompt:
   ```cmd
   excel-to-json.exe
   ```

## Contributions
Contributions are welcome! Please open an issue or create a pull request with suggested changes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
