!#/bin/bash

00 10 * * * $GOPATH/bin/walk -root /mnt/c/Users/micah/Documents/WorkSpaces/Projects/command_line_projects/file_system_crawler -ext .log -size 10485760 -log file_crawler_deleted_files.log -del