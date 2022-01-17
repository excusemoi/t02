from google.transliteration import transliterate_word
suggestions = transliterate_word('wildberries', lang_code='ru')
print(suggestions)