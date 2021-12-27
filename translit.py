from google.transliteration import transliterate_word
suggestions = transliterate_word('Wildberries', lang_code='ru')
print(suggestions)