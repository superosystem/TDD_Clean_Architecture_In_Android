package dev.gusrill.tdd.cocktails.game.model

class Game{
    var currentScore = 0
        private set

    var highestScore = 0
        private set

    fun incrementScore() {
        currentScore++
        highestScore++
    }
}