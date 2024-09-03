## Strategy Employed
My initial thought was that there should be some api in the gorilla websockets package that would allow me to create a connection with another socket after having the connection with another socket. 
So I was going to go about it by specifying the ip that I wanted the message to be sent to and then adding the message from the client side and then on the server side I parse it, get the ip address I'm supposed to sending the message to and then use the api that 
I thought would exist to create another connection with another socket and then send the message. As it turns out, it is not possible for a websocket server to have more than one client on the same connection, which kind of makes sense as I think about it.

My other try was to close the initial connection(not that smart) and then create a TCP connection with the intended ip address we want to send the message to. In my head this was supposed to work but there is a problem with the way I think I'm supposed to test it.
Maybe If I test it right I could actually get it work. But then the other client can't mantain a websocket connection with the server for this to work. Which kind of defeats the purpose of using websockets for real time communication so maybe it's not the same thing.
