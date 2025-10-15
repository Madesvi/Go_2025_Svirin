#!/bin/bash

echo "----- Добро пожаловать в игру \"Поле чудес!\"-----"

#	Обявляем переменную с нашим словом
HIDDEN_WORD="HOME"
#	Для проверки двух букв
#HIDDEN_WORD_2="SWEET" # С двумя буквами EE проверка
#	Считаем кол-во символов и записываем в переменную
word_length=${#HIDDEN_WORD}
#echo "Длина слова $word_length"

#	Приветсвие
echo "----- Ниже вы увидете загаданное слово под звёздочками -----"
echo "----- Подсказка: Обычно SWEET ..."
#echo "----- Подсказкза -----  Обычно с HOME"

#	Создаем переменную hidden_display и заполняем кол-вом * = кол-ву в
#	word_length циклом for
hidden_display=""
for (( i=0; i<$word_length; i++ )); do
	hidden_display+="*"
done

#	Объявляем while - пока в hidden_display есть * do
#	Первый вариант с двумя [[]]
# while [[ "$hidden_display" == *"*"* ]]; do
# -q - только возвращает код, не выводит результат
while echo "$hidden_display" | grep -q '\*'; do
	echo "----- Загаданное слово: $hidden_display"
	read -p "Напишите букву, которая как вам кажется есть в этом слове: " user_input
#	echo "-----Вы ввели $user_input"

#	Объявляем две переменные одна для обновления hidden_display
#	Вторая для условия найденных или нет букв 0 - не найдены 1 - найдеы
	new_display=""
	found=0
	for (( i=0; i<$word_length; i++ )); do
#		current_char в цикле for пробегаемся по нашему слову
#		сверяя каждый символ слова с введенным пользователем
#		$user_input
		current_char="${HIDDEN_WORD:$i:1}"
		existing_char="${hidden_display:$i:1}"

#		В данном if мы проверяем посимвольно в случае совпадения
#		Устанавливаем found=1
		if [ "$current_char" = "$user_input" ]; then
			new_display+="$current_char"
			found=1
		elif [ "$existing_char" != "*" ]; then
			new_display+="$existing_char"
		else
			new_display+="*"
		fi

#		if [ "$current_char" = "$user_input" ] || [ "${hidden_display:$i:1}" != "*" ];
#			then
#			new_display+="$current_char"
#			found=1
#			else
#			new_display+="*"
#		fi

done
	if [ $found -eq 1 ]; then
		echo "-----Вы угадали букву $user_input"
	else
		echo "Вы не угадали букву... Попробуйте ещё раз"
#Found must be 0!
	fi
#	found=0
	echo

	hidden_display="$new_display"
#	echo "На данный момент: $hidden_display"
#	echo "---"
done

echo "Поздравляем! Вы угадали слово $HIDDEN_WORD!"

#function readLetters {
#	read -p "Введите букву: " user_input
#	echo "Вы ввели букву: " $user_input
#}








