Es una traduccion del repo 

https://github.com/robincoello/linux-a-distancia

1 Este documento
2
3
4 
5
6
7
8
9
10 
11 
12 ** 
13 
14 
15 
16 
17 
18 
19 
20 
21 
22 
23 
24 
25 
26 

Los comandos que son en ingles propios de linux no se debe traducir (dejar en ingles)

1 Sistema Operativo - 1 Secuencia 

Instalación de su sistema – 1

Utilizaremos la distribución Debian GNU/Linux (última version estable). Se trata de un sistema código abierto (gratuito).

Étant donné que votre ordinateur est vraisemblablement doté d’un système Windows, une solution simple, pour y faire fonctionner le système Debian GNU/Linux,
est d’y installer d’abord un logiciel appelé : « machine virtuelle » ou « logiciel de virtualisation ».

Debido a que probablemente su computadora posea un sistema Windows, una solución simple para hacer funcionar el sistema Debian GNU/Linux es la de instalar un software llamada "maquina virtual" o "solución de virtualización"

usaremos Oracle VM VirtualBox. Il s’agit d’un logiciel libre (gratuit).

usaremos Oracle Vm VirtualBox. El cual es un software(gratuito) -la traducción literal sería "se trata de un software libre(gratuito)", aunque me parece más adecuado lo que escribí anteriormente-

Site de Oracle VM VirtualBox :
Sitio de Oracle VM VirtualBox :

https://www.virtualbox.org/

Pour le téléchargement de ce logiciel : Downloads – Oracle VM VirtualBox :
Para descargar este software :

https://www.virtualbox.org/wiki/Downloads

o :

http://www.oracle.com/technetwork/server-storage/virtualbox/downloads/index.html

Ce logiciel de virtualisation est appelé : « système hôte ». Il peut être installé sur les systèmes d’exploitation Windows, Mac OS X, Linux, ou Solaris.
Este software de virtualización se llama "sistema host o sistema anfitrión -no se si sería más adecuado en este caso solución de virtualizacipon o hypervisor-", Puede ser instalado en sistemas operativos Windows, Mac OS X, Linux, o Solaris

Lorsqu’il est installé et exécuté, on peut alors y installer un ou plusieurs systèmes d’exploitation appelés : « systèmes invités », puis les faire fonctionner. Un tel système$
Una vez que 

En réalité, un système invité ne fonctionne pas directement sur le matériel et n’accède pas directement à celui-ci. Lorsqu’il est installé dans Oracle VM VirtualBox, tout ce q$

Pour vous initier à l’utilisation de Oracle VM VirtualBox : [PDF] Oracle VM VirtualBox Manuel de l'utilisateur :

http://download.virtualbox.org/virtualbox/UserManual_fr_FR.pdf

LABORATOIRE – Installation de votre système – 1
Laboratorio - Instalación de 

Téléchargez et prenez connaissance du manuel de l’utilisateur de Oracle VM VirtualBox (les « Premiers pas » de ce manuel vous suffiront).

Téléchargez et installez Oracle VM VirtualBox sur votre ordinateur.

2 Installation de votre système – 2

Oracle VM VirtualBox étant maintenant installé, il faut télécharger ce qui est nécessaire en vue de l’installation du système invité : le système Debian GNU/Linux.

Site de Debian GNU/Linux :

https://www.debian.org/index.fr.html

Si vous disposez d’une connexion à Internet, il n’est pas nécessaire de télécharger tous les paquets de composants du système (tous les CD ou tous les DVD). Une procédure d’installation possible consiste à ne télécharger d’abord que le premier CD. Ensuite, lors de l’installation dans Oracle VM VirtualBox, puis lors d’installations ultérieures de composants supplémentaires, les paquets nécessaires qui manquent seront téléchargés depuis une localisation dans l’Internet, appelée : « miroir ».

Images ISO de CD/DVD :

https://www.debian.org/CD/

Télécharger les images des CD ou DVD par HTTP ou FTP :

https://www.debian.org/CD/http-ftp/

Puis : images officielles des CD ou DVD de la distribution « stable » :

https://www.debian.org/CD/http-ftp/#stable

Il faut ensuite sélectionner le type de processeur de votre ordinateur (voir les caractéristiques de celui-ci).

Si votre ordinateur est un PC 64 bits, sélectionnez : amd64.

Et si la référence de la dernière version stable est 8.6.0 :

http://cdimage.debian.org/debian-cd/8.6.0/amd64/iso-cd/

Il faut ensuite télécharger le fichier : debian-8.6.0-amd64-CD-1.iso

C’est le premier fichier d’image disque du système : le fichier du CD-1 au format « .iso ».

Son URL est donc :

http://cdimage.debian.org/debian-cd/8.6.0/amd64/iso-cd/debian-8.6.0-amd64-CD-1.iso

En ce qui concerne la documentation d’installation du système Debian GNU/Linux :

Accéder au : Guide de l'installation :

https://www.debian.org/releases/stable/i386/apa

Accéder au : Manuel d'installation, selon le type de processeur de votre ordinateur :

https://www.debian.org/releases/stable/installmanual



LABORATOIRE – Installation de votre système – 2



Selon le type de processeur de votre ordinateur, téléchargez le fichier CD-1.iso du système Debian GNU/Linux (dernière version stable).

Prenez connaissance des documents : Guide de l'installation et Manuel d'installation, selon le type de processeur de votre ordinateur.



3 Installation de votre système

Oracle VM VirtualBox ayant été installé, et le fichier CD-1.iso du système Debian GNU/Linux (dernière version stable) étant maintenant téléchargé et enregistré dans le système de fichiers de votre ordinateur, il faut installer le système Debian GNU/Linux (système invité) dans Oracle VM VirtualBox (système hôte).

LABORATOIRE – Installation de votre système – 3

Démarrez VirtualBox. Une fenêtre appelée « gestionnaire VirtualBox » apparaît.

Il faut d’abord créer une nouvelle machine virtuelle pour Debian GNU/Linux, en vous référant au manuel de l’utilisateur de Oracle VM VirtualBox.

Cliquez sur le bouton « Nouvelle » de la fenêtre du gestionnaire. Un assistant apparaît pour vous guider à travers le paramétrage de la nouvelle machine virtuelle (VM).

L’assistant vous demande notamment :

le nom de la machine virtuelle (par exemple : debian-8.6.0)
le type de système d’exploitation
la taille de mémoire (RAM) que VirtualBox devra affecter à chaque démarrage de la machine virtuelle
la création d’un nouveau disque virtuel, c’est-à-dire d’un nouveau fichier image de disque (par défaut au format « .vdi »), dans le dossier de la nouvelle machine virtuelle, créé par VirtualBox dans le système de fichiers de votre ordinateur

Après avoir terminé, dans la fenêtre du « gestionnaire VirtualBox », la nouvelle machine virtuelle apparaît.

Pour que, pendant l’installation du système Debian GNU/Linux, les paquets nécessaires puissent être téléchargés depuis une localisation dans l’Internet, appelée : « miroir », allez dans « Configuration », « Réseau » et sélectionnez « Accès par pont ».

Démarrez la nouvelle machine virtuelle. Une fenêtre s’ouvre, dans laquelle vous pourrez installer le nouveau système d’exploitation. Lors du premier démarrage de cette machine virtuelle, un assistant apparaît pour la sélection d’un média d’installation. Sélectionnez donc le fichier CD-1.iso du système Debian GNU/Linux enregistré dans votre système de fichiers. Il peut être nécessaire d’aller dans « Configuration », « Stockage » et de le sélectionner.

Ensuite, installez Debian GNU/Linux, en vous référant aux documents : Guide de l'installation et Manuel d'installation.



4 # Préliminaire : définition d’un système d’exploitation

Avant d’utiliser Debian GNU/Linux dans la séquence 5, un peu de lecture concernant la notion de système d’exploitation, en particulier le système Linux.

0.     Préliminaire : définition d’un système d’exploitation

Vision matérielle simplifiée d’un ordinateur

Un ordinateur monoprocesseur est composé, dans son « cœur », d’un processeur central, lui-même composé d’une unité de commande, d’une unité arithmético-logique (UAL), où sont exécutées les opérations binaires, et de registres, d’une mémoire centrale, et bien souvent d’autres mémoires intermédiaires.

Toutes les informations manipulées par les utilisateurs (images, sons, textes, etc.), afin de pouvoir être traitées par la machine, sont décomposées en séquences d’informations élémentaires, représentées par des suites de chiffres binaires (0 ou 1), appelés « bits » (binary digits), structurées généralement en séquences d’« octets » ou de « bytes » (suites de 8 bits).

Les informations manipulées par les utilisateurs (les « données »), ainsi que les « programmes », sont rangés dans des fichiers, eux-mêmes regroupés en dossiers, conservés sur des mémoires externes (disques magnétiques, bandes magnétiques, disques optiques ou magnéto-optiques, etc.).

En outre, les utilisateurs communiquent avec la machine via des périphériques, reliés au système au moyen d’interfaces (cartes de périphériques).

Un périphérique d’« entrée » permet à un utilisateur de donner de l’information au coeur du système. Un périphérique de « sortie » permet à un utilisateur de recevoir de l’information du coeur du système.

Bien sûr, certains périphériques sont à la fois d’entrée et de sortie. C’est le cas par exemple des « écrans tactiles ». Les mémoires externes sur lesquelles on peut lire et écrire peuvent aussi être considérées comme faisant partie de cette catégorie. Les lecteurs de DVD ou de CD-ROM (Compact Disk - Read Only Memory) n’en font évidemment pas partie. Par contre, les « graveurs » en font partie.

Comme périphériques d’entrée/sortie, il faut encore mentionner les interfaces permettant à un ordinateur d’être relié à un réseau. Il peut s’agir d’une interface purement digitale (traitant des signaux binaires) ou d’un modem (modulateur-démodulateur : interface digitale-analogique réalisant la conversion des signaux binaires en signaux sinusoïdaux et vice versa).



0.1.    Structure en couches logicielles d’un ordinateur


Un ordinateur est une machine capable de résoudre des problèmes en appliquant des instructions préalablement définies. La suite des instructions décrivant la façon dont l’ordinateur doit effectuer un certain travail est appelée programme. Il s’agit des instructions qui peuvent être « exécutées » par l’unité de contrôle de l’ordinateur.

Les circuits électroniques de chaque ordinateur ne pouvant reconnaître et exécuter directement qu’un nombre très limité d’instructions, tout programme doit être, avant son exécution, converti pour n’être exprimé qu’avec ces instructions. L’ensemble des instructions exécutables directement par un ordinateur forme un langage qui permet en fin de compte aux gens de communiquer avec cet ordinateur. C’est ce qu’on appelle le « langage machine ». Lorsqu’on conçoit un nouvel ordinateur, il convient de choisir les instructions qui formeront son langage machine. En général on cherche à ce que ces instructions soient les plus simples possible, compte tenu des performances attendues de la machine, afin de réduire la complexité (et donc le coût) des circuits électroniques nécessaires. Le problème est que ces langages machine sont alors si primitifs qu’il est extrêmement pénible et fastidieux de les utiliser.

Il y a deux façons de résoudre ce problème qui, toutes les deux, visent à construire un nouveau jeu d’instructions plus pratique à utiliser que le langage machine. L’ensemble de ces nouvelles instructions forme un langage, que nous appellerons LPP, de la même façon que l’ensemble des instructions machine forme le langage appelé LM. Les deux techniques qui existent diffèrent suivant la manière dont les programmes écrits en LPP sont traités par l’ordinateur qui, en dernière analyse, ne peut exécuter que des programmes écrits dans son langage machine LM.

La première façon d’exécuter un programme écrit en LPP est de remplacer dans un premier temps chaque instruction de ce programme par la suite d’instructions en LM qui est équivalente. On obtient ainsi un nouveau programme entièrement écrit en LM. L’ordinateur exécute alors ce nouveau programme en LM et non pas l’ancien programme en LPP. Cette technique est appelée traduction ou compilation, et le programme qui traduit chaque instruction LPP en séquence LM s’appelle un compilateur. Le résultat de la traduction réside dans des fichiers exécutables.

La seconde façon est d’écrire un programme capable, après avoir examiné chaque instruction d’un programme en LPP, d’exécuter directement la séquence d’instructions en LM équivalente. Cette technique, avec laquelle on n’a pas besoin de générer tout un programme en LM, est appelée interprétation, et le programme capable de réaliser ce traitement s’appelle un interpréteur.

Traduction et interprétation se ressemblent beaucoup. Dans chaque cas toute instruction en LPP est finalement convertie en une suite équivalente d’instructions en LM. Mais il faut noter que dans le cas de la traduction tout le programme en LPP est d’abord converti en un programme en LM ; puis le programme en LPP « disparaît » et c’est le programme en LM qui est exécuté. En revanche, avec l’interprétation chaque instruction du programme en LPP est analysée puis immédiatement exécutée. On n’obtient donc pas de programme traduit. Ces deux méthodes sont toutes deux très utilisées.

Plutôt que de raisonner en termes de traduction ou d’interprétation, il est souvent plus facile d’imaginer l’existence d’un ordinateur hypothétique, ce qu’on appelle une machine virtuelle, dont le langage machine est LPP. Si on pouvait construire réellement cette machine à un coût raisonnable, on n’aurait pas besoin du langage LM ni d’une machine capable d’exécuter ce langage LM. Les gens écriraient tout simplement leurs programmes en LPP et la machine exécuterait directement ces programmes. Bien qu’il soit trop coûteux de construire une telle machine, on peut néanmoins écrire des programmes en LPP à condition que ces programmes puissent être traduits ou interprétés par un ordinateur (réel celui-ci) dont le langage est LM. En d’autres termes, on peut écrire des programmes pour des machines virtuelles comme si elles existaient réellement.

Pour que la traduction ou l’interprétation reste assez simple il faut que les langages LM et LPP ne soient pas trop différents. Ceci implique que très souvent LPP, bien que « meilleur » que LM, n’est pas le langage idéal pour l’écriture d’un programme. Ceci est quelque peu contradictoire avec l’idée même qui avait présidé à l’élaboration de LPP : faire un langage plus simple et plus pratique à utiliser par des êtres humains que LM, qui était plutôt fait pour la machine.

