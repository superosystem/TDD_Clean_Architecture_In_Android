package dev.gusrill.tdd.wishlist

import androidx.arch.core.executor.testing.InstantTaskExecutorRule
import androidx.lifecycle.Observer
import dev.gusrill.tdd.wishlist.domain.Wishlist
import dev.gusrill.tdd.wishlist.persistence.RepositoryImpl
import dev.gusrill.tdd.wishlist.persistence.WishlistDao
import dev.gusrill.tdd.wishlist.persistence.WishlistDaoImpl
import dev.gusrill.tdd.wishlist.ui.wishlist.DetailViewModel
import org.junit.Rule
import org.junit.Test
import org.mockito.Mockito
import org.mockito.kotlin.any
import org.mockito.kotlin.mock
import org.mockito.kotlin.verify

class DetailViewModelTest {
    @get:Rule
    var instantTaskExecutorRule = InstantTaskExecutorRule()

    private val wishlistDao: WishlistDao = Mockito.spy(WishlistDaoImpl())
    private val viewModel = DetailViewModel(RepositoryImpl(wishlistDao))

    @Test
    fun saveNewItemCallsDatabase() {
        viewModel.saveNewItem(
            Wishlist(
                "Victoria",
                listOf("RW Android Apprentice Book", "Android phone"), 1
            ),
            "Smart watch"
        )

        verify(wishlistDao).save(any())
    }

    @Test
    fun saveNewItemSavesData() {
        val wishlist = Wishlist(
            "Victoria",
            listOf("RW Android Apprentice Book", "Android phone"), 1
        )
        val name = "Smart watch"
        viewModel.saveNewItem(wishlist, name)

        val mockObserver = mock<Observer<Wishlist>>()
        wishlistDao.findById(wishlist.id)
            .observeForever(mockObserver)
        verify(mockObserver).onChanged(
            wishlist.copy(wishes = wishlist.wishes + name)
        )
    }

    @Test
    fun getWishListCallsDatabase() {
        viewModel.getWishlist(1)

        verify(wishlistDao).findById(any())
    }

    @Test
    fun getWishListReturnsCorrectData() {
        val wishlist = Wishlist(
            "Victoria",
            listOf("RW Android Apprentice Book", "Android phone"), 1
        )
        wishlistDao.save(wishlist)
        val mockObserver = mock<Observer<Wishlist>>()
        viewModel.getWishlist(1).observeForever(mockObserver)
        verify(mockObserver).onChanged(wishlist)
    }
}
