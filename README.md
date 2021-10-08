# AppGo

AppGo is an application that is intended to read a plain text log file and deliver an encoded polyline.


## Installation

To run AppGo it is necessary to download and install [Go] (https://golang.org/doc/install) and follow the steps according to your Operating System.

Once installed go run the following command to verify the installation
bash
go version

You should see the installed version.
## How to use

1. Unzip the contents of the .zip.
2. Open the command console and go to the path where the .zip was unzipped
3. Once in the path, execute "go install", this will install the necessary dependencies

bash
cd [ruta_donde_se_descomprimió]/AppGo
go install

### Execution
#### 4.1 
For Windows Operating System an AppGo.exe file will be created which is the CLI executable.

bash
./AppGo.exe

#### 4.2
For Linux systems
bash
~/go/bin/AppGo


5. The executable will read the path of the file that is entered. (must be a local file)
6. The result will be a coded line (s)

bash
Please enter your path: C:/Downloads/example.log.txt
# Outcome:
_p~iF~ps|U_ulLnnqC_mqNvxq`@