On peut être alors amené à définir un nouvel ensemble d’instructions plus proches de l’utilisateur final et donc moins dépendantes de la machine que celles de LPP. Ces nouvelles instructions forment un langage LN à l’aide duquel, comme on l’a vu plus haut, on pourra écrire des programmes comme si une machine basée sur LN existait. Ces programmes seront, bien sûr, traduits en LPP par un compilateur ou interprétés par un interpréteur.

On peut ainsi concevoir toute une série ou « hiérarchie » de langages, chacun étant un peu plus pratique que son prédécesseur, jusqu’à ce qu’on en obtienne un jugé « convenable ». Chaque langage s’appuie sur son prédécesseur de telle sorte qu’on peut voir un ordinateur comme un empilement de couches ou de niveaux. Le langage « du bas » est le plus simple, celui « du haut » est le plus complexe.

Un ordinateur composé de n couches peut être vu comme n machines virtuelles distinctes, chaque machine virtuelle ayant son propre langage.



Vision logique simplifiée d’un ordinateur

La plupart des programmeurs travaillant sur une machine virtuelle de niveau n ne sont intéressés que par le niveau supérieur, celui qui ressemble le moins au langage machine du bas niveau. Cependant, les gens qui s’intéressent au fonctionnement d’un ordinateur doivent étudier tous les niveaux. Ceux qui veulent concevoir de nouvelles machines ou de nouvelles couches (c’est-à-dire de nouvelles machines virtuelles) doivent aussi connaître d’autres couches que les couches supérieures.

Le niveau 0, le premier à intéresser l’informaticien, est appelé niveau physique. Les objets manipulés à ce niveau sont des portes qui, bien que fabriquées à l’aide de composants analogiques comme les transistors, peuvent être vues comme des composants logiques. Chaque porte dispose d’une ou plusieurs entrées logiques (signaux représentant 0 et 1) et rend comme résultat une fonction simple de ces entrées comme ET ou OU. Pour être complet, nous devrions mentionner l’existence d’un niveau encore inférieur au niveau 0. Ce niveau, appelé « niveau composant », est du domaine du génie électrique. Ce qui est manipulé à ce niveau, ce sont des transistors qui sont en quelque sorte les « atomes » des concepteurs de machines. Bien sûr, on pourrait aussi parler du fonctionnement des transistors, mais cela nous entraînerait dans la physique du solide !

A la différence du niveau 0 dans lequel il n’y a pas trace de programme (ou suite d’instructions à exécuter), le niveau 1 comprend un programme, appelé microprogramme, qui se situe généralement dans des mémoires mortes (ROM - Read Only Memory) et dont le travail est d’interpréter les instructions du niveau supérieur. Nous appellerons ce niveau le niveau microprogrammé. Deux ordinateurs différents n’ont jamais exactement des niveaux microprogrammés identiques. Peu de machines disposent à ce niveau de plus de vingt instructions et la plupart de ces instructions consistent à transférer des données ou à exécuter des tests très simples.

Chaque microprogramme définit implicitement un langage de niveau 2 (et une machine virtuelle dont le langage machine est ce langage). Toutes les machines de niveau 2 ont beaucoup de points communs, même si ces machines sont de constructeurs différents. Ce niveau sera appelé niveau machine traditionnel. Les constructeurs d’ordinateurs proposent avec chaque machine qu’ils vendent un manuel souvent intitulé « Manuel de référence du langage machine » ou quelque chose d’approchant. Ces manuels se rapportent au niveau 2 (machine virtuelle) et non au niveau 1. Le jeu d’instructions machine qui y est décrit (entre 50 et 300 instructions) est l’ensemble des instructions interprétées par le microprogramme et non pas l’ensemble des instructions exécutables directement par le matériel.

Notons cependant que certains ordinateurs, et surtout les plus anciens, ne disposent pas de couche microprogrammée. Sur ces machines, les instructions de niveau 2 sont traitées directement par le matériel (niveau 0), sans intervention d’un niveau 1. Nous ne tiendrons pas compte de ces exceptions et nous dirons que le niveau machine conventionnel est le niveau 2.

Le troisième niveau, appelé système d’exploitation (Operating System), est souvent un niveau hybride. La plupart des instructions de son langage figurent également dans le langage du niveau 2. Il n’y a en effet aucune raison pour qu’une instruction d’un niveau ne puisse être présente à un autre niveau. On y trouve également des instructions spécifiques à ce niveau (dites « de niveau système d’exploitation »), une organisation de la mémoire différente, la capacité d’exécuter plusieurs programmes en parallèle, etc. Il existe plus de différences entre les machines de niveau 3 qu’entre les machines de niveau 1 ou 2. Les instructions du niveau 3 qui sont identiques à des instructions du niveau 2 sont traitées directement par le niveau inférieur et non par le système d’exploitation. C’est en cela que ce niveau est hybride.

Il y a un fossé entre les niveaux 3 et 4. Les quatre niveaux inférieurs ne concernent pas directement le programmeur moyen. Ils sont là principalement pour supporter les traducteurs et interpréteurs dont les niveaux supérieurs ont besoin et sont écrits par des spécialistes de la réalisation de machines virtuelles : les programmeurs système. Les niveaux 4 et au-dessus concernent le programmeur d’applications qui a un problème concret à résoudre.

Autre caractéristique qui apparaît au niveau 4 : la méthode de support des niveaux supérieurs. Les niveaux 2 et 3 sont toujours interprétés. Les niveaux 4 et 5 sont souvent, mais pas toujours, associés à des compilateurs (traducteurs).

Il y a d’ailleurs une autre différence entre les niveaux 1, 2 et 3 d’un côté et les niveaux supérieurs de l’autre : c’est la nature du langage fourni. Les langages « machine » des niveaux 1, 2 et, originellement, 3 sont numériques, donc très difficiles à utiliser puisqu’ils consistent en de longues séries de nombres. Au contraire, à partir du niveau 4, on trouve des langages qui contiennent des mots ou des abréviations beaucoup plus compréhensibles par les êtres humains (mais pas par la machine).

Le niveau 4 est appelé assembleur (il s’agit d’un traducteur). Le langage d’assemblage est une forme symbolique d’un des langages sous-jacents. Ce niveau permet d’écrire des programmes pour les niveaux 1, 2 et 3 dans une forme moins déplaisante. Les programmes en langage d’assemblage sont d’abord traduits en langage de niveau 1, 2 ou 3 puis sont interprétés par la machine virtuelle ou réelle correspondante. Si ce langage d’assemblage était autrefois très important, on peut dire qu’aujourd’hui son importance est bien moindre. En effet, au cours du temps, de nouvelles couches ont été insérées dans la structure de manière à assurer une certaine « portabilité » des développements de haut niveau. C’est par exemple l’objectif des plates-formes « JAVA » ou « .NET ».

Au niveau 5 on trouve des compilateurs ou des interpréteurs de langages conçus pour être utilisés par des programmeurs d’applications. Ces langages, souvent appelés langages de haut niveau, sont extrêmement nombreux : il en existe des centaines.

Les niveaux 6 et au-dessus sont des ensembles de programmes conçus pour créer des machines virtuelles destinées à traiter des applications déterminées.

En résumé, on peut dire qu’un ordinateur peut être vu comme une suite de couches, chaque couche englobant toutes les couches précédentes. Une couche représente un certain niveau d’abstraction et comporte divers objets et opérations sur ces objets. L’ensemble des types de données, des opérations et des caractéristiques de chaque niveau s’appelle l’architecture de ce niveau.




0.2.    Logiciel système




On peut dire que les logiciels (les programmes) se répartissent en deux grandes catégories : les programmes système qui permettent le fonctionnement de l’ordinateur, et les programmes d’application qui résolvent les problèmes des utilisateurs.

Le système d’exploitation (Operating System) est le plus important des programmes système. Il contrôle les ressources de l’ordinateur et fournit la base sur laquelle seront construits les programmes d’application.

Le reste du logiciel système se trouve au-dessus du système d’exploitation. On y trouve l’interpréteur de commandes (Shell), les éditeurs, les compilateurs et d’autres programmes qui ne dépendent pas non plus des programmes d’application. Ces programmes ne font pas partie du système d’exploitation, même s’ils sont fournis par les constructeurs d’ordinateurs.

Le système d’exploitation est la partie du logiciel système qui fonctionne en mode noyau ou mode superviseur. Elle est protégée par le matériel contre les erreurs de manipulation des utilisateurs. Les éditeurs et les compilateurs fonctionnent en mode utilisateur. Si un utilisateur n’apprécie pas un compilateur donné, il peut, s’il le souhaite, écrire son propre compilateur ; en revanche, il ne peut pas écrire son propre gestionnaire des interruptions du disque car celui-ci fait partie du système d’exploitation et est protégé par le matériel contre les tentatives des utilisateurs pour le modifier.

Enfin, au-dessus des programmes système, on trouve les programmes d’application. Ces programmes sont écrits par les utilisateurs ou par les éditeurs de logiciels pour résoudre des problèmes spécifiques tels que le traitement des données commerciales, les calculs scientifiques ou les jeux.




0.3.    A quoi sert un système d’exploitation ?




L’interface entre le système d’exploitation et les programmes de l’utilisateur est constituée d’un ensemble d’« instructions étendues » fournies par le système d’exploitation. Ces instructions étendues sont qualifiées d’appels système. Les appels système créent, détruisent et utilisent divers objets logiciels gérés par le système d’exploitation. Les processus et les fichiers sont les plus importants de ces objets.

Fondamentalement, un système d’exploitation rend des services en tant que machine étendue et en tant que gestionnaire de ressources. Cette double vocation d’un système d’exploitation est développée ci-après.



0.3.1.   Le système d’exploitation en tant que machine étendue



Un ordinateur se compose d’un ou de plusieurs processeurs (unités de contrôle), d’une mémoire principale, d’horloges, de mémoires secondaires ou externes, de terminaux, et d’interfaces de connexion à d’autres périphériques d’entrées/sorties (E/S) et à des réseaux. Il s’agit d’un système fort complexe. L’écriture de programmes qui prennent en compte tous ces éléments et les gèrent correctement est un travail extrêmement difficile.

Au niveau du langage machine, les périphériques d’E/S sont contrôlés en chargeant des valeurs spécifiques dans des registres spéciaux, les registres des périphériques. On peut, par exemple, mettre un disque en mode lecture/écriture en chargeant ses registres avec l’adresse du disque, l’adresse de la mémoire principale, le nombre d’octets à transférer et le sens de transfert (lecture ou écriture). Dans la pratique, il faut de nombreux autres paramètres et le code retourné par l’unité de disque à la fin de l’opération est très complexe. De plus, pour de nombreux périphériques d’E/S, le temps est un facteur important lors de la programmation.

A titre d’exemple, la programmation des E/S des lecteurs de disquettes au moyen du contrôleur NEC PD765 (utilisé sur l’IBM PC et de nombreux autres ordinateurs personnels) s’effectue selon 16 commandes qui consistent toutes à charger entre 1 et 9 octets dans un registre de données. Ces commandes permettent de lire et d’écrire des données, de déplacer le bras du lecteur de disquettes, de formater les disquettes ainsi que d’initialiser, tester, restaurer et recalibrer le contrôleur et les lecteurs. Les commandes fondamentales sont READ (lire) et WRITE (écrire), chacune demandant 13 paramètres regroupés dans 9 octets. Ces paramètres spécifient des éléments tels que l’adresse du secteur à lire, le nombre de secteurs par piste, le mode d’enregistrement utilisé sur le support physique, la distance entre deux secteurs consécutifs et ce qu’il faut faire à la rencontre d’une marque d’effacement de données. A la fin de l’opération, le contrôleur retourne 23 champs d’état et d’erreur regroupés dans 7 octets. En outre, le programmeur du lecteur de disquettes doit toujours vérifier si le moteur est allumé ou éteint. S’il est éteint, il faut le mettre en route (et attendre que sa vitesse se stabilise) avant de pouvoir lire ou écrire les données. Il ne faut cependant pas laisser le moteur allumé trop longtemps car cela userait la disquette. Le programmeur système doit donc trouver un compromis entre de longs délais de démarrage et l’usure des disquettes (et les risques de perte de données qui s’ensuivraient).

Bien peu de programmes seraient développés si chaque programmeur devait connaître le fonctionnement détaillé des disques et toutes les erreurs qui peuvent apparaître lors de la lecture d’un bloc de données. Il est donc apparu clairement qu’il fallait trouver un moyen de libérer les programmeurs de la complexité du matériel. Ce moyen, qui a évolué petit à petit, consiste à enrober le matériel d’une couche de logiciel qui gère l’ensemble du système. Il faut en fait présenter au programmeur une interface ou machine virtuelle plus facile à comprendre et à programmer. Cette couche de logiciel est le système d’exploitation.

Le programmeur d’applications souhaite, pour programmer efficacement, une abstraction simple et de haut niveau. Dans le cas des disques, une abstraction typique est que le disque contient des fichiers nommés. Chaque fichier peut alors être ouvert en lecture ou en écriture. Puis il est lu ou écrit et finalement fermé.

Le système d’exploitation est donc le programme qui soustrait le matériel aux regards du programmeur et offre une vue simple et agréable de fichiers nommés qui peuvent être lus ou écrits. Le système d’exploitation masque de la même manière beaucoup d’aspects fastidieux en ce qui concerne les interruptions, les horloges, la gestion de la mémoire et les autres tâches de bas niveau. Dans tous les cas, l’abstraction présentée à l’utilisateur est plus simple et plus facile à utiliser que le matériel sous-jacent.

De ce point de vue, la fonction du système d’exploitation est de présenter à 1’utilisateur l’équivalent d’une machine étendue ou machine virtuelle plus facile à programmer que le matériel.



0.3.2.   Le système d’exploitation en tant que gestionnaire 
            de ressources



Nous venons de voir le système d’exploitation en tant qu’interface de programmation plus commode. Mais on peut également le considérer comme un gestionnaire des fonctions complexes d’un système lui aussi complexe. Les ordinateurs se composent d’une multitude d’éléments. Dans cette optique, le travail du système d’exploitation consiste à ordonner et à contrôler l’allocation des processeurs aux processus et à gérer leur exécution, à gérer l’allocation des mémoires, à gérer l’accès aux fichiers et aux périphériques d’Entrées/Sorties, pour les différents programmes qui en ont besoin.

