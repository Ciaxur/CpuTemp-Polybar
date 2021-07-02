# CpuTemp-Polybar
> Displays the current CPU's Temperature Information for Polybar

## Build 🔨
Use the `go build` Command to compile the application.
``` bash
go build -o cpu-temp ./src
```

## Arguments
All arguments are optional, though it depends on personal use. If you have an **AMD CPU** then you'll need to pass in the **CPU Type** argument.
- `-c`: CPU Type (*default is intel*)
  - Options: `amd`, `intel`
- `-i`: CPU Icon Hex Code Color (*default is #FF*)
- `-s`: CPU Temperature String Hex Code Color (*default is #FF*)


## Usage 🚀
There are two optional Arguments to Change the Color Output of the Icon and String
``` bash
# Defaults Color to White
cpu-temp

# Arguments passed are for colors (Icon & Temperature String)
cpu-temp -i "#FF" -s "#22"
```

## License
This project is licensed under the [MIT License](LICENSE).