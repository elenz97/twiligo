# twiligo
Twilio Call Go Integration
# Local installation
Set the environment variables specified in `.env.example`:
```
export TWILIO_ACCOUNT_SID=<SID>
export TWILIO_ACCOUNT_TOKEN=<TOKEN>
export TWILIO_ENDPOINT_URL=<URL>   // (http|https) endpoint of the server running the go binary
export TWILIO_MP3_URL=<URL*.mp3$>  // e.g. `http://www.evidenceaudio.com/wp-content/uploads/2014/10/monsterslap.mp3`
export TWILIO_RECEIVER=012345678   // number you wish to receive calls on. Without '+' but with specific country prefix
export TWILIO_SENDER=112345678     // number you wish to make calls from. Without '+', but with specific country prefix
```

Build binary:
```
CGO_ENABLED=0 go build -installsuffix 'static' -o /build/twiligo .
```

Run server:
```
./twiligo
```
Navigate to `http://<your_machine>:2255/makecall` to make your first call.

Navigate to `http://<your_machine>:2255/twiml` to see the XML that will be passed to your TwiML integration (the .mp3 file it will be using).

_Alternatively_, you can use `curl`:

`curl http://<your_machine>:2255/makecall`

__Beware__, as this requires you to open your machine's IP to the public. Otherwise, Twilio will not be able to reach the `/twiml` endpoint.

An unsuccessful operation (most likely in the case of Twilio not being able to make a connection to your endpoint) will return an error for your request.
Please refer to logs, as they will tell you how the request to the Twilio API was specified.

For a list of calling code prefixes of all countries, see [Wikipedia](https://en.wikipedia.org/wiki/List_of_mobile_telephone_prefixes_by_country).

# Using Docker

Image size: 13.6MB

To quickly run the latest image, set your environment variables via `.env.example` and rename the file to `.env`, or specify them via `-e|--env`.

Start the container:
```
docker run --rm --env-file .env -dt -p 127.0.0.1:2255:2255 elenz/twiligo:latest
```

# Twilio configuration
- Log in to your [Twilio Console](https://www.twilio.com/console/).
    (If you do not have an account yet, [try it](https://www.twilio.com/try-twilio) for free!)
- Navigate to your [Twilio Phone Numbers](https://www.twilio.com/console/phone-numbers) and select the desired number
- In the 'configure' tab of your desired phone number, configure the `incoming call` webhook to use `http://<TWILIO_ENDPOINT_URL>:2255/twiml` `HTTP GET`

_(TWILIO_ENDPOINT_URL equals the accessible hostname of your machine running twiligo)_


That's it! By hitting `http://<your_machine>/2255/makecall` you should receive a call on the set number of `TWILIO_RECEIVER` using your custom MP3 file.

## inspired by

#### Ricky Robinett
https://www.twilio.com/blog/author/rrobinett

see the following for code references:

https://www.twilio.com/blog/2014/10/making-and-receiving-phone-calls-with-golang.html

https://gist.github.com/rickyrobinett/25e6b56ef2b4d709d124

#### Swatto
https://github.com/Swatto/promtotwilio
