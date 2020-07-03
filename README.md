# nodewatch
sync markdown notes by git

1. ### support options

   - config     Configure the watch directory (default : the current directory)

     - ```shell
       eg: nodesWatch config
       ```

   - install     install service and the app run as service

   - uninstall  uninstall service

   - stop  stop the service

   - start start the service

2. ### how to use

1. make your markdown dir as contol of git

2. store your git username and password

3. clone the project to your local dir

4. cd to the project dir and run go build

5. copy exe file to your markdown dir

6. run  install command

   1. ```
      eg: nodeswatch.exe install
      ```

7. edit server nodesWatch  let the service run by current user

8. it's over