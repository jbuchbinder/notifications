# NOTIFICATIONS

Universal notifications interface to push messages to chat services. Currently supports:

* dummy (testing output)
* hipchat
* slack

## EXAMPLE

```
        n, err := NotifierByName("dummy")
	if err != nil {
		panic(err)
	}
        n.Init()
        n.Notify("Test message from new client")
```
