Robomongo is now Robo 3T. Following are the updated steps:

Download the tar file from robomongo site. The current file is robo3t-1.1.1-linux-x86_64-c93c6b0.tar.gz, but yours could be different.

Open up the terminal, switch to download directory and run the following commands:

$ tar -xvzf robo3t-1.1.1-linux-x86_64-c93c6b0.tar.gz
$ sudo mkdir /usr/local/bin/robomongo
$ sudo mv  robo3t-1.1.1-linux-x86_64-c93c6b0/* /usr/local/bin/robomongo
$ cd /usr/local/bin/robomongo/bin
$ sudo chmod +x robo3t 
$ sudo gedit ~/.bashrc
Add the following line to the end of .bashrc file:

alias robomongo='/usr/local/bin/robomongo/bin/robo3t'

Save and close the file. Now reload it using the following command:

$ source ~/.bashrc
Then you can run robomongo from your terminal and it will work:

$ robomongo