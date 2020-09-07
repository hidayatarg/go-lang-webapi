## Glide

- Install Glide ``brew install glide`` for mac and for other operating system 
https://github.com/Masterminds/glide

- Install go pgsql package ``glide get github.com/go-pg/pg``

- To build go ``go build`` and then run via ``./go-postgresql``
#### NOTE
For mac permissions need to do the following incase you get homebrew problems
fix owner of files and folders recursively

``sudo chown -vR $(whoami) /usr/local /opt/homebrew-cask /Library/Caches/Homebrew``

fix read/write permission of files and folders recursively

``chmod -vR ug+rw /usr/local /opt/homebrew-cask /Library/Caches/Homebrew``

fix execute permission of folders recursively

``find /usr/local /opt/homebrew-cask /Library/Caches/Homebrew -type d -exec chmod -v ug+x {} +``