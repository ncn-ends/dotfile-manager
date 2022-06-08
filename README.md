
# dotfile-manager

A tool to make it easy to manage dotfiles on your system.
Currently, it only manages configuration files in csgo. Additional configuration and development is required to manage all dotfiles.



## Installation and Usage

1.  Clone the repo
2.  Create an instance of the repo in your local system in any directory you prefer.
3.  Update the cfg files via GitHub or locally to your liking.
- Note: If you have many cfg files, you can organize them under under as many sub directories as you like within the local `cfg` directory. At runtime, they will all be moved to cfg directory without being nested in subdirectories.
4.  Open your terminal in the directory. You should see `program.go` when you entry `ls`.
5.  Type `./run.sh`

Repeat steps 3-5 whenever you want to update your cfg files.
- You can edit cfg files directly from GitHub and run the tool to update your cfg files.

## Authors

- [@ncn-ends](https://www.github.com/ncn-ends)


## License

[MIT](https://choosealicense.com/licenses/mit/)