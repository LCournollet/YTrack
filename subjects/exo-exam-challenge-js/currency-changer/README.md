# Currency Changer

## Instruction

Durant cet exercice vous allez devoir créer un convertisseur de monnaie. Le but est de pouvoir convertir la valeur d'entrée d'un input, faire le calcul de la conversion et de ressortir la valeur dans un input de sortie

Cette conversion devra être affichée : 

- Lorsque l'utilisateur arrête d'écrire et quitte l'input (cf: STEP 2)
- Que l'utilisateur change de monnaie d'entrée ou de sortie

Le client pourra choisir quelle devise utiliser parmi les 4 devises présente dans la page HTML fourni [ici](https://github.com/Lyon-Ynov-Campus/YTrack/blob/master/subjects/exo-exam-challenge-js/currency-changer/index.html).

## STEP 1

Pour effectuer votre conversion, vous devrez utiliser l'objet suivant : 

``` js
const money = {
dollar: 0.84,
euro: 1,
livre: 1.14,
bitcoin: 40000
}
```

La formule pour calculer la valeur de sortie est la suivante : 

`(devise_entrée * quantité) / devise_sortie`

où : 

- **devise_entrée** est la monnaie à convertir
- **quantité** est la valeur de l'input avec `id="entry"`
- **devise_sortie** est la monnaie dans laquelle on veut convertir

## STEP 2

- Changer le style des devises sélectionnées 
- Ajouter un eventListener sur l'input avec `id="entry"` ([change event](https://developer.mozilla.org/fr/docs/Web/API/HTMLElement/change_event))

## ⚠️Attention⚠️

- Les attributs onClick des éléments HTML sont à compléter
- Vous devrez utiliser les propriétés **data-...** présentent sur les `buttons` pour déterminer la devise sur laquelle vous avez cliqué ([dataset js](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset))
- Le fichier js principal sera nommé ```cash.js```

## Files:
[index.html](https://github.com/Lyon-Ynov-Campus/YTrack/blob/master/subjects/exo-exam-challenge-js/currency-changer/index.html)

## Lien utile

[access object](https://stackoverflow.com/questions/4255472/javascript-object-access-variable-property-by-name-as-string)

## Exemple

![exemple](https://i.imgur.com/7OFe6Dd.gif)