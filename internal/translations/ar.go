package translations

func initArabicTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "يتطلب هذا الموقع تشغيل javascript للعمل بشكل صحيح")

	translation.put("start-the-game", "استعد!")
	translation.put("force-start", "تشغيل اجباري")
	translation.put("force-restart", "إعادة تشغيل اجبارية")
	translation.put("game-not-started-title", "لم تبدأ اللعبة بعد")
	translation.put("waiting-for-host-to-start", "انتظر حتى يبدأ المضيف اللعبة")

	translation.put("now-spectating-title", "أنت الآن مشاهد للعبة")
	translation.put("now-spectating-text", "يمكن مغادرة وضع المشاهدة بالضغط على زر العين.")
	translation.put("now-participating-title", "أنت الآن مشارك في اللعبة")
	translation.put("now-participating-text", "بمكنك الدخول في وضع المشاهدة بالضعط على أيقونة العين في الأعلى")

	translation.put("spectation-requested-title", "تم طلب وضع المشاهدة")
	translation.put("spectation-requested-text", "ستصبح مشاهداً بعد هذه الجولة")
	translation.put("participation-requested-title", "نم طلب المشاركة")
	translation.put("participation-requested-text", "ستشارك في الجولة القادمة")

	translation.put("spectation-request-cancelled-title", "تم إلغاء طلب المشاهدة")
	translation.put("spectation-request-cancelled-text", "تم إلغاء طلب المشاهدة، ستستمر كلاعب.")
	translation.put("participation-request-cancelled-title", "ألغي طلب المشاركة في اللعبة")
	translation.put("participation-request-cancelled-text", "ألغي طلب مشاركتك. ستبقى كمشاهد.")

	translation.put("round", "جولة")
	translation.put("toggle-soundeffects", "تثبيت التأثيرات الصوتية")
	translation.put("toggle-pen-pressure", "تثبيت ضغط القلم")
	translation.put("change-your-name", "الاسم")
	translation.put("randomize", "عشوائية")
	translation.put("apply", "تطبيق")
	translation.put("save", "حفظ")
	translation.put("toggle-fullscreen", "وضع الشاشة الكاملة")
	translation.put("toggle-spectate", "الدخول في وضع المشاهد")
	translation.put("show-help", "المساعدة")
	translation.put("votekick-a-player", "صوت لطرد اللاعب")

	translation.put("last-turn", "(آخر جولة:%s)")

	translation.put("drawer-kicked", "بما أنه تم طرد اللاعب الذي يرسم، لن يحصل أحد على نقاط هذه الكلمة")
	translation.put("self-kicked", "تم طردك")
	translation.put("kick-vote", "(%s/%s) لاعبون صوتو لطرد %s.")
	translation.put("player-kicked", "تم طرد اللاعب")
	translation.put("owner-change", "%sهو مدير اللعبة الجديد")

	translation.put("change-lobby-settings-tooltip", "تغيير إعدادات اللعبة")
	translation.put("change-lobby-settings-title", "إعدادات اللعبة")
	translation.put("lobby-settings-changed", "تم تغيير إعدادات اللعبة")
	translation.put("advanced-settings", "اعدادات متقدمة")
	translation.put("chill", "للاستمتاع")
	translation.put("competitive", "للمنافسة")
	translation.put("chill-alt", "قد تفوز إن أجبت بسرعة، لكن المهم هو الاستمتاع، نتيجة الإجابة الصحيحة لا تتناقص بسرعة، خذ وقتك!")
	translation.put("competitive-alt", "كلما أجبت بسرعة أكبر، كانت نتيجتك أكبر، تتناقص نتيجة الإجابة مع الوقت بسرعة.")
	translation.put("score-calculation", "النتيجة")
	translation.put("word-language", "اللغة")
	translation.put("drawing-time-setting", "وقت الرسم")
	translation.put("rounds-setting", "الجولات")
	translation.put("max-players-setting", "اكبر عدد من اللاعبين")
	translation.put("public-lobby-setting", "لعبة عامة")
	translation.put("custom-words", "الكلمات المخصصة")
	translation.put("custom-words-info", "ادخل الكلمات الإضافية وضع بينها فاصلة")
	translation.put("custom-words-per-turn-setting", "الكلمات المخصصة في الدور")
	translation.put("players-per-ip-limit-setting", "عدد اللاعبين المسموح من عنوان واحد")
	translation.put("save-settings", "احفظ الاعدادات")
	translation.put("input-contains-invalid-data", "بعض المعلومات غير مقبولة:")
	translation.put("please-fix-invalid-input", "صحح المدخلات الخاطئة وحاول ثانية")
	translation.put("create-lobby", "إنشاء لعبة")
	translation.put("create-public-lobby", "إنشاء لعبة عامة")
	translation.put("create-private-lobby", "إنشاء لعبة خاصة")

	translation.put("players", "لاعبون")
	translation.put("refresh", "تحديث")
	translation.put("join-lobby", "ادخل للعبة")

	translation.put("message-input-placeholder", "اكتب تعليقاتك وتخميناتك هنا")

	translation.put("word-choice-warning", "الكلمة في حال لم تختر ضمن الوقت")
	translation.put("choose-a-word", "اختر كلمة")
	translation.put("waiting-for-word-selection", "في انتظار اختيار الكلمة")
	// This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "يختار كلمة")

	translation.put("close-guess", "'%s' قريبة جدا من الكلمة الصحيحة")
	translation.put("correct-guess", "لقد خمنت الكلمة الصحيحة")
	translation.put("correct-guess-other-player", "'%s' خمن الكلمة الصحيحة")
	translation.put("round-over", "انتهى الوقت، لم يتم اختيار كلمة")
	translation.put("round-over-no-word", "انتهى الوقت، كانت الكلمة: '%s'.")
	translation.put("game-over-win", "مبروك! لقد فزت.")
	translation.put("game-over-tie", "إنه تعادل!")
	translation.put("game-over", "ترتيبك%s. بمجموع %s نقاط")

	translation.put("change-active-color", "غير اللون المستخدم")
	translation.put("use-pencil", "استخدم القلم")
	translation.put("use-eraser", "استخدم الممحاة")
	translation.put("use-fill-bucket", "استخدم دلو الدهان (يملأ المكان المحدد)")
	translation.put("change-pencil-size-to", "تغيير حجم القلم او الممحاة إلى%s")
	translation.put("clear-canvas", "تنظيف اللوحة")
	translation.put("undo", "التراجع عن اخر تغيير ( لا تعمل بعد \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "تم فقدان الاتصال")
	translation.put("connection-lost-text", "تتم الآن محاولة الاتصال"+
		" ...\n\nMake sure your internet connection works.\nIf the "+
		"problem persists, contact the webmaster.")
	translation.put("error-connecting", "خطأ أثناء الاتصال")
	translation.put("error-connecting-text",
		"Scribble.rs couldn't establish a socket connection.\n\nWhile your internet "+
			"connection seems to be working, either the\nserver or your firewall hasn't "+
			"been configured correctly.\n\nTo retry, reload the page.")

	translation.put("message-too-long", "رسالتك طويلة للغاية")

	// Help dialog
	translation.put("controls", "التحكم")
	translation.put("pencil", "القلم")
	translation.put("eraser", "الممحاة")
	translation.put("fill-bucket", "دلو الدهان")
	translation.put("switch-tools-intro", "يمكنك التبديل بين الأدوات باستخدام الاختصارات")
	translation.put("switch-pencil-sizes", "يمكنك تغيير حجم القلم باستخدام المفاتيح%s إلى %s.")

	// Generic words
	// "close" as in "closing the window"
	translation.put("close", "إغلاق")
	translation.put("no", "لا")
	translation.put("yes", "نعم")
	translation.put("system", "النظام")

	translation.put("source-code", "الكود المصدري")
	translation.put("help", "مساعدة")
	translation.put("submit-feedback", "اعطنا رأيك")
	translation.put("stats", "الحالة")

	RegisterTranslation("ar", translation)
	RegisterTranslation("ar-ar", translation)

	return translation
}
