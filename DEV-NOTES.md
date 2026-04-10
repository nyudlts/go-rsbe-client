# DEV-NOTES

The following are some notes in case you are trying to set up a test  
environment in the following configuration:  
```
1.) Operating System    : Windows  
2.) Rails RSBE Instance : running in a Vagrant box      listening on port 3000  
3.) go-rsbe    Instance : running in a Podman container listening on port 3001 [1]
4.) test suite          : running in an Ubuntu WSL session   
```


If you're trying to do this, you may run into a situation where you cannot reach  
the Rails RSBE instance from within WSL.

Here's how to make the Rails RSBE instance accessible to WSL:
1.  In Windows, go to your user profile folder: `%USERPROFILE%`.
2.  Create or edit a file named **`.wslconfig`**.
3.  Add the following lines:
    ```ini
    [wsl2]
    networkingMode=mirrored
    ```
4.  Restart WSL (In PowerShell: `wsl --shutdown`).
5.  Now WSL and Windows share the same `localhost`, and WSL should see   
    the port Vagrant is pushing to Windows.

6. In your **Vagrantfile**, change your forward rule to this:
    ```ruby
    config.vm.network "forwarded_port", guest: 3000, host: 3000, host_ip: "0.0.0.0"
    ```
7. You must run `vagrant reload` after making this change.
  * NOTE: Windows Firewall might pop up a prompt asking if you want to    
          allow "Vagrant" or "VirtualBox" to communicate on Public/Private networks.  
          **You must allow this**, or Windows will block WSL's attempt to cross over into that port.

[1] A repo with the `go-rsbe` Podman setup can be found [here](https://github.com/nyudlts/podman-go-rsbe-client-test)