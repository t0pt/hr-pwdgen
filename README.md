<h1>Human readable password generator</h1>

So, I had a problem with creating some passwords, that I could actually remember, without any need for the third party password-keeping software. Here I just have a list of about 300k words that can be combined with each other and result with a password. Also between the words some special characters and/or numbers will be put. Thanks to the kaggle.com/bwandowando for the words dataset!

It chages "a" to the @ randomly, "i" to the 1 and "o" to the 0 randomly, if the fancyfying (-f) is enabled. 
It capitalizes some random chars, if capitalization (-c) is enabled.
If you disable special characters (-sc) (enabled by default), it will only use random numbers as spacers between words.
Also you have the hotness scale to decide, how often it should alter the characters. Use -hottness and then specify an integer. You will have 10:n probability of editing a character (higher n results in more changes), it defaults to 5.