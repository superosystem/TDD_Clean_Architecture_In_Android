package dev.gusrill.tdd.wishlist.ui

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import dev.gusrill.tdd.wishlist.domain.Wishlist
import dev.gusrill.tdd.wishlist.persistence.Repository

class MainViewModel(
    private val repository: Repository
) : ViewModel() {
    fun saveNewList(name: String) {
        repository.saveWishlist(Wishlist(name, listOf()))
    }

    fun getWishlists(): LiveData<List<Wishlist>> {
        return repository.getWishlists()
    }
}