#கொற்கை - சொல் திரட்டி - A Tamil corpus builder
கொற்கை என்ற இந்த நிரல் தமிழ்ச் சொற்களை ஆய்ந்தெடுக்கும் கருவியாகும்.

ப்ளாகர், வேர்டுப்ரசு தளங்கள் உங்கள் பதிவுகளையும்,  விக்கிப்பீடியா தன் கட்டுரைகளையும், ஒரே கோப்பாக மொத்தமாகத்  தரவிறக்கிக் கொள்ள வழி கொடுத்துள்ளன. அவ்வாறு இறக்கப்பட்டுள்ள XML கோப்புகளில் இருந்து, தமிழ்ச் சொற்களை மட்டும் ஆய்ந்தெடுத்து அவற்றை அகரமுதலாகவும், பயன்படுத்தப்பட்ட எண்ணிக்கை வாரியாகவும் அடுக்கித் தரும் மென்பொருளாகும் இந்த கொற்கை என்னும் செயலி.

#உரிமம்
இந்த நிரல் கிரியேட்டிவ் காமன்சு சீரோ லைசென்சு என்னும் உரிமத்தின் கீழ் உங்களுக்கு வழங்கப்படுகின்றது. இதனை நீங்கள் வணிக நோக்கில் பயன்படுத்திக் கொள்ளலாம்.

மேலதிகத் தகவல்களுக்கு:		http://creativecommons.org/publicdomain/zero/1.0/
இவ்வுரிமத்தின் முழு சட்ட விளக்கம்:	http://creativecommons.org/publicdomain/zero/1.0/legalcode

இந்த நிரலுக்கு ஏதேனும் மாற்றங்களை நீங்கள் கொடையாக அளிக்க விரும்பினால், மேற்சொன்ன அதே உரிமத்தின் படி மட்டுமே தர வேண்டும்.

#நிறுவுதல்
இந்த நிரலின் மூலக்கோப்புகளை கம்பைல் (நிரல்மாற்றி வாயிலாக) செய்வதன் மூலம் மட்டுமே இந்த நிரலியை பயன்படுத்த முடியும்.

* http://golang.org என்ற தளத்திலிருந்து கோ நிரல்மாற்றியை நிறுவவும் (Install Go programming language compiler)
* ப்ளாகர்/வேர்டுப்ரசு/விக்கிப்பீடியா கோப்புகளைத் தரவிறக்கிக் கொள்ளவும் (Download your blogger/wordpress/wikipedia archive files)
* $> go run wikipedia-archives.go <wikipedia-archive>.xml
* $> go run blogger-xmldump-parser.go <blogger-archive>.xml
* $> go run wordpress-xmldump-parser.go <wordpress-archive>.xml
* $> ls *.tokens *.stats

*.tokens அகரமுதலாகவும் *.stats பயன்படுத்தப்பட்ட எண்ணிக்கை வாரியாகவும் அடுக்கித் தரும்

உங்களுக்கு ஒன்றுக்கும் மேற்பட்ட பதிவுகள் இருந்தால், அவற்றின் பெயர்களை ஒன்றின் பின் ஒன்றாகத் தரலாம். 

பதிவுகளில் உள்ள தமிழ்ச்சொற்கள் மட்டுமே பிரித்தெடுக்க எடுத்துக்கொண்டிருக்கிறோம். பின்னூட்டங்கள் கணக்கில் சேராது.

#XML கோப்புகளைத் தரவிறக்குதல்
###ப்ளாகர்
blogger.com -> Settings -> Other -> Export blog
பொதுவாக blog-MM-DD-YYYY.xml என்ற பெயரில் சேமிக்கப்படும்
###வேர்டுப்ரசு
TODO: Get the instructions from some wordpress user and fill here
###விக்கிப்பீடியா
http://dumps.wikimedia.org/tawiki/latest/tawiki-latest-pages-articles.xml.bz2 ஐ சேமித்துக் கொள்ளவும்.

லினக்சு பயனாளிகள் tar, file-roller மூலமும், விண்டோசு பயனாளிகள் http://gnuwin32.sourceforge.net/packages/bzip2.htm மூலமும் மேற்சொன்ன கோப்பிலிருந்து  கோப்பினை பிசைந்தெடுக்கவும் !? ;) (uncompress extract)

#கொற்கை பெயர்க்காரணம்
கொற்கை என்பது பழங்காலப் பாண்டி நாட்டுத் தலைநகரங்களுல் ஒன்றாகும். உலகின் தலைசிறந்த முத்துகள் இவ்வூரில் இருந்துதான் எடுக்கப்பட்டு உலகம் முழுமைக்கும் (உரோமாபுரி, சீனம் என்று) அனுப்பப்பட்டது. துறைமுக நகரமாகவும், பாண்டிய இளவரசனின் தலைநகரமாகவும் (மன்னனின் தலைநகரமாக மதுரையும்) விளங்கிற்று. இந்நகரம் குறித்து மேலும் அறிய திரு. சாண்டில்யன் அவர்கள் எழுதியுள்ள இராச(ஜ)முத்திரை என்ற நூலை வாசிக்கவும்.

எப்படி அந்நாளைய கொற்கை, முத்துகளை கண்டெடுக்க உதவியதோ, அதே போல, தமிழ் சொற்களாகிய முத்துகளைக் கண்டெடுக்க உதவுவதால் இந்த நிரலும் கொற்கை என்று வழங்கப்படுகிறது. (முத்துகளைப் போல தமிழ் சொற்களும் தற்காலத்தில் அரிதாகிக் கொண்டே வருவது வேறு கதை)

கொற்கை: 
* http://ta.wikipedia.org/wiki/%E0%AE%95%E0%AF%8A%E0%AE%B1%E0%AF%8D%E0%AE%95%E0%AF%88 
* https://www.goodreads.com/book/show/7715528-raja-muththirai

#நன்றி
என் ஆராய்ச்சிக்கு உறுதுணையாக தங்கள் பதிவுகளின் XML கோப்புகளைத் தந்து வழங்கிய கீழ்கண்ட நல்லுள்ளங்களுக்கு மிக மிக நன்றி.

* கார்க்கி - http://iamkarki.blogspot.com
* சொக்கன் - http://nchokkan.wordpress.com
* ப்ரியா கதிரவன் - http://priyakathiravan.blogspot.com
* இலவசக் கொத்தனார் - http://elavasam.blogspot.com
* இராகவன் - http://gragavan.blogspot.com
* சரவண கார்த்திகேயன் - http://www.writercsk.com

சோனியா கீசு (Soniya Keys) மற்றும் கைல் லெமன்சு (Kyle Lemons) ஆகியோர் கோ (Go programming language) குறித்து அறிந்து கொள்ள உதவினர்.

#Developer Notes
A hashtable is used to store the words in memory before dumping the data to the
files. We can improve the performance substantially by using a Trie instead of a
hastable. However, we have used a hashtable to keep the code short and more
easily understandable for anyone with just the programming language knowledge. 

Using the built in datastructures makes the maintenance of the code easier too.
However, if you are planning to use this in a production machine that requires 
realtime responses, you are recommended to change the hashtable to something 
that will be more fitting for your needs.

The wikipedia extractor program will use goroutines and hence all the CPUs in
your machine, thereby increasing the performance. In my thinkpad T430 with 4
CPUs and 4 GB RAM, it finishes the word extraction part in about 50 seconds.

The extractor programs for blogger and wordpress are not parallelized since the
load is usually very less and the benefits of parallellizing may not be worth 
the effort. This too may be done sometime in the future though, depending on the
need and more importantly how bored I am.
