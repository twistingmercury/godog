#Godog

"Godog" is "Godog" backwards...

     _____           _               _
    |  __ \         | |             | |
    | |  \/ ___   __| | ___   __ _  | |
    | | __ / _ \ / _` |/ _ \ / _` | | |
    | |_\ \ (_) | (_| | (_) | (_| | | |
    \____/ \___/ \__,_|\___/ \__, | |_|
                              __/ /  _ 
                             |___/  |_| 
  
                    Who's a good boy!?!
    +-----------------+-------------------------------------------------------------+
    | COMMAND         | DESCRIPTION                                                 |
    +-----------------+-------------------------------------------------------------+
    | --nu            | Add a new user to Datadog.                                  |
    | --ru            | Disable an existing user in Datadog.                        |
    | --lu            | List all of the currently active and pending Datadog users. |
    | --fu            | Find a specific user by email or name.                      |
    | --lr            | Lists all of the available roles supported by Datadog.      |
    | --help | ? | -h | Shows command line help.                                    |
    +-----------------+-------------------------------------------------------------+

To add a new user:

```bash
$ godog --add-user --email="new.user@mcg.com" --name="New User" --role=s
```