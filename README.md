# Lets Go Chat
Currently this package allows you to hash any string with **SHA256** algorithm, and compare a row string with already hashed value.

# Info

You can use password hasher as a stand alone program.

## How to build executable file

In project folder open terminal and type `go build`.
In the same folder **lets-go-chat.exe** executable file will appear.

## How to run executable file

If you just open exe file you will see message "Error! No arguments passed", so you need to provide some arguments.
You can pass up to two parameters. As first parameter you can provide any not blank string:
`.\lets-go-chat.exe VerySecurePassword`
In console you will see hash value for provided string: 
`66b532b8105c587bc6dd4a3099d0d92ebb5121cfcca6b1612735b3acc529215d`

Or, if you already have a hashed value and want to check some password, you can pass the `password` as first parameter and `hash` as a second one:
`.\lets-go-chat.exe <password> <hash>`
In console output you will see message like this: `Is password correct: true/false`

## License

[BSD](LICENSE)