Par exemple, si trois programmes qui s’exécutent sur un ordinateur essayent simultanément d’imprimer leurs résultats sur la même imprimante, le système d’exploitation peut transférer les résultats à imprimer dans un fichier tampon sur disque. Lorsqu’un programme se termine, le système d’exploitation peut alors copier ses résultats du fichier où ils ont été sauvegardés vers l’imprimante. Simultanément, un autre programme peut continuer à générer des résultats sans se rendre compte qu’il ne les envoie pas directement à l’imprimante.

Lorsqu’un ordinateur est partagé entre plusieurs utilisateurs, la nécessité de gérer et de protéger toutes les ressources est encore plus évidente. De plus, il est souvent nécessaire de partager des informations entre des personnes qui travaillent ensemble. Le rôle principal du système d’exploitation, de ce point de vue, est de connaître à tout moment l’utilisateur d’une ressource, de gérer les accès à cette ressource, d’en accorder l’usage et d’éviter les conflits d’accès entre différents programmes ou entre différents utilisateurs.




0.4.    Historique des systèmes d’exploitation




Les systèmes d’exploitation ont évolué au fil des ans, au gré de l’évolution de l’architecture des ordinateurs sur lesquels ils fonctionnent. Il est donc nécessaire de passer en revue les générations successives d’ordinateurs pour en examiner les apports successifs au niveau du fonctionnement des systèmes d’exploitation.



0.4.1.   La première génération (1945-1955) : les tubes à vide et les 
             cartes enfichables



Vers le milieu des années 40, Howard Aiken de Harvard, John von Neumann de Princeton, Konrad Zuse en Allemagne et d’autres encore ont réussi à construire des machines à calculer au moyen de tubes électroniques. Ces machines étaient énormes, remplissaient des salles entières avec des centaines de tubes et étaient extrêmement lentes et peu fiables.

A cette époque, un seul groupe de personnes concevait, construisait, programmait, utilisait et effectuait la maintenance d’une machine. La programmation se faisait intégralement en langage machine, souvent en reliant électriquement des cartes pour contrôler les fonctions de base de l’ordinateur. Les systèmes d’exploitation étaient alors inconnus. Le mode de fonctionnement consistait à réserver une tranche de temps pour chaque programmeur. Ce dernier se rendait à la salle de l’ordinateur et y insérait son circuit. Pratiquement tous les programmes ne servaient qu’à effectuer des calculs numériques.

Au début des années 50, la procédure s’est améliorée grâce à l’introduction des cartes perforées. Il devint alors possible d’écrire des programmes sur ces cartes et de les faire lire à l’ordinateur plutôt que d’utiliser des cartes électriques à enficher ; le reste de la procédure ne changea pas.



0.4.2.   La deuxième génération (1955-1965) : les transistors et le 
             traitement par lots



L’introduction du transistor vers le milieu des années 50 a complètement modifié la situation. Les ordinateurs devinrent suffisamment fiables pour être construits et vendus à des clients. Pour la première fois, une séparation nette existait entre les concepteurs, les constructeurs, les opérateurs, les programmeurs et le personnel de maintenance.

Ces machines étaient enfermées dans des pièces dont l’air était conditionné et elles étaient surveillées par des opérateurs professionnels. Seules les grandes compagnies, les administrations ou les universités pouvaient se permettre une telle dépense. Pour lancer un travail (job), c’est-à-dire un programme ou un ensemble de programmes, le programmeur devait commencer par écrire le programme en assembleur ou en FORTRAN, puis il devait le mettre sur des cartes perforées. Il apportait ensuite ses cartes à la salle de soumission des travaux et les passait à l’un des opérateurs, qui les donnait à lire à l’ordinateur. Si le compilateur FORTRAN se révélait nécessaire, il fallait également le charger dans l’ordinateur. La majeure partie du temps était perdue en manipulations.

Etant donné le coût élevé des équipements, on chercha à réduire les pertes de temps en adoptant la solution du traitement par lots (batch processing). L’idée directrice était de collecter un ensemble de travaux puis de les transférer sur des bandes magnétiques en utilisant pour cela un petit ordinateur (relativement) peu onéreux, comme l’IBM 1401. D’autres machines, bien plus chères, comme l’IBM 7094, étaient utilisées pour les calculs. L’opérateur devait donc y charger un programme spécial (l’ancêtre des systèmes d’exploitation d’aujourd’hui) qui lisait chaque programme successif et l’exécutait, les résultats étant écrits sur une autre bande magnétique au lieu d’être imprimés directement. A la fin de l’ensemble des travaux, l’opérateur changeait les bandes d’entrée et de sortie, puis les résultats pouvaient être imprimés par un IBM 1401 qui fonctionnait de manière autonome.

A cette époque, la structure d’un travail apporté à la salle de soumission des travaux est la suivante : elle commence par une carte $JOB qui spécifie le temps maximal d’exécution exprimé en minutes, le numéro de compte et le nom du programmeur. Puis vient une carte $FORTRAN qui indique au système d’exploitation de charger le compilateur FORTRAN à partir de la bande système. Elle est suivie de cartes contenant le programme à compiler (fichier source), puis d’une carte $LOAD qui indique au système d’exploitation de charger le code objet qui vient d’être généré (les programmes compilés étaient fréquemment écrits dans des fichiers temporaires et devaient être chargés explicitement). Une carte $RUN indique ensuite à l’ordinateur d’exécuter le programme avec les données qui figurent sur les cartes suivantes. Enfin, une carte $END marque la fin du travail. Ces cartes de contrôle, quoique primitives, étaient les précurseurs des langages de contrôle et des interpréteurs de commandes modernes.

Les ordinateurs de la seconde génération étaient surtout utilisés pour les calculs scientifiques et d’ingénierie. Les systèmes d’exploitation les plus connus étaient FMS (Fortran Monitor System) et IBSYS, le système d’exploitation d’IBM pour le 7094.



0.4.3.   La troisième génération (1965-1980) : les circuits intégrés, 
             la multiprogrammation, etc.



Au début des années 60, la plupart des constructeurs d’ordinateurs avaient deux lignes de produits. Il y avait d’une part les gros ordinateurs scientifiques et d’autre part les ordinateurs commerciaux, largement utilisés par les banques et les compagnies d’assurance pour le tri des bandes et l’impression. Le maintien de ces deux lignes de produits différentes était trop coûteux pour les fabricants. De plus, de nombreux utilisateurs demandaient au départ un petit ordinateur, mais souhaitaient par la suite un ordinateur plus puissant.

IBM tenta de résoudre d’un coup ces deux problèmes en introduisant le système OS/360, conçu à la fois pour les calculs scientifiques et commerciaux, donnant naissance à une série d’ordinateurs compatibles entre eux au niveau du logiciel et ne différant qu’en termes de prix et de performance (mémoire maximale, vitesse du processeur, nombre de périphériques d’E/S connectables, etc.). Comme toutes ces machines possédaient, en théorie du moins, la même architecture et le même jeu d’instructions, les programmes développés pour une machine pouvaient s’exécuter sur toutes les autres. La série 360 fut la première ligne d’ordinateurs à utiliser des circuits intégrés qui lui ont permis d’offrir un rapport performance/coût bien supérieur à celui des ordinateurs de la deuxième génération.

Le point fort du concept d’une « famille unique » fut également son point faible. L’idée de départ était que tout logiciel, y compris le système d’exploitation, devait fonctionner sur tous les modèles. Cependant, il devait être efficace sur ces ordinateurs aux vocations souvent si différentes et il devait répondre à des besoins souvent contradictoires. Il en résulta un OS/360 énorme et très complexe. Il était constitué de millions de lignes en assembleur écrites par des centaines de programmeurs. Elles contenaient des centaines de bogues qu’il fallait sans cesse corriger. Chaque nouvelle version corrigeait quelques bogues mais en introduisait aussi des nouveaux, ce qui fait que le nombre de bogues est pratiquement resté constant au cours du temps.

Malgré son énorme taille et ses problèmes, l’OS/360 et les autres systèmes d’exploitation de la troisième génération ont raisonnablement satisfait leurs utilisateurs, en mettant en application quelques techniques qui faisaient défaut aux ordinateurs de la deuxième génération. La multiprogrammation est probablement la plus importante d’entre elles. Pour que le processeur ne reste pas inoccupé pendant les opérations d’E/S, la mémoire a été partitionnée. Chaque partition différente contient un travail ou une « tâche ». Lorsqu’un travail attend la fin d’une E/S, un autre travail peut utiliser le processeur. Si la mémoire principale peut contenir suffisamment de travaux, le processeur peut être utilisé pratiquement en permanence. La présence simultanée de plusieurs travaux en mémoire nécessite des procédures de contrôle pour protéger chaque travail contre les intrusions ou les erreurs des autres. Les ordinateurs de la troisième génération possédaient ces procédures.



La possibilité de transférer les travaux des cartes vers le disque dès leur arrivée dans la salle des ordinateurs constitue un autre apport des systèmes d’exploitation de la troisième génération. Ainsi, dès qu’un travail se termine, le système d’exploitation peut charger, à partir du disque, un nouveau travail dans la partition qui vient de se libérer, puis il peut lancer son exécution. Cette technique s’appelle SPOOL (Simultaneous Peripheral Operation On Line). Elle est également utilisée pour les sorties, vers les imprimantes par exemple. Les manipulations de bandes et les ordinateurs auxiliaires moins chers sont devenus ainsi inutiles.

Ensuite, pour que les programmeurs ne doivent plus attendre quelques heures entre le moment de la soumission d’un travail et le moment du retrait des résultats, un système à temps partagé (CTSS ou Compatible Time Sharing System) a été mis au point au M.I.T. (Massachusetts Institute of Technology). Il s’agit d’une variante de la multiprogrammation dans laquelle chaque utilisateur possède un terminal en ligne, le processeur étant réparti entre les tâches actives selon un algorithme d’attribution de tour de rôle, le système étant fondé sur l’énorme différence entre les temps de réaction typiques d’un cerveau humain et d’un processeur central. L’ordinateur peut donc répondre rapidement à un nombre important d’utilisateurs et peut même continuer un traitement par lots en arrière-plan lorsque le processeur est libre. Le partage du temps ne s’est vraiment développé qu’avec la généralisation des protections matérielles.

Le M.I.T., les laboratoires Bell et General Electric se sont dès lors attelés au développement d’une machine pouvant supporter la connexion simultanée de plusieurs centaines d’utilisateurs. Le système d’exploitation s’appelait MULTICS (MULTiplexed Information and Computing Service). Depuis lors, le concept d’un gros ordinateur central s’est effrité. Néanmoins, MULTICS a beaucoup influencé les systèmes ultérieurs. Autre événement majeur de la troisième génération : le développement phénoménal des mini-ordinateurs, avec la série des PDP de DEC, dont le plus performant était le PDP-11. L’un des informaticiens des laboratoires Bell, Ken Thompson, qui avait participé au projet MULTICS, se mit à écrire sur un PDP-7 une version simplifiée de MULTICS, appelée d’abord UNICS et ensuite UNIX. Depuis, le système d’exploitation UNIX a bien évolué et s’est répandu dans le marché des 
mini-ordinateurs, des stations de travail et des serveurs, sous diverses versions, jusqu’à son « descendant » LINUX.



0.4.4.   La quatrième génération (depuis 1980) : les LSI et VLSI, 
             les ordinateurs personnels, etc.



L’époque de l’ordinateur personnel est venue avec le développement des circuits LSI (Large Scale Integration), puis VLSI (Very Large Scale Integration). Au niveau de l’architecture, les ordinateurs personnels ne sont pas sensiblement différents des 
mini-ordinateurs. La différence se situe au niveau du prix et de la performance. Le mini-ordinateur a permis aux différents départements des sociétés et des universités d’avoir leur propre ordinateur. Le micro-ordinateur permet à chacun d’avoir son ordinateur. Quant aux « stations de travail », ce ne sont en fait que des ordinateurs personnels puissants. Ces stations sont habituellement connectées à un réseau local (LAN = Local Area Network), concept qui s’est généralisé depuis 1980.

La large disponibilité de puissance de calcul, particulièrement lorsqu’elle est associée à une forte interactivité et à un excellent graphisme, a entraîné le développement d’une importante industrie de production de logiciels pour les ordinateurs individuels.

Des systèmes d’exploitation tels que Microsoft MS-DOS puis Microsoft Windows ont dominé le marché des ordinateurs personnels, tandis que UNIX, puis les versions suivantes de Microsoft Windows ainsi que LINUX ont dominé celui des stations de travail et des serveurs. Les dernières versions de MS-DOS et Windows ont intégré des caractéristiques provenant d’UNIX, ce qui n’est guère surprenant lorsqu’on sait que Microsoft est un des grands fournisseurs des versions successives d’UNIX. Les « stations de travail » (workstations & servers) ont progressivement acquis la puissance d’un mini-ordinateur, pour finir par être dotées d’une puissance même supérieure. Il est donc logique qu’elles utilisent un système d’exploitation initialement conçu pour les mini-ordinateurs.

Depuis le milieu des années 80, on assiste au développement d’interfaces utilisateur conviviales, ainsi qu’à l’élaboration des réseaux d’ordinateurs individuels. En réalité, les idées qui sont à la base de ces développements sont nées il y a bien plus longtemps, dans des lieux tels que le « Xerox Parc ».

Le Xerox Palo Alto Research Center est fondé en 1970 dans le parc industriel de l’université de Stanford, en plein coeur d’une bande de terrain coincée entre la baie de San Francisco et la chaîne côtière de Santa Cruz. A l’époque, ce territoire d’une cinquantaine de kilomètres de longueur, largement à l’état de friche, n’est pas encore baptisé la « Silicon Valley ». Et le Parc entame alors sa mission : imaginer une société du futur, que les responsables de Xerox ont l’intelligence de pressentir sans papier mais avec un ordinateur sur chaque bureau.

