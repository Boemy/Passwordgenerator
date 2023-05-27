# Password Generator

This project was an assignment from school with the objective of creating a tool that generates a good password. The program allows users to specify the length of the password and what types of characters they want to include, such as numbers, logograms, punctuation, quotations, dashes/slashes, math symbols, and brackets.

After generating a password, the program checks if the password already exists in a MySQL database. If it does, a new password is generated and checked until a password is created that is not in the database. Once a new password is created, it is added to the database.

## Flags

You can customize the password generation by using the following flags:

    L: Length of the password (Default value: 16)
    n: Should numbers be used? (Default value: False)
    l: Should logograms be used? (Default value: False)
    p: Should punctuation be used? (Default value: False)
    q: Should quotations be used? (Default value: False)
    ds: Should dashes & slashes be used? (Default value: False)
    m: Should math symbols be used? (Default value: False)
    b: Should brackets be used? (Default value: False)

You can set the flags to true or false. For example, to generate a password with length 32, using punctuation and math symbols, you can run:

    go run main.go -L=32 -p=true -m=true

## Dependencies

This project uses the following dependency:

    github.com/go-sql-driver/mysql

To install the dependency, you can run the following command:

    go get github.com/go-sql-driver/mysql

## Database Configuration

The MySQL database configuration is stored in the Config.json file. You can edit this file to specify your database settings.
```
{
    "db_login": [
        {
            "db_name": "password_generator",
            "db_user": "your_username",
            "db_pass": "your_password"
        }
    ]
}
```

How to use

1. Clone the repository: 
```
git clone https://github.com/[username]/password-generator.git
```

2. Navigate to the project directory:
```
cd passwordgenerator
```
3. Install dependencies:
```
go get github.com/go-sql-driver/mysql
```
4. Edit the Config.json file to specify your database settings.

5. Run the program with your desired flags:
```
go run main.go -L=[password length] -n=[true/false] -l=[true/false] -p=[true/false] -q=[true/false] -ds=[true/false] -m=[true/false] -b=[true/false]
```

## Contributors

This project was created as a school assignment by Boemy. If you would like to contribute to this project, feel free to fork the repository and submit a pull request.
