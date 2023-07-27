package dev.gusrill.tdd.wishlist.persistence

import androidx.lifecycle.LiveData
import dev.gusrill.tdd.wishlist.domain.Wishlist

interface Repository {
    fun saveWishlist(wishlist: Wishlist)
    fun getWishlists(): LiveData<List<Wishlist>>
    fun getWishlist(id: Int): LiveData<Wishlist>
    fun saveWishlistItem(wishlist: Wishlist, name: String)
}

class RepositoryImpl(
    private val wishlistDao: WishlistDao
) : Repository {
    override fun saveWishlist(wishlist: Wishlist) {
        wishlistDao.save(wishlist)
    }

    override fun getWishlists(): LiveData<List<Wishlist>> {
        return wishlistDao.getAll()
    }

    override fun getWishlist(id: Int): LiveData<Wishlist> {
        return wishlistDao.findById(id)
    }

    override fun saveWishlistItem(wishlist: Wishlist, name: String) {
        wishlistDao.save(wishlist.copy(wishes = wishlist.wishes + name))
    }
}