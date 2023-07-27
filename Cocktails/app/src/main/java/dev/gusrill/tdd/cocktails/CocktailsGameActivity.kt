package dev.gusrill.tdd.cocktails

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import dev.gusrill.tdd.cocktails.databinding.ActivityGameBinding

class CocktailsGameActivity : AppCompatActivity() {

    private lateinit var binding: ActivityGameBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityGameBinding.inflate(layoutInflater)
        setContentView(binding.root)
    }
}