Alors que le PC n’existe pas encore, pas plus que Microsoft (1975) ou Apple (1976), le fabricant de photocopieuses imagine déjà les effets d’une société basée sur la communication électronique. Et les chercheurs de haut vol rassemblés au Parc reçoivent pour ambitieuse mission de provoquer cette évolution, afin de mieux l’accompagner et surtout d’en prévenir les conséquences pour Xerox. Concrètement, ils doivent tenter de rendre les ordinateurs faciles à utiliser, ce qui les amènera en réalité à réinventer la relation homme-machine, sur base de leurs innovations techniques mais aussi sur base des travaux de certains psychologues comme le suisse Piaget ou l’américain Bruner.

Il faut préciser qu’au moment de se lancer dans cette entreprise, la compagnie Xerox avait elle-même été marquée par les idées d’un certain Doug Engelbart. Dès les années 50, cet ingénieur visionnaire prétendit en effet qu’un jour le commun des mortels s’installerait devant un ordinateur pour y manipuler textes et graphiques et parvint à convaincre le Pentagone de lui fournir les moyens de réaliser un prototype matérialisant ses concepts. C’est en 1968, lors d’une conférence informatique se tenant à San Francisco, que l’ingénieur présenta son appareil : une console branchée, par le biais d’un réseau de télévision, sur l’ordinateur du Stanford Research Institute. A l’aide d’un écran, d’un clavier, d’un dispositif de pointage (un ancêtre de la souris) et d’un système complexe de clefs à bascule lui permettant de fournir des instructions à l’ordinateur, Engelbart fit d’un seul coup la démonstration du premier traitement de texte et du premier système fonctionnel d’hypertexte.

Mais à l’époque, les éminences n’ont pas encore foi en l’informatique personnelle, et restent largement insensibles à cette démonstration. Critiqué pour le coût de ses travaux, Engelbart doit mettre la clef de son laboratoire sous le paillasson.

Plus tard, Xerox embauche une partie des ingénieurs d’Engelbart ainsi que d’autres, et dès 1973, le fabricant de photocopieuses dispose d’un ordinateur qui influencera tous les autres. Cet ordinateur, baptisé « Alto », dispose d’une souris à trois boutons, d’une unité à cassettes et est connecté, via un réseau local Ethernet, à d’autres ordinateurs ainsi qu’à une imprimante laser (autant d’inventions dues en partie à Xerox). Quant à l’interface graphique, elle met en scène des icônes et des fenêtres à menus déroulants, tandis que Smalltalk, le langage de programmation, est orienté objet et repose sur la notion d’informatique distribuée.

Confronté à la concurrence des Japonais sur le marché de la photocopie, Xerox commet l’énorme bourde de ne pas donner suite à ces inventions. L’histoire démontre que Xerox aurait été bien avisé de se lancer dans ce créneau commercial encore vierge à l’époque. Car en effet Steve Wozniak et Steve Jobs ont été admirer l’ordinateur Alto. Les fondateurs d’Apple s’en inspireront en créant le « Lisa » en 1983 puis, l’année suivante, leur premier Macintosh, lequel inspira à son tour un certain Bill Gates ...

Aujourd’hui, les réseaux d’ordinateurs individuels peuvent fonctionner sous des systèmes d’exploitation en réseau ou sous des systèmes d’exploitation distribués.

Sous un système d’exploitation en réseau, les utilisateurs connaissent l’existence des différents ordinateurs, peuvent se connecter sur une machine distante et transférer des fichiers d’une machine à une autre. Chaque ordinateur fonctionne sous son propre système d’exploitation.

A l’opposé, un système d’exploitation distribué apparaît aux yeux de ses utilisateurs comme une seule machine, même si ce n’est pas le cas. Dans un vrai système distribué, les utilisateurs n’ont pas à savoir où se trouvent leurs fichiers et où leurs programmes sont exécutés. Le système d’exploitation doit s’acquitter de sa tâche de manière transparente pour l’utilisateur.

Par contre, un système d’exploitation en réseau ne diffère pas fondamentalement d’un système classique. Il requiert évidemment une interface réseau et son logiciel de contrôle de bas niveau, ainsi que des programmes qui permettent une connexion distante et un accès aux fichiers qui se trouvent sur les différentes machines. Mais ces ajouts ne modifient pas la structure fondamentale du système d’exploitation.




0.5.    Linux




Voici des sites à consulter :

https://fr.wikipedia.org/wiki/Linux

http://linuxfr.org/

Linux est le nom donné à tout système d’exploitation libre (à toute « distribution ») fonctionnant avec le noyau Linux. C’est un développement libre du système UNIX respectant les spécifications POSIX.

Ce système est né de la rencontre entre le mouvement du logiciel libre et le modèle de développement collaboratif et décentralisé via Internet. Son nom vient du créateur du noyau Linux, Linus Torvalds.

La Free Software Foundation utilise plutôt le nom GNU/Linux, étant donné que depuis 1992, le noyau est distribué sous la Licence publique générale GNU (GNU GPL, écrite par Richard Stallman).

En ce qui concerne l’identification du noyau, toute version PAIRE (par exemple 2.6) est une version de noyau STABLE, tandis que toute version IMPAIRE (par exemple 2.5) est une version de noyau INSTABLE (en phase de test).

Les distributions Debian GNU/Linux, Ubuntu, etc. utilisent les paquets au format « .deb », tandis que les distributions Red Hat Enterprise Linux, CentOS, etc. utilisent les paquets au format « .rpm ».

Une connexion au système Linux (Ouverture de session ou Logging on), dans l’interface graphique ou dans une interface texte, nécessite la saisie de deux informations : Username et Password.

Pour administrer le système, il faut se connecter en « superutilisateur » (Username = « root »), qui a tous les droits sur le système, et non en utilisateur ordinaire.

La totalité des tâches d’administration du système (station de travail ou serveur) peut être réalisée par des commandes en mode texte.

Les tâches les plus simples peuvent être accomplies via des outils en mode graphique. Les outils sont propres à chaque distribution. Cependant, certains outils, tels que linuxconf ou Webmin, sont multi distribution.


5 Connexions, outils et accès aux manuels – 1

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Debian GNU/Linux étant installé, vous pouvez ouvrir une session en tant qu’utilisateur de ce système.

Dans le « bureau » qui se présente alors à vous (interface graphique), dans le menu « Activités », puis « Afficher les applications », vous pouvez lancer un interpréteur de commandes (Shell), c’est-à-dire un « Terminal ».

Lorsqu’on tape une commande dans un terminal puis qu’on pousse sur la touche « Entrée » pour l’exécuter, c’est un programme (soit « système », soit « d’application ») que l’on exécute. Une commande possède un « nom » et, selon l’utilisation que l’on en fait, peut être assortie d’une ou de plusieurs « options » et d’un « argument ». Par exemple, dans la commande : « ls -l /usr/bin/who », « ls » est le nom, « -l » est une option et « /usr/bin/who » est l’argument.

Un système Linux est absolument sensible à la casse ; par exemple, la commande : « ls » DOIT être tapée « ls », et NON « LS » ou « Ls » ! Il en est de même pour tout nom d’utilisateur, de groupe, de fichier, de répertoire, etc.



Quelques commandes :



Chaque commande est annoncée avec son numéro de section de manuel (affichez au sujet des sections le manuel de la commande man en tapant la commande : « man man »).

Dans une page de manuel, dans la rédaction d’un synopsis de commande, ce qui est non obligatoire est écrit entre crochets droits : « [   ] ».



PWD(1)

NOM
       pwd - Afficher le nom du répertoire de travail actuel

SYNOPSIS
       pwd [OPTION]...

DESCRIPTION
       Afficher le nom complet du répertoire de travail actuel.



WHOAMI(1)

NOM
       whoami - Afficher l'identifiant d'utilisateur

SYNOPSIS
       whoami [OPTION]...

DESCRIPTION
       Afficher le nom d'utilisateur associé à l'utilisateur courant. Identique à id -un.



ID(1)

NOM
       id - Afficher les identifiants d'utilisateur et de groupe effectifs et réels

SYNOPSIS
       id [OPTION]... [UTILISATEUR]

DESCRIPTION
       Afficher les informations sur l'utilisateur et le groupe pour UTILISATEUR,
       ou (si UTILISATEUR est omis) pour l'utilisateur courant.



WHO(1)

NOM
       who - Montrer qui est connecté

SYNOPSIS
       who [OPTION] ... [ FICHIER | PARAM1 PARAM2 ]

DESCRIPTION
       Afficher les informations sur les utilisateurs actuellement connectés.



EXIT(3)

NAME
       exit - Cause normal process termination



En particulier, la commande « exit » permet de terminer l’exécution d’un processus bash (Interpréteur de commandes GNU Bourne-Again Shell).



CLEAR(1)

NAME
       clear - Clear the terminal screen

SYNOPSIS
       clear

DESCRIPTION
       clear clears your screen if this is possible. It looks in the environment
       for the terminal type and then in the terminfo database to figure out
       how to clear the screen.



MAN(1)

NOM
       man - Interface de consultation des manuels de référence en ligne



INFO(1)

NAME
       info - Read Info Documents

SYNOPSIS
       info [OPTION]... [MENU-ITEM...]

DESCRIPTION
       Read documentation in Info format.



WHATIS(1)

NOM
       whatis - Afficher la description des pages de manuel

SYNOPSIS
       whatis [-dlhvV] [-r|-w] [-s liste] [-m système[,...]] [-M chemin]
       [-L locale] [-C fichier] nom ...

DESCRIPTION
       Chaque page de manuel comporte une description courte.
       whatis recherche des pages de manuel dont le nom correspond à nom
       et affiche leur description courte.



HOSTNAME(1)

NAME
       hostname - Show or set the system's host name



UNAME(1)

NOM
       uname - Afficher des informations sur le système

SYNOPSIS
       uname [OPTION]...

DESCRIPTION
       Afficher certaines informations concernant le système.
       Le comportement est identique si OPTION est omis ou -s est utilisée.



DATE(1)

NOM
       date - Afficher ou configurer la date et l'heure du système

SYNOPSIS
       date [OPTION]... [+FORMAT]
       date [-u|--utc|--universal] [MMJJhhmm[[CC]AA][.ss]]

DESCRIPTION
       Afficher dans le FORMAT indiqué ou configurer la date et l'heure du système.



CAL(1)

NAME
       cal, ncal - Displays a calendar and the date of Easter

SYNOPSIS
       cal [-3hjy] [-A number] [-B number] [[month] year]
       cal [-3hj] [-A number] [-B number] -m month [year]
       ncal [-3bhjJpwySM] [-A number] [-B number] [-s country_code] [[month] year]
       ncal [-3bhJeoSM] [-A number] [-B number] [year]
       ncal [-CN] [-H yyyy-mm-dd] [-d yyyy-mm]

DESCRIPTION
       The cal utility displays a simple calendar in traditional format
       and ncal offers an alternative layout, more options and the date of Easter.
       The new format is a little cramped but it makes a year fit on a 25x80 terminal.
       If arguments are not specified, the current month is displayed.



PASSWD(1)

NOM
       passwd - Modifier le mot de passe d'un utilisateur

SYNOPSIS
       passwd [options] [LOGIN]

DESCRIPTION
       La commande passwd modifie les mots de passe des comptes d'utilisateurs.
       Un utilisateur normal ne peut changer que son propre mot de passe, alors que
       le superutilisateur peut changer le mot de passe associé à n'importe quel compte.
       passwd modifie également les dates de fin de validité du compte
       ou du mot de passe associé.



Un fichier :



PASSWD(5)

NOM
       passwd - Fichier des mots de passe (en fait : d’informations sur les comptes utilisateurs)

DESCRIPTION
       /etc/passwd contient différentes informations sur les comptes utilisateurs.
       Ces informations consistent en sept champs séparés par des deux-points (« : ») :

       ·   nom de connexion de l'utilisateur (« login »)

       ·   un mot de passe chiffré optionnel

       ·   l'identifiant numérique de l'utilisateur

       ·   l'identifiant numérique du groupe de l'utilisateur

       ·   le nom complet de l'utilisateur ou un champ de commentaires

       ·   le répertoire personnel de l'utilisateur

       ·   l'interpréteur de commandes de l'utilisateur (optionnel)

       Le champ du mot de passe chiffré peut être vide. Dans ce cas,
       aucun mot de passe n'est nécessaire pour s'authentifier avec le compte donné.
       Cependant, certaines applications qui lisent le fichier /etc/passwd peuvent décider
       de ne donner aucun accès si le mot de passe est vide. Si le mot de passe est un « x »
       minuscule, alors le mot de passe chiffré se trouve dans le fichier shadow(5) ;
       il doit y avoir une ligne correspondante dans le fichier shadow,
       sinon le compte de l'utilisateur n'est pas valide.



LABORATOIRE – Connexions, outils et accès aux manuels – 1



Dans la distribution Debian GNU/Linux :



Une connexion au système (Ouverture de session ou Logging on), dans l’interface graphique ou dans une interface texte (il y en a six : tty1 à tty6), nécessite la saisie de deux informations : Username + Password.

Connectez-vous au système en mode graphique (Username + Password)

Vous venez de lancer une « session utilisateur » en mode graphique.



Etant connecté au système en mode graphique, basculez en mode texte (dans tty6)

CTRL+ALT+F6



Etant en mode texte dans tty6, basculez en mode texte dans tty2

ALT+F2



Connectez-vous (Username + Password)

Vous venez de lancer une « session utilisateur » en mode texte.

Dans la suite, pour la nécessité des explications, le mot « user » représentera le nom de l’utilisateur ordinaire.

En tête de ligne de commande, on trouve : « user@localhost:~$ » ; « ~ » signifie que le working directory est : répertoire personnel de user (en principe : /home/user) et « $ » signifie que le current logged user est « user », c’est-à-dire un utilisateur ordinaire et non pas l’administrateur du système, c’est-à-dire le superutilisateur, nommé « root ».



Affichez le chemin complet du current working directory

pwd



Affichez sous quel compte vous êtes connecté

whoami



Affichez l’identification système de votre compte

id



Affichez la liste des utilisateurs connectés

who



Déconnectez-vous

exit



Basculez en mode graphique

ALT+F7



Basculez en mode texte (dans tty3)

CTRL+ALT+F3



