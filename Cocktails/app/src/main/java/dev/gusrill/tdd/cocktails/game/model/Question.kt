package dev.gusrill.tdd.cocktails.game.model

class Question(
    val correntOption: String,
    val incorrectOption: String
) {
    var answeredOption: String? = null
        private set

    fun answer(option: String): Boolean {
        answeredOption = option

        return correntOption == answeredOption
    }
}