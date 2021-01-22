# Reading-Arguments-from-CLI
1) os.Args:--> It is of type slice of strings that can be used for elementary command-line argument processing, 
            which is reading arguments that are given on the command-line when the program is started.
            
2) Space (" ");---> command line condidered " " as one argument.

3) strings.Join(os.Args[1:] , "&&")--> Take arguments from user in CLI and after every arguments it adds  '&&'.