Connectez-vous (Username + Password)



Affichez le manuel (commande Unix) de la commande clear

man clear

…

q

« q » permet de terminer cet affichage.



Affichez le manuel (commande Gnu) de la commande clear

info clear



Effacez l’écran du terminal

clear



Affichez le manuel de la commande whatis

man whatis



Affichez une courte description de la page de manuel de clear

whatis clear

man -f clear



Affichez le manuel des commandes hostname et uname

man hostname

man uname



Affichez le nom d’hôte de la machine locale

hostname



Affichez toutes les informations sur le système (version du noyau, etc.)

uname -a

Version IMPAIRE (par exemple 2.5) = version de noyau INSTABLE
Version PAIRE (par exemple 2.6) = version de noyau STABLE



Affichez le manuel des commandes date et cal

man date

man cal



Affichez la date et l’heure système

date



Affichez le calendrier courant

cal



Affichez le manuel de la commande man

man man

Les pages de manuel sont réparties dans huit [neuf] sections.

       Le tableau ci-dessous indique le numéro des sections de manuel ainsi que le type de pages qu’elles contiennent.

       1   Programmes exécutables ou commandes de l’interpréteur de commandes (shell) ;
       2   Appels système (Fonctions fournies par le noyau) ;
       3   Appels de bibliothèque (fonctions fournies par les bibliothèques des programmes) ;
       4   Fichiers spéciaux (situés généralement dans /dev) ;
       5   Formats des fichiers et conventions. Par exemple /etc/passwd ;
       6   Jeux ;
       7   Divers (y compris les macropaquets et les conventions). Par exemple, man(7), groff(7) ;
       8   Commandes de gestion du système (généralement réservées au superutilisateur) ;
       9   Sous-programmes du noyau [hors standard].



Affichez la page de manuel concernant le fichier passwd

man 5 passwd



Affichez le manuel de la commande passwd

man 1 passwd

man passwd



Déconnectez-vous

exit

6 Connexions, outils et accès aux manuels – 2

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Pour rappel, les pages de manuel sont réparties dans huit [neuf] sections.

       Le tableau ci-dessous indique le numéro des sections de manuel ainsi que le type de pages qu’elles contiennent.

       1   Programmes exécutables ou commandes de l’interpréteur de commandes (shell) ;
       2   Appels système (Fonctions fournies par le noyau) ;
       3   Appels de bibliothèque (fonctions fournies par les bibliothèques des programmes) ;
       4   Fichiers spéciaux (situés généralement dans /dev) ;
       5   Formats des fichiers et conventions. Par exemple /etc/passwd ;
       6   Jeux ;
       7   Divers (y compris les macropaquets et les conventions). Par exemple, man(7), groff(7) ;
       8   Commandes de gestion du système (généralement réservées au superutilisateur) ;
       9   Sous-programmes du noyau [hors standard].



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.



Quelques commandes :



Chaque commande est annoncée avec son numéro de section de manuel.



USERADD(8)

NOM
       useradd - créer un nouvel utilisateur ou modifier les informations par défaut
       appliquées aux nouveaux utilisateurs

SYNOPSIS
       useradd [options] LOGIN
       useradd -D
       useradd -D [options]

DESCRIPTION
       useradd is a low level utility for adding users. On Debian, administrators should
       usually use adduser(8) instead.



FORK(2)

NAME
       fork - create a child process



APROPOS(1)

NOM
       apropos - Chercher le nom et la description des pages de manuel

SYNOPSIS
       apropos [-dalv?V] [-e|-w|-r] [-s liste] [-m système[,...]] [-M chemin] [-L locale]
       [-C fichier] mot-clé ...

DESCRIPTION
       Chaque page de manuel comporte une courte description. apropos recherche
       et affiche cette description pour chaque page correspondant à mot-clé.



MORE(1)

NOM
       More - Filtre lecteur de fichier

SYNOPSIS
       more [-dlfpcsu] [-num] [+/motif] [+numligne] [fichier ...]

DESCRIPTION
       more est un filtre permettant de se déplacer dans un texte, écran par écran.
       Cette version est particulièrement primitive. less(1) constitue une excellente émulation
       de more(1) avec en plus de nombreuses améliorations.



HIER(7)

NOM
       hier - Description de la hiérarchie du système de fichiers



CAT(1)

NOM
       cat - Concaténer des fichiers et les afficher sur la sortie standard

SYNOPSIS
       cat [OPTION]... [FICHIER]...

DESCRIPTION
       Concaténer le(s) FICHIER(s) ou l'entrée standard, et les afficher sur la sortie standard.



SU(1)

NOM
       su - Changer d'identifiant d'utilisateur ou devenir superutilisateur

SYNOPSIS
       su [options] [nom_utilisateur]

DESCRIPTION
       La commande su permet de devenir un autre utilisateur pour la durée d'une session.
       Invoqué sans nom d'utilisateur, le comportement par défaut de su est de devenir
       superutilisateur. Le paramètre optionnel - permet d'obtenir un environnement
       similaire à celui que l'utilisateur aurait obtenu lors d'une connexion directe.



CD(1)

NAME
       cd - Change the shell working directory



SCRIPT(1)

NOM
       Script - Faire une transcription d'une session d'un terminal

SYNOPSIS
       script [-a] [-c commande] [-e] [-f] [-q] [-t[=fichier]] [-V] [-h] [fichier]

DESCRIPTION
       script fait une transcription de tout ce qui est affiché sur le terminal.



MANDB(8)

NOM
       mandb - Créer ou mettre à jour les bases de données d'indexation des pages de manuel

SYNOPSIS
       mandb [-dqsucpt|-h|-V] [-C fichier] [chemin_vers_man]
       mandb [-dqsut] [-C fichier] -f nom_de_fichier ...

DESCRIPTION
       mandb est employé pour initialiser ou mettre à jour manuellement les bases de données
       d'indexation qui sont habituellement exploitées par man. Ces bases de données forment
       un « cache » du système de fichiers qui contient les pages de manuel.
       Elles contiennent l'état actuel du système de pages de manuel
       ainsi que les informations qui y sont stockées.



Un fichier :



ISSUE(5)

NOM
       issue - Message d'identification du système avant la connexion

DESCRIPTION
       Le fichier /etc/issue est un fichier texte contenant un message de bienvenue
       ou d'identification du système, il est affiché avant l'invite de connexion.
       Ce fichier peut contenir diverses séquences de caractères @car et \car,
       s'ils sont gérés par le programme de type getty utilisé sur le système.



LABORATOIRE – Connexions, outils et accès aux manuels – 2



Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous au système en mode graphique (Username + Password)



Dans le « bureau » qui se présente alors à vous (interface graphique), dans le menu « Activités », puis « Afficher les applications », vous pouvez lancer un interpréteur de commandes (Shell), c’est-à-dire un « Terminal ».



Lancez un interpréteur de commandes (Shell), c’est-à-dire un « Terminal »



Affichez la page de manuel concernant la commande d’administration useradd

man 8 useradd

man useradd



Affichez la page de manuel concernant l’appel système fork

man 2 fork

man fork



Lancez un navigateur et, dans un moteur de recherche sur Internet, recherchez une page de manuel concernant l’appel système fork



« TAB » représente la touche de tabulation sur le clavier.

Affichez tous les noms de commande commençant par « user »

user TAB TAB



Recherchez la description courte des pages de manuel correspondant au mot-clé « user »

apropos user

man -k user



Affichez le manuel de la commande more

man more



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.

Affichez les pages de manuel associées à « user » en filtrant

man -k user | more



Affichez la description de l’arborescence du système de fichiers

man hier



Affichez le manuel de la commande cat

man cat



Affichez la distribution du système et sa version

cat /etc/issue



Rééditez successivement les commandes déjà exécutées dans votre terminal

Flèche vers le Haut / Flèche vers le Bas (touches sur le clavier)



Affichez le manuel de la commande su

man su



Connectez-vous sous le compte « root », c’est-à-dire le superutilisateur (Password du root)

su

Vous avez maintenant tous les droits d’administration. Vous avez la responsabilité d’en faire usage avec prudence !



Affichez l’identification système du compte root

id



Affichez la liste des utilisateurs connectés

who



Affichez le chemin complet du current working directory

pwd



Essayez d’afficher une page de manuel concernant la commande cd

man cd



Lancez un navigateur et, dans un moteur de recherche sur Internet, recherchez une page de manuel concernant la commande cd



Changez de working directory de façon à ce qu’il devienne le répertoire personnel de root

cd

cd ~

En tête de ligne de commande, on trouve : « localhost:~# » ; « ~ » signifie que le working directory est « /root » et « # » signifie que le current logged user est « root ».



Affichez le chemin complet du current working directory

pwd



Déconnectez-vous du compte root

exit



Affichez sous quel compte vous êtes connecté

whoami



Affichez le chemin complet du current working directory

pwd



« . » représente le répertoire de travail courant.
« .. » représente le répertoire parent.

Changez de working directory de façon à remonter d’un niveau dans l’arborescence

cd ..



Affichez le chemin complet du current working directory



Exécutez une commande afin que le working directory devienne « /usr/bin »

cd /usr/bin



Affichez le chemin complet du current working directory



Changez de working directory de façon à remonter à la racine de l’arborescence

cd /



Affichez le chemin complet du current working directory



Affichez le fichier /etc/passwd en filtrant

more /etc/passwd



Changez de working directory de façon à ce qu’il devienne le répertoire personnel de user

cd

cd ~



Affichez le manuel de la commande script

man script



Démarrez l’enregistrement des commandes qui seront exécutées ensuite

script -a



Affichez sous quel compte vous êtes connecté



Affichez l’identification système de votre compte



Affichez le chemin complet du current working directory



Affichez la date et l’heure système



Mettez fin à l’enregistrement

exit



Visualisez la partie de session enregistrée

more typescript



Affichez le manuel de la commande mandb

man mandb



Connectez-vous sous le compte root (Password du root)

su –



Affichez le chemin complet du current working directory

pwd

Avez-vous remarqué l’effet du « – » dans la commande « su – » ?



(Re)créez la base de données whatis (sortie verbeuse par défaut)

mandb



Déconnectez-vous du compte root

exit



Déconnectez-vous du terminal

exit






7 Accès au système de fichiers – 1

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Dans un système Linux, tous les répertoires et tous les fichiers prennent place logiquement dans une arborescence unique qui débute à la « racine » (« root »), notée : « / ».

La commande « man hier » affiche la description de la hiérarchie de ce système de fichiers.

Le chemin absolu d’un fichier est sa localisation logique dans l’arborescence, exprimée à partir de la racine. Il débute toujours par « / ». Par exemple : /var/log/messages

Un chemin relatif de fichier est sa localisation logique dans l’arborescence, exprimée par rapport au répertoire de travail courant. Par exemple : log/messages, si le répertoire de travail courant est : /var

Dans l’énoncé d’un chemin, « . » représente le répertoire de travail courant et « .. » représente le répertoire parent (du niveau supérieur dans l’arborescence).

« ~ » représente toujours le répertoire de connexion de l’utilisateur courant.

Tout « fichier » (au sens large) possède les caractéristiques suivantes :

Champ 1 : Type

                        Fichier ordinaire (-)
                        Répertoire (d)
                        Fichier spécial de périphérique (b – bloc ; c – caractère)
                        Lien symbolique (l)
                        Fichier IPC, pipe et socket (p ; s)

Champ 2 : Droits (Permissions)
Champ 3 : Nombre de liens
Champ 4 : Propriétaire
Champ 5 : Groupe
Champ 6 : Taille
Champ 7 : Date et heure de dernière modification
Champ 8 : Nom et extension



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.



Syntaxe du shell : les « jokers » (caractères de remplacement) :

Le caractère « ? » représente un caractère quelconque.
Le caractère « * » représente une chaîne quelconque de caractères, même vide.
[abc] représente a, b ou c.
[!abc] représente un caractère différent de a, b, c.
[a-e] représente un caractère de a à e.
[!a-e] représente un caractère différent des caractères de a à e.



Quelques commandes :



HIER(7)

NOM
       hier - Description de la hiérarchie du système de fichiers



CD(1)

NAME
       cd - Change the shell working directory



LS(1)

NOM
       ls - Afficher le contenu de répertoires

SYNOPSIS
       ls [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher les informations des FICHIERs (du répertoire courant par défaut).
       Les entrées sont triées alphabétiquement si aucune des options
       -cftuvSUX ou --sort n'est indiquée.



MORE(1)

NOM
       More - Filtre lecteur de fichier

SYNOPSIS
       more [-dlfpcsu] [-num] [+/motif] [+numligne] [fichier ...]

DESCRIPTION
       more est un filtre permettant de se déplacer dans un texte, écran par écran.
       Cette version est particulièrement primitive. less(1) constitue une excellente émulation
       de more(1) avec en plus de nombreuses améliorations.



CAT(1)

NOM
       cat - Concaténer des fichiers et les afficher sur la sortie standard

SYNOPSIS
       cat [OPTION]... [FICHIER]...

DESCRIPTION
       Concaténer le(s) FICHIER(s) ou l'entrée standard, et les afficher sur la sortie standard.



LESS(1)

NAME
       less - opposite of more

SYNOPSIS
       less -?
       less --help
       less -V
       less --version
       less [-[+]aABcCdeEfFgGiIJKLmMnNqQrRsSuUVwWX~]
              [-b space] [-h lines] [-j line] [-k keyfile]
              [-{oO} logfile] [-p pattern] [-P prompt] [-t tag]
              [-T tagsfile] [-x tab,...] [-y lines] [-[z] lines]
              [-# shift] [+[+]cmd] [--] [filename]...
       (See the OPTIONS section for alternate option syntax with long option names.)

DESCRIPTION
       less is a program similar to more (1), but it has many more features.



Un fichier :



« /etc/services » : fichier qui contient la liste générale des services réseau TCP/IP avec leur numéro de port et le protocole de transport associé



LABORATOIRE – Accès au système de fichiers – 1

Les exercices des séquences 7 à 11 s’enchaînent.
Il faut donc les accomplir dans l’ordre prévu.

Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous dans un terminal



Affichez la description de l’arborescence du système de fichiers

man hier



Un chemin absolu ou complet commence par « / », c’est-à-dire par la racine.
Un chemin relatif commence par le répertoire ou le fichier adéquat, dans le répertoire de travail courant.

Faites en sorte que le working directory devienne « /usr/bin » (emploi du chemin absolu)

cd /usr/bin



Changez de working directory de façon à remonter à la racine de l’arborescence

cd /



Faites en sorte que le working directory devienne « /usr/bin » (emploi du chemin relatif)

cd usr/bin



Changez de working directory de façon à remonter d’un niveau dans l’arborescence

cd ..



Faites en sorte que le working directory devienne « /usr/bin » (emploi du chemin relatif)

cd bin



Faites en sorte que le working directory devienne « /var/log » (emploi du chemin relatif)

cd ../../var/log



Faites en sorte que le working directory devienne « /var/lib » (emploi du chemin relatif)

cd ../lib



Faites en sorte que le working directory devienne « /bin » (emploi du chemin absolu)

cd /bin



Affichez le manuel de la commande ls

man ls



Affichez le contenu du working directory

ls



Affichez le contenu du working directory, en utilisant un format d’affichage long

ls -l

Champ 1 : Type

                        Fichier ordinaire (-)
                        Répertoire (d)
                        Fichier spécial de périphérique (b – bloc ; c – caractère)
                        Lien symbolique (l)
                        Fichier IPC, pipe et socket (p ; s)

Champ 2 : Droits (Permissions)
Champ 3 : Nombre de liens
Champ 4 : Propriétaire
Champ 5 : Groupe
Champ 6 : Taille
Champ 7 : Date et heure de dernière modification
Champ 8 : Nom et extension



Affichez le manuel des commandes more, cat et less

man more

man cat

man less



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.

Affichez le contenu du working directory, en utilisant un format d’affichage long et en filtrant

ls -l | more



Syntaxe du shell : les « jokers » (caractères de remplacement) :

Le caractère « ? » représente un caractère quelconque.
Le caractère « * » représente une chaîne quelconque de caractères, même vide.
[abc] représente a, b ou c.
[!abc] représente un caractère différent de a, b, c.
[a-e] représente un caractère de a à e.
[!a-e] représente un caractère différent des caractères de a à e.

Affichez, en utilisant un format d’affichage long, les fichiers dans le répertoire /usr/bin dont le nom commence par un caractère quelconque suivi de la chaîne de caractères « sh », puis par deux caractères quelconques suivis du caractère « h »

ls -l /usr/bin/?sh

ls -l /usr/bin/??h



Affichez les fichiers dans le working directory dont le nom est formé de cinq caractères, en utilisant un format d’affichage long

ls -l ?????



Affichez les fichiers dans le working directory commençant par la lettre « c », en utilisant un format d’affichage long

ls -l c*



Affichez les fichiers dans le working directory finissant par « h » et dont le nom est formé de trois caractères, en utilisant un format d’affichage long

ls -l ??h



Affichez les fichiers dans le working directory commençant par la lettre « a » ou « d », en utilisant un format d’affichage long

ls -l [ad]*



Affichez les fichiers dans le working directory commençant par un caractère différent de « a » et « c », en utilisant un format d’affichage long

ls -l [!ac]*



Affichez les fichiers dans le working directory commençant par un caractère de « k » à « p », en utilisant un format d’affichage long

ls -l [k-p]*



Affichez les fichiers dans le working directory commençant par un caractère différent des caractères de « c » à « v », en utilisant un format d’affichage long

ls -l [!c-v]*



« ~ » représente le répertoire de connexion de l’utilisateur courant.

Affichez le contenu du répertoire /home/user, en utilisant un format d’affichage long

ls -l /home/user

ls -l ~



Affichez le contenu du répertoire /home/user, y compris les entrées cachées (débutant par « . »), en utilisant un format d’affichage long

ls -al /home/user

ls -al ~



Affichez le contenu du répertoire /var/log classé par date, c'est-à-dire du plus récent au plus ancien, en utilisant un format d’affichage long

ls -lt /var/log



Affichez les fichiers dans le working directory en ajoutant un caractère (parmi « */=>@| ») à chaque entrée de façon à les classer selon leur type, en utilisant un format d’affichage long

ls -lF



Affichez le contenu du répertoire /var/log classé par type d’extension, en utilisant un format d’affichage long

ls -lX /var/log



Affichez les fichiers dans le working directory classés par taille décroissante, en utilisant un format d’affichage long

ls -lS



Visualisez les caractéristiques du fichier /etc/services

ls -l /etc/services



Visualisez les caractéristiques du fichier de commande /usr/bin/who

ls -l /usr/bin/who



Déconnectez-vous

exit



8 Accès au système de fichiers – 2

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Syntaxe du shell : les caractères de protection ou d’échappement :

\        Annulation de la signification du caractère spécial suivant
'...'        Annulation de la signification de tous les caractères spéciaux
"..."        Annulation de la signification des caractères spéciaux, sauf : « ` », « \ » et « $ »



Quelques commandes :



FILE(1)

NAME
       file - determine file type

C’est grâce à la commande « file » que l’on peut savoir si un fichier est un fichier de programme exécutable, ou un fichier de texte, ou etc.



CD(1)

NAME
       cd - Change the shell working directory



MKDIR(1)

NOM
       mkdir - Créer des répertoires

SYNOPSIS
       mkdir [OPTION]... RÉPERTOIRE...

DESCRIPTION
       Créer les RÉPERTOIREs s'ils n'existent pas.



LS(1)

NOM
       ls - Afficher le contenu de répertoires

SYNOPSIS
       ls [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher les informations des FICHIERs (du répertoire courant par défaut).
       Les entrées sont triées alphabétiquement si aucune des options
       -cftuvSUX ou --sort n'est indiquée.



CP(1)

NOM
       cp - Copier des fichiers et des répertoires

SYNOPSIS
       cp [OPTION]... [-T] SOURCE CIBLE
       cp [OPTION]... SOURCE... RÉPERTOIRE
       cp [OPTION]... -t RÉPERTOIRE SOURCE...

DESCRIPTION
       Copier la SOURCE vers la CIBLE, ou de multiples SOURCEs vers le RÉPERTOIRE.



RM(1)

NOM
       rm - Effacer des fichiers et des répertoires

SYNOPSIS
       rm [OPTION]... FICHIER...

DESCRIPTION
       Le programme rm efface chaque fichier listé.
       Par défaut, il n'efface pas les répertoires.



DATE(1)

NOM
       date - Afficher ou configurer la date et l'heure du système

SYNOPSIS
       date [OPTION]... [+FORMAT]
       date [-u|--utc|--universal] [MMJJhhmm[[CC]AA][.ss]]

DESCRIPTION
       Afficher dans le FORMAT indiqué ou configurer la date et l'heure du système.



LABORATOIRE – Accès au système de fichiers – 2

Les exercices des séquences 7 à 11 s’enchaînent.
Il faut donc les accomplir dans l’ordre prévu.

Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous dans un terminal



Affichez le manuel de la commande file

man file



Visualisez le type des fichiers /etc/services, /etc/issue, /usr/bin/who, /etc/passwd

file /etc/services

file /etc/issue

file /usr/bin/who

file /etc/passwd



Changez de working directory de façon à ce qu’il devienne le répertoire personnel de user

cd

cd ~



Affichez le manuel de la commande mkdir

man mkdir



Créez un répertoire nommé : « Un_dossier »

mkdir Un_dossier



Listez



Syntaxe du shell : les caractères de protection ou d’échappement :

\        Annulation de la signification du caractère spécial suivant
'...'        Annulation de la signification de tous les caractères spéciaux
"..."        Annulation de la signification des caractères spéciaux, sauf : « ` », « \ » et « $ »

Créez un répertoire nommé : « Un dossier »

mkdir 'Un dossier'



Listez



Créez la suite de répertoires suivante : dir1/doc1

mkdir dir1

ls

cd dir1

mkdir doc1

ls



Changez de working directory de façon à ce qu’il devienne le répertoire personnel de user

cd

cd ~



Créez la suite de répertoires dir2/doc2/textes en exécutant une seule commande

mkdir -p dir2/doc2/textes



Visualisez le résultat puis faites en sorte que le working directory redevienne le répertoire personnel de user



Affichez le manuel de la commande cp

man cp



Copiez le fichier /etc/issue, dans ~

cp /etc/issue ~



Listez



Copiez le fichier /etc/issue en le fichier nommé : « distrib-vers », dans ~

cp /etc/issue distrib-vers



Listez



Copiez le fichier /etc/issue en le fichier distrib-vers, dans Un_dossier (chemin relatif)

cp /etc/issue Un_dossier/distrib-vers

cp /etc/issue ./Un_dossier/distrib-vers



Copiez le fichier /etc/issue en le fichier distrib-vers, dans Un_dossier (chemin absolu)

cp /etc/issue /home/user/Un_dossier/distrib-vers

cp /etc/issue ~/Un_dossier/distrib-vers

ATTENTION !        Si vous effectuez une copie d’un fichier sur un fichier qui existe dejà, celui-ci est effacé et remplacé par le nouveau fichier !



Copiez le fichier ~/issue en le fichier distrib-vers, dans Un_dossier, avec avertissement

cp -i issue Un_dossier/distrib-vers



Copiez le fichier ~/issue en le fichier distrib-vers, dans Un_dossier, avec avertissement et sauvegarde du fichier précédent dans le répertoire de destination

cp -ib issue Un_dossier/distrib-vers



Affichez le contenu du répertoire /home/user/Un_dossier, y compris les entrées cachées (débutant par « . »), en utilisant un format d’affichage long

ls -al Un_dossier



Copiez ~/Un_dossier/distrib-vers en d-v-ldur, dans le répertoire « Un dossier », en créant un lien « dur »

cp -l Un_dossier/distrib-vers 'Un dossier'/d-v-ldur



Visualisez les caractéristiques du fichier Un_dossier/distrib-vers

ls -l Un_dossier/distrib-vers



Visualisez les caractéristiques du fichier d-v-ldur

ls -l 'Un dossier'/d-v-ldur



Copiez ~/Un_dossier/distrib-vers en d-v-lsymb, dans le répertoire « Un dossier », en créant un lien symbolique

cd 'Un dossier'
cp -s ../Un_dossier/distrib-vers d-v-lsymb
cd ..



Visualisez les caractéristiques du fichier Un_dossier/distrib-vers

ls -l Un_dossier/distrib-vers



Visualisez les caractéristiques du fichier d-v-lsymb

ls -l 'Un dossier'/d-v-lsymb



Copiez récursivement le répertoire Un_dossier et son contenu dans le répertoire dir1, avec affichage des fichiers copiés (sortie verbeuse)

cp -rv Un_dossier dir1



Visualisez le résultat en listant



Affichez le manuel de la commande rm

man rm



Supprimez le fichier ~/distrib-vers, avec avertissement

rm -i distrib-vers



Supprimez récursivement et de manière forcée le répertoire ~/dir1/Un_dossier et son contenu, sans avertissement

rm -rf dir1/Un_dossier

Cette dernière opération n’est évidemment à exécuter qu’avec grande précaution !



Affichez la date et l’heure système, dans le style :

DATE : samedi 24 mars 2029
HEURE : 16 heures 25 minutes 44 secondes

date '+DATE : %A %d %B %C%y%nHEURE : %H heures %M minutes %S secondes'



Déconnectez-vous

exit

9 Accès au système de fichiers – 3

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Syntaxe du shell :

Annonce d’un commentaire                        #

Séparation de commandes                        ; ou un saut de ligne ou une tabulation

Evaluation d’une variable                        $



Syntaxe du shell : les redirections de fichier :

        <        Redirection en entrée
        >        Redirection en sortie (création ou réécriture)
        >>        Redirection en sortie (création ou ajout)



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.



Quelques commandes :



MV(1)

NOM
       mv - Déplacer ou renommer des fichiers

SYNOPSIS
       mv [OPTION]... [-T] SOURCE CIBLE
       mv [OPTION]... SOURCE... RÉPERTOIRE
       mv [OPTION]... -t RÉPERTOIRE SOURCE...

DESCRIPTION
       Renommer la SOURCE en CIBLE ou déplacer la SOURCE vers le RÉPERTOIRE.



CP(1)

NOM
       cp - Copier des fichiers et des répertoires

SYNOPSIS
       cp [OPTION]... [-T] SOURCE CIBLE
       cp [OPTION]... SOURCE... RÉPERTOIRE
       cp [OPTION]... -t RÉPERTOIRE SOURCE...

DESCRIPTION
       Copier la SOURCE vers la CIBLE, ou de multiples SOURCEs vers le RÉPERTOIRE.



ECHO(1)

NOM
       echo - Afficher une ligne de texte

SYNOPSIS
       echo [OPTION-COURTE]... [CHAÎNE]...
       echo OPTION-LONGUE

DESCRIPTION
       Afficher la(les) CHAÎNE(s) en écho sur la sortie standard.



MORE(1)

NOM
       More - Filtre lecteur de fichier

SYNOPSIS
       more [-dlfpcsu] [-num] [+/motif] [+numligne] [fichier ...]

DESCRIPTION
       more est un filtre permettant de se déplacer dans un texte, écran par écran.
       Cette version est particulièrement primitive. less(1) constitue une excellente émulation
       de more(1) avec en plus de nombreuses améliorations.



CAT(1)

NOM
       cat - Concaténer des fichiers et les afficher sur la sortie standard

SYNOPSIS
       cat [OPTION]... [FICHIER]...

DESCRIPTION
       Concaténer le(s) FICHIER(s) ou l'entrée standard, et les afficher sur la sortie standard.



LESS(1)

NAME
       less - opposite of more

SYNOPSIS
       less -?
       less --help
       less -V
       less --version
       less [-[+]aABcCdeEfFgGiIJKLmMnNqQrRsSuUVwWX~]
              [-b space] [-h lines] [-j line] [-k keyfile]
              [-{oO} logfile] [-p pattern] [-P prompt] [-t tag]
              [-T tagsfile] [-x tab,...] [-y lines] [-[z] lines]
              [-# shift] [+[+]cmd] [--] [filename]...
       (See the OPTIONS section for alternate option syntax with long option names.)

DESCRIPTION
       less is a program similar to more (1), but it has many more features.



HEAD(1)

NOM
       head - Afficher le début des fichiers

SYNOPSIS
       head [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher par défaut les 10 premières lignes de chaque FICHIER sur la sortie standard.
       Avec plus d'un FICHIER, faire précéder chacun d'un en-tête donnant le nom du fichier.
       L'entrée standard est lue quand FICHIER est omis ou quand FICHIER vaut « - ».



TAIL(1)

NOM
       tail - Afficher la dernière partie de fichiers

SYNOPSIS
       tail [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher par défaut les 10 dernières lignes de chaque FICHIER sur la sortie standard.
       Lorsqu'il y a plus d'un FICHIER, faire précéder chaque groupe de lignes
       d'un en-tête donnant le nom du fichier. L'entrée standard est lue
       quand FICHIER est omis ou quand FICHIER vaut « - ».



WC(1)

NOM
       wc - Afficher le nombre de lignes, de mots et d'octets d'un fichier

SYNOPSIS
       wc [OPTION] ... [FICHIER] ...
       wc [OPTION] ... --files0-from=FICHIER

DESCRIPTION
       Afficher le décompte de lignes, de mots et d'octets pour chaque FICHIER
       et le nombre total de lignes si plus d'un FICHIER est indiqué.
       L'entrée standard est lue quand FICHIER est omis ou quand FICHIER vaut « - ».
       Un mot est une séquence non vide de caractères délimités par des espaces.
       Les options permettent de choisir quels décomptes sont affichés.
       Ils le sont toujours dans l'ordre suivant : lignes, mots, caractères,
       octets, longueur maximale des lignes.



TOUCH(1)

NOM
       touch - Modifier l'horodatage d'un fichier

SYNOPSIS
       touch [OPTION]... FICHIER...

DESCRIPTION
       Mettre à jour les dates (dates et heures) d'accès et de dernière modification
       de chaque FICHIER selon la date actuelle. Un paramètre FICHIER qui n'existe pas
       sera créé vide, à moins que les options -c ou -h ne soient fournies.



LABORATOIRE – Accès au système de fichiers – 3

Les exercices des séquences 7 à 11 s’enchaînent.
Il faut donc les accomplir dans l’ordre prévu.

Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous dans un terminal



Affichez le manuel de la commande mv

man mv



Renommez le fichier ~/issue en « distrib-vers »

mv issue distrib-vers



Listez



Renommez le répertoire ~/Un_dossier en « Un_catalogue »

mv Un_dossier Un_catalogue



Listez



Déplacez le fichier ~/distrib-vers dans le répertoire ~/dir1/doc1

mv distrib-vers dir1/doc1



Visualisez le résultat en listant



Copiez le fichier ~/dir1/doc1/distrib-vers dans le répertoire ~/dir2/doc2

cp ~/dir1/doc1/distrib-vers ~/dir2/doc2



Déplacez le fichier ~/Un_catalogue/distrib-vers dans le répertoire ~/dir2/doc2, avec avertissement

mv -i Un_catalogue/distrib-vers dir2/doc2



Visualisez le résultat en listant



Copiez le fichier ~/dir1/doc1/distrib-vers en le fichier nommé : « d-v-2 », dans le répertoire ~/dir2/doc2

cp ~/dir1/doc1/distrib-vers ~/dir2/doc2/d-v-2



Renommez le fichier ~/dir2/doc2/distrib-vers en « d-v-2 », avec avertissement et sauvegarde du fichier précédent

mv -ib dir2/doc2/distrib-vers dir2/doc2/d-v-2



Visualisez le résultat en listant



Affichez le manuel de la commande echo

man echo



Syntaxe du shell : les redirections de fichier :

        <        Redirection en entrée
        >        Redirection en sortie (création ou réécriture)
        >>        Redirection en sortie (création ou ajout)

Créez le fichier de texte ~/dir2/doc2/textes/F_texte contenant la chaîne de caractères « Fichier de texte » en utilisant la commande echo et une redirection

echo "Fichier de texte" > dir2/doc2/textes/F_texte



Affichez le manuel des commandes more, cat et less

man more

man cat

man less



Affichez le contenu du fichier ~/dir2/doc2/textes/F_texte

cat dir2/doc2/textes/F_texte



Créez le fichier de texte ~/dir2/doc2/textes/F_texte2 contenant la chaîne de caractères « Fichier de texte numéro 2 » en utilisant la commande echo et une redirection

echo "Fichier de texte numéro 2" > dir2/doc2/textes/F_texte2



Affichez le contenu du fichier ~/dir2/doc2/textes/F_texte2

cat dir2/doc2/textes/F_texte2



Listez afin de bien visualiser quel est le plus récent des fichiers de texte

ls -l dir2/doc2/textes



Essayez de renommer le plus ancien en lui donnant le nom du plus récent, avec demande de ne pas remplacer un fichier plus récent par un plus ancien

mv -u dir2/doc2/textes/F_texte dir2/doc2/textes/F_texte2



Redirigez le calendrier courant dans un fichier ~/calendrier.txt

cal > calendrier.txt



Affichez le contenu du fichier

cat calendrier.txt



Ajoutez-y la date et l’heure

date >> calendrier.txt



Affichez le contenu du fichier

cat calendrier.txt



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.

Affichez le calendrier 2008 page par page

cal 2008 | more



Affichez le manuel des commandes head, tail et wc

man head

man tail

man wc



Affichez les premières lignes du fichier /etc/services

head /etc/services



Affichez les 7 premières lignes du fichier /etc/services

head -7 /etc/services



Affichez les dernières lignes du fichier /etc/services

tail /etc/services



Affichez les 5 dernières lignes du fichier /etc/services

tail -5 /etc/services



Affichez le nombre de lignes, de mots et de caractères du fichier /etc/services

wc /etc/services



Affichez seulement le nombre de lignes du fichier /etc/services

wc -l /etc/services



Affichez seulement le nombre de mots du fichier /etc/issue

wc -w /etc/issue



Affichez seulement le nombre d’octets ou de caractères du fichier /etc/passwd

wc -c /etc/passwd

wc -m /etc/passwd



Affichez le manuel de la commande touch

man touch



Syntaxe du shell :

Annonce d’un commentaire                        #

Séparation de commandes                        ; ou un saut de ligne ou une tabulation

Evaluation d’une variable                        $

Avec touch, créez dans ~ un fichier nommé : « #Commentaire »

touch \#Commentaire

# est un caractère spécial ; sa signification a été annulée grâce à l’utilisation de \



Listez



Avec touch, créez dans ~ un fichier nommé : « Les roses sont bleues », puis listez, en une seule ligne de commande

touch 'Les roses sont bleues' ; ls -l



Déconnectez-vous

exit

10 Accès au système de fichiers – 4

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Syntaxe du shell :

Annonce d’un commentaire                        #

Séparation de commandes                        ; ou un saut de ligne ou une tabulation

Evaluation d’une variable                        $



Toute application (tout processus) possède un espace mémoire dit d’environnement, contenant des variables d’environnement ayant chacune un nom et une valeur sous forme d’une chaîne de caractères. Cet espace est hérité de processus père en processus fils.

Quelques variables d’environnement du shell :

PS1                Prompt
HOME                Chemin absolu du répertoire de connexion
PATH                Chemins des répertoires où sont recherchées les commandes (séparés par « : »)



Syntaxe du shell : les caractères de protection ou d’échappement :

\        Annulation de la signification du caractère spécial suivant
'...'        Annulation de la signification de tous les caractères spéciaux
"..."        Annulation de la signification des caractères spéciaux, sauf : « ` », « \ » et « $ »



Syntaxe du shell : les redirections de fichier :

        <        Redirection en entrée
        >        Redirection en sortie (création ou réécriture)
        >>        Redirection en sortie (création ou ajout)



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.



Syntaxe du shell : les substitutions de commande :

`commande`                        la commande est interprétée
$(commande)                        la commande est interprétée



! ! !   Caractères différents utilisés :        '  quote        "   double quote        `   antiquote   ! ! !



Quelques commandes :



ECHO(1)

NOM
       echo - Afficher une ligne de texte

SYNOPSIS
       echo [OPTION-COURTE]... [CHAÎNE]...
       echo OPTION-LONGUE

DESCRIPTION
       Afficher la(les) CHAÎNE(s) en écho sur la sortie standard.



IFCONFIG(8)

NOM
       ifconfig - configure une interface réseau

SYNOPSIS
       ifconfig [interface]
       ifconfig interface [aftype] options | adresse ...

DESCRIPTION
       ifconfig est utilisé pour configurer (et maintenir ensuite) les interfaces réseau
       résidentes dans le noyau. Si aucun argument n'est donné, ifconfig affiche
       simplement l'état des interfaces actuellement définies. Si seul le paramètre interface
       est donné, il affiche seulement l'état de l'interface correspondante. Si seul
       le paramètre -a est fourni, il affiche l'état de toutes les interfaces, même celles
       qui ne sont pas actives.



BASH(1)

NOM
       bash - Interpréteur de commandes GNU Bourne-Again SHell

SYNOPSIS
       bash [options] [chaîne_de_commande | fichier]

DESCRIPTION
       bash est un interpréteur de commandes (shell) compatible sh qui exécute les
       commandes lues depuis l'entrée standard ou depuis un fichier. bash inclut aussi
       des fonctionnalités utiles des interpréteurs de commandes ksh et csh.



GREP(1)

NOM
       grep, egrep, fgrep, rgrep - Afficher les lignes correspondant à un motif donné

SYNOPSIS
       grep [OPTIONS] MOTIF [FICHIER...]
       grep [OPTIONS] [-e MOTIF | -f FICHIER] [FICHIER...]

DESCRIPTION
       grep recherche dans les FICHIERs indiqués les lignes correspondant à un certain
       MOTIF. Par défaut, grep affiche les lignes qui contiennent une correspondance
       au motif. L'entrée standard est lue si FICHIER est omis ou si FICHIER vaut « - ».

       Trois variantes du programme sont disponibles : egrep, fgrep et rgrep ; egrep
       est identique à grep -E, fgrep est identique à grep -F et rgrep est identique à
       grep -r. L'appel direct à egrep ou fgrep est déconseillé, mais est toujours possible pour
       permettre à d'anciennes applications qui les utilisent de fonctionner sans modification.



Quelques fichiers :



« ~/.profile » (ou « ~/.bash_profile », selon les versions) :
        fichier de script de démarrage du shell bash spécifique d’un utilisateur

« /etc/profile » :
        fichier de script de démarrage commun à l’ensemble des utilisateurs



LABORATOIRE – Accès au système de fichiers – 4

Les exercices des séquences 7 à 11 s’enchaînent.
Il faut donc les accomplir dans l’ordre prévu.

Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous dans un terminal



« ~ » représente le répertoire de connexion de l’utilisateur courant.

Avec echo, affichez le répertoire de connexion

echo ~



Variables d’environnement :

PS1                Prompt
HOME                Chemin absolu du répertoire de connexion
PATH                Chemins des répertoires où sont recherchées les commandes (séparés par « : »)

En utilisant la variable PS1, affichez le prompt

echo $PS1



En utilisant la variable HOME, affichez le répertoire de connexion

echo $HOME



En utilisant la variable HOME, affichez : « Voici mon répertoire de connexion : /home/user »

echo "Voici mon répertoire de connexion : $HOME"



En utilisant la variable HOME, ajoutez dans le fichier ~/#Commentaire : « Voici mon répertoire de connexion : /home/user »

echo "Voici mon répertoire de connexion : $HOME" >> \#Commentaire

# est un caractère spécial ; sa signification a été annulée grâce à l’utilisation de \



Affichez le contenu du fichier ~/#Commentaire

cat \#Commentaire



Avec echo, affichez la phrase : « Voici la date et l’heure :  », suivie de la date et de l’heure

echo "Voici la date et l’heure : " ; date



Syntaxe du shell : les substitutions de commande :

`commande`                        la commande est interprétée
$(commande)                        la commande est interprétée

"..."        Annulation de la signification des caractères spéciaux, sauf : « ` », « \ » et « $ »

Avec echo, affichez la phrase : « Voici la date et l’heure :  », suivie de la date et de l’heure passées en arguments par la commande date à la commande echo

echo "Voici la date et l’heure : `date`"

On a ainsi généré dynamiquement des arguments pour la commande echo.



Affichez les chemins des répertoires où sont recherchées les commandes

echo $PATH



Essayez d’afficher la version de la commande d’administration ifconfig

ifconfig --version



Recommencez en précisant le chemin absolu de cette commande

/sbin/ifconfig --version



Faites en sorte que le working directory devienne « /sbin »

cd /sbin



Essayez d’afficher la version de la commande d’administration ifconfig

ifconfig --version



Recommencez en précisant le chemin de cette commande

./ifconfig --version



Changez de working directory de façon à ce qu’il devienne le répertoire personnel de user

cd

cd ~



Modifiez transitoirement la valeur de la variable PATH de façon à ce qu’elle comprenne aussi le chemin « /sbin »

PATH=$PATH:/sbin

Une modification permanente de la valeur de PATH nécessite une mise à jour du fichier système : « ~/.profile » (ou « ~/.bash_profile », selon les versions).



Affichez les chemins des répertoires où sont recherchées les commandes

echo $PATH



Affichez la version de la commande d’administration ifconfig

ifconfig --version



Affichez le manuel de la commande bash

man bash



Lancez un shell bash fils

bash



Affichez les chemins des répertoires où sont recherchées les commandes

echo $PATH

Toute application (tout processus) possède un espace mémoire dit d’environnement, contenant des variables d’environnement ayant chacune un nom et une valeur sous forme d’une chaîne de caractères. Cet espace est hérité de processus père en processus fils.



Déconnectez-vous

exit

Extinction du processus fils.

exit

Extinction du processus père.



Reconnectez-vous dans un terminal



Affichez les chemins des répertoires où sont recherchées les commandes

echo $PATH



Affichez le manuel de la commande grep

man grep



Affichez les lignes du fichier /etc/services qui contiennent la chaîne « HTTP »

grep HTTP /etc/services



Exécutez la même tâche en ne tenant pas compte de la casse

grep -i HTTP /etc/services



Affichez les lignes du fichier /etc/services qui contiennent la chaîne « HTTP » en les numérotant

grep -n HTTP /etc/services



Affichez le nombre d’occurrences de la chaîne « HTTP » dans le fichier /etc/services

grep -c HTTP /etc/services



Affichez les fichiers dans le répertoire /etc qui contiennent la chaîne « HTTP »

grep -l HTTP /etc/*



Affichez les lignes du fichier /etc/services qui commencent par la chaîne « http »

grep '^http' /etc/services



Affichez les lignes du fichier /etc/services qui ne contiennent pas la chaîne « HTTP »

grep -v HTTP /etc/services



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.

Exécutez la même tâche page par page

grep -v HTTP /etc/services | more



Affichez page par page les lignes du fichier /etc/services qui ne commencent pas par la chaîne « http »

grep -v '^http' /etc/services | more



Créez dans ~ trois fichiers nommés : « f1.txt », « f2.txt » et « f3.txt », contenant chacun quelques lignes de texte non triées et contenant plusieurs fois la chaîne « html »



Affichez la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

cat f1.txt f2.txt f3.txt



Affichez la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt en numérotant les lignes non vides

cat -b f1.txt f2.txt f3.txt



Affichez la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt en numérotant toutes les lignes

cat -n f1.txt f2.txt f3.txt



Ecrivez la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt dans le fichier nommé : « f4.txt », en le créant dans ~

cat f1.txt f2.txt f3.txt > f4.txt



Déconnectez-vous

exit


11 Accès au système de fichiers – 5

Légende :

Explications sur fond blanc
Exercice à accomplir sur fond rouge clair
Une ou plusieurs solutions possibles pour accomplir l’exercice sur fond vert clair



Syntaxe du shell : les redirections de fichier :

        <        Redirection en entrée
        >        Redirection en sortie (création ou réécriture)
        >>        Redirection en sortie (création ou ajout)



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.



Un éditeur est un programme de traitement de texte, pourvu de fonctionnalités plus ou moins complexes.

« gedit » est un éditeur en mode graphique, disponible dans l’interface graphique du système, via le menu « Activités », puis « Afficher les applications ». Il est pourvu de fonctionnalités complexes pour assurer la mise en évidence, selon un mode de coloration, des éléments de nature semblable (commandes, instructions, variables, etc.) d’un texte de script ou de code source écrit dans un langage de programmation.

En mode texte, « nano » est un éditeur installé par défaut. Par contre, « emacs » et « joe » sont des éditeurs à installer.



Quelques commandes :



SORT(1)

NOM
       sort - Trier les lignes de fichiers texte

SYNOPSIS
       sort [OPTION]... [FICHIER]...
       sort [OPTION]... --files0-from=F

DESCRIPTION
       Afficher sur la sortie standard la concaténation triée de tous les FICHIERs.



HEAD(1)

NOM
       head - Afficher le début des fichiers

SYNOPSIS
       head [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher par défaut les 10 premières lignes de chaque FICHIER sur la sortie standard.
       Avec plus d'un FICHIER, faire précéder chacun d'un en-tête donnant le nom du fichier.
       L'entrée standard est lue quand FICHIER est omis ou quand FICHIER vaut « - ».



TAIL(1)

NOM
       tail - Afficher la dernière partie de fichiers

SYNOPSIS
       tail [OPTION]... [FICHIER]...

DESCRIPTION
       Afficher par défaut les 10 dernières lignes de chaque FICHIER sur la sortie standard.
       Lorsqu'il y a plus d'un FICHIER, faire précéder chaque groupe de lignes
       d'un en-tête donnant le nom du fichier. L'entrée standard est lue
       quand FICHIER est omis ou quand FICHIER vaut « - ».



GREP(1)

NOM
       grep, egrep, fgrep, rgrep - Afficher les lignes correspondant à un motif donné

SYNOPSIS
       grep [OPTIONS] MOTIF [FICHIER...]
       grep [OPTIONS] [-e MOTIF | -f FICHIER] [FICHIER...]

DESCRIPTION
       grep recherche dans les FICHIERs indiqués les lignes correspondant à un certain
       MOTIF. Par défaut, grep affiche les lignes qui contiennent une correspondance
       au motif. L'entrée standard est lue si FICHIER est omis ou si FICHIER vaut « - ».

       Trois variantes du programme sont disponibles : egrep, fgrep et rgrep ; egrep
       est identique à grep -E, fgrep est identique à grep -F et rgrep est identique à
       grep -r. L'appel direct à egrep ou fgrep est déconseillé, mais est toujours possible pour
       permettre à d'anciennes applications qui les utilisent de fonctionner sans modification.



CAT(1)

NOM
       cat - Concaténer des fichiers et les afficher sur la sortie standard

SYNOPSIS
       cat [OPTION]... [FICHIER]...

DESCRIPTION
       Concaténer le(s) FICHIER(s) ou l'entrée standard, et les afficher sur la sortie standard.



APT(8)

NOM
       apt - interface en ligne de commande

SYNOPSIS
       apt [-h] [-o=chaîne_de_configuration] [-c=fichier_de_configuration]
       [-t=publication_cible] [-a=architecture] {list | search | show | update |
       install paquet [{=numéro_version_paquet | /publication_cible}]... |
       remove paquet... | upgrade | full-upgrade | edit-sources |
       {-v | --version} | {-h | --help}}

DESCRIPTION
       apt (Advanced Package Tool) est un outil en ligne de commande pour gérer les paquets.
       Il fournit une interface en ligne de commande au système de gestion de paquets. Voir
       aussi apt-get(8) et apt-cache(8) pour davantage d'options en ligne de commande.



APT-GET(8)

NOM
       apt-get - Utilitaire APT pour la gestion des paquets - interface en ligne de commande

SYNOPSIS
       apt-get [-asqdyfmubV] [-o=chaîne_de_configuration] [-c=fichier_de_configuration]
       [-t=publication_cible] [-a=architecture] {update | upgrade | dselect-upgrade |
       dist-upgrade | install paquet [{=numéro_version_paquet | /publication_cible}]... |
       remove paquet... | purge paquet... |
       source paquet [{=numéro_version_paquet | /publication_cible}]... |
       build-dep paquet [{=numéro_version_paquet | /publication_cible}]... |
       download paquet [{=numéro_version_paquet | /publication_cible}]... |
       check | clean | autoclean | autoremove | {-v | --version} | {-h | --help}}

DESCRIPTION
       apt-get est le programme en ligne de commande pour la gestion des paquets. Il peut être
       considéré comme l'outil de base pour les autres programmes de la bibliothèque APT.
       Plusieurs interfaces utilisateur existent, comme aptitude(8), synaptic(8) et wajig(1).



APTITUDE(8)

NOM
       aptitude - interface évoluée pour le gestionnaire de paquets

SYNOPSIS
       aptitude [<options>...] {autoclean | clean | forget-new | keep-all | update}
       aptitude [<options>...] {full-upgrade | safe-upgrade} [<paquets>...]
       aptitude [<options>...] {build-dep | build-depends | changelog | download |
       forbid-version | hold | install | markauto | purge | reinstall | remove | show |
       unhold | unmarkauto | versions} <paquets>...
       aptitude extract-cache-subset <répertoire-sortie> <paquets>...
       aptitude [<options>...] search <motifs>...
       aptitude [<options>...] {add-user-tag | remove-user-tag} <étiquettes> <paquets>...
       aptitude [<options>...] {why | why-not} [<motifs>...] <paquet>
       aptitude [-S <nom-fonct>] [--autoclean-on-startup | --clean-on-startup | -i | -u]
       aptitude help

DESCRIPTION
       aptitude est une interface en mode texte pour le gestionnaire de paquets
       de Debian GNU/Linux. Elle permet à l'utilisateur de connaître la liste des paquets et
       de réaliser des tâches d'administration comme l'installation, la mise à jour ou
       la suppression de paquets. Ces tâches peuvent être réalisées en mode « interactif »
       ou à partir de la « ligne de commande ».

ACTIONS EN LIGNE DE COMMANDE
       Le premier argument qui ne commence pas par un tiret (« - ») sera considéré comme étant
       la commande que le programme doit réaliser. Si aucune commande n'est donnée,
       aptitude démarrera en mode interactif.



NANO(1)

NAME
       nano - Nano's ANOther editor, an enhanced free Pico clone

SYNOPSIS
       nano [OPTIONS] [[+LINE,COLUMN] FILE]...

DESCRIPTION
       nano is a small, free and friendly editor which aims to replace Pico, the default
       editor included in the non-free Pine package. Rather than just copying Pico's look
       and feel, nano also implements some missing (or disabled by default) features in
       Pico, such as "search and replace" and "go to line and column number".



EMACS(1)

NAME
       emacs - GNU project Emacs

SYNOPSIS
       emacs [ command-line switches ] [ files ... ]

DESCRIPTION
       GNU Emacs is a version of Emacs, written by the author of the original (PDP-10) Emacs,
       Richard Stallman. The user functionality of GNU Emacs encompasses everything
       other editors do, and it is easily extensible since its editing commands are written in Lisp.



JOE(1)

NAME
       joe - Joe's Own Editor

SYNTAX
       joe [global-options] [ [local-options] filename ]...

       jstar [global-options] [ [local-options] filename ]...
       jmacs [global-options] [ [local-options] filename ]...
       rjoe [global-options] [ [local-options] filename ]...
       jpico [global-options] [ [local-options] filename ]...

DESCRIPTION
       joe is a powerful ASCII-text screen editor. It has a "mode-less" user interface which
       is similar to many user-friendly PC editors. Users of Micro-Pro's WordStar or
       Borland's "Turbo" languages will feel at home. joe is a full featured UNIX
       screen-editor though, and has many features for editing programs and text.
       joe also emulates several other editors.



Un fichier :



SERVICES(5)

NOM
       services - Liste des services réseau internet

DESCRIPTION
       L'emplacement du fichier services est défini par _PATH_SERVICES dans <netdb.h>.
       Il est habituellement défini à /etc/services
       services est un fichier texte ASCII fournissant une correspondance entre des noms
       textuels faciles à mémoriser pour les services et les numéros de ports qui leur sont
       assignés, ainsi que les types de protocoles. Chaque programme en relation avec le réseau
       peut consulter ce fichier pour obtenir le numéro de port (et le protocole) pour ses services.



LABORATOIRE – Accès au système de fichiers – 5

Les exercices des séquences 7 à 11 s’enchaînent.
Il faut donc les accomplir dans l’ordre prévu.

Dans la distribution Debian GNU/Linux :



Dans la suite, le nom de l’utilisateur ordinaire sera représenté par le mot « user ».



Connectez-vous dans un terminal



Affichez le manuel de la commande sort

man sort



Affichez le fichier /etc/services trié

sort /etc/services



Affichez le fichier /etc/services trié en n’affichant qu’une seule fois les lignes multiples

sort -u /etc/services



Affichez le fichier /etc/passwd trié par ordre décroissant

sort -r /etc/passwd



Triez le fichier /etc/passwd, avec enregistrement de la sortie dans ~/f_comptes

sort -o f_comptes /etc/passwd



Affichez le fichier ~/f1.txt trié

sort f1.txt

sort < f1.txt



« | » symbolise un « pipe », c’est-à-dire un « tube de communication » entre deux processus.

Affichez les 5 premières lignes du fichier /etc/services trié

sort /etc/services | head -5

sort < /etc/services | head -5



Affichez les 3 dernières lignes du fichier /etc/services trié

sort /etc/services | tail -3

sort < /etc/services | tail -3



Ecrivez le fichier /etc/services trié dans le fichier ~/Tri-serv en le créant

sort /etc/services > Tri-serv

sort < /etc/services > Tri-serv



Affichez triées, page par page, les lignes du fichier /etc/services qui ne commencent pas par « # »

grep -v '^#' /etc/services | sort | more



Affichez le tri de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

sort f1.txt f2.txt f3.txt

cat f1.txt f2.txt f3.txt | sort



Ecrivez le tri de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt dans le fichier nommé : « f5.txt », en le créant dans ~

sort f1.txt f2.txt f3.txt > f5.txt

cat f1.txt f2.txt f3.txt | sort > f5.txt



Affichez les lignes contenant la chaîne « html » de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

cat f1.txt f2.txt f3.txt | grep html



Affichez les lignes contenant la chaîne « html » du tri de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

sort f1.txt f2.txt f3.txt | grep html

cat f1.txt f2.txt f3.txt | sort | grep html

cat f1.txt f2.txt f3.txt | grep html | sort



Créez dans ~ le fichier nommé : « f6.txt » et écrivez-y les lignes contenant la chaîne « html » de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

cat f1.txt f2.txt f3.txt | grep html > f6.txt



Créez dans ~ le fichier nommé : « f7.txt » et écrivez-y les lignes contenant la chaîne « html » du tri de la concaténation des fichiers ~/f1.txt, ~/f2.txt et ~/f3.txt

sort f1.txt f2.txt f3.txt | grep html > f7.txt

cat f1.txt f2.txt f3.txt | sort | grep html > f7.txt

cat f1.txt f2.txt f3.txt | grep html | sort > f7.txt



Un éditeur est un programme de traitement de texte, pourvu de fonctionnalités plus ou moins complexes.

« gedit » est un éditeur en mode graphique, disponible dans l’interface graphique du système, via le menu « Activités », puis « Afficher les applications ». Il est pourvu de fonctionnalités complexes pour assurer la mise en évidence, selon un mode de coloration, des éléments de nature semblable (commandes, instructions, variables, etc.) d’un texte de script ou de code source écrit dans un langage de programmation.

En mode texte, « nano » est un éditeur installé par défaut. Par contre, « emacs » et « joe » sont des éditeurs à installer.



Affichez le manuel de la commande nano

man nano



Essayez les commandes principales de l’éditeur nano



Installez l’éditeur emacs

apt-get install emacs

aptitude install emacs



Affichez le manuel de la commande emacs

man emacs



Essayez les commandes principales de l’éditeur emacs



Installez l’éditeur joe

apt-get install joe

aptitude install joe



Affichez le manuel de la commande joe

man joe



Essayez les commandes principales de l’éditeur joe



Déconnectez-vous

exit









