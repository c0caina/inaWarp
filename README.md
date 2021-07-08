# inaWarp
Warp system for [Dragonfly](https://github.com/df-mc/dragonfly "Source code")

# Commands
| Command | Sub command     | Description                                                   |
| :------ | :-------------- | :-----------------------------------------------------------  |
| /warp   | set [Name warp] | Creates a teleportation point relative to your location.      |
| /warp   | del [Name warp] | Removes teleportation point.                                  |
| /warp   | tp  [Name warp] | Teleports you to a specific point.                            |
| /warp   | list            | Displays information about all registered points in the chat. |

# Usage
### Import package
```import "github.com/c0caina/inaWarp"```

### LoadPlugin
Insert the `inaWarp.Run()` into the main file after initializing the server.