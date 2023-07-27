package dev.gusrill.tdd.wishlist.ui.wishlist

import androidx.lifecycle.LiveData
import androidx.lifecycle.ViewModel
import dev.gusrill.tdd.wishlist.domain.Wishlist
import dev.gusrill.tdd.wishlist.persistence.Repository

class DetailViewModel(
    private val repository: Repository
) : ViewModel() {
    fun saveNewItem(wishlist: Wishlist, name: String) {
        repository.saveWishlistItem(wishlist, name)
    }

    fun getWishlist(id: Int): LiveData<Wishlist> {
        return repository.getWishlist(id)
    }
}