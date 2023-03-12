Napisz program w Go, który sprawdzi/wygeneruje elementy ciągu Collatza (powiązanego z tzw. Hipotezą Collatza - dotychczas nieudowodnionym twierdzeniem w matematyce). Elementy takiego ciągu to liczby całkowite, które oblicza się w sposób następujący:


	
Sprawdź czy dana liczba jest parzysta.
	
Jeżeli liczba jest parzysta, podziel ją przez dwa i zapamiętaj.
	
Jeżeli liczba nie jest parzysta, pomnóż ją przez trzy i dodaj jeden, i zapamiętaj.
	
Wróć do punktu 1.

Liczby generowane za pomocą tego algorytmu mają bardzo ciekawą cechę: ZAWSZE kończy się na dojściu do cyklu 4->2->1, i tak w kółko. W programie, możesz zbadać ile iteracji trzeba wykonać zanim liczba, którą sprawdzasz wpadnie w taki cykl. Oto zadania do napisania, i zbadania. Przy okazji - zachęcam do zastanowienia, dlaczego to tak się dzieje. Dlaczego? Tego obecnie NIKT na świecie nie wie. Może Ty będziesz osobą, która wpadnie na rozwiązanie!
Zadanie 1:

sprawdź, jaka liczba z zakresu 10-100 daje najdłuższy ciąg Collatza.
	
możesz napisać swój program od razu, aby sprawdzał zakresy, np. 10-100, albo 1000-2000, itp. - wtedy pierwszy punkt też będziesz mieć rozwiązany

Zadanie 2:

zobacz, czy w przedziałach co 1000 największe długości cyklu mają podobne liczby, czy zupełnie "losowe" - to znaczy: sprawdź przedział 1000-2000, 2000-3000, i tak dalej, i zobacz, czy dane liczby generujące najdłuższy cykl są do siebie jakoś podobne. Może coś je łączy? A może nic?

Zadanie 3:
	
wygeneruj długości ciągów Collatza dla liczb w zakresie 100000. Zrób wykres zależności długości takiego ciągu od liczby, która go generuje (czyli: na osi X masz liczby, a na osi Y długości ciągów, które te liczby generują). Wykresy można zrobić za pomocą znakomitego programu Gnuplot (polecenie: gnuplot). Jest to profesjonalny program do tworzenia wykresów naukowych - bardzo polecam żeby go sobie ogarnąć, przyda się Wam na studiach smile
	https://www.gnuplot.info/

Zadanie 4:

wylicz długości ciągów Collatza i policz ich statystykę, w sensie zbiorczej liczby cyfr. Ile jest wśród tych liczb jedynek, dwójek, itd.? Która cyfra występuje najczęściej? Jeżeli zauważysz pewne dziwne zależności (np. jakaś cyfra będzie występować najcześciej) - zapoznaj się z Rozkładem Benforda, i popatrz, czy przypadkiem Twoje liczby go nie spełniają...



