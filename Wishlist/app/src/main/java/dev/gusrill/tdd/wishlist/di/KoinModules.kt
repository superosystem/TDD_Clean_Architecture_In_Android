package dev.gusrill.tdd.wishlist.di

import dev.gusrill.tdd.wishlist.persistence.Repository
import dev.gusrill.tdd.wishlist.persistence.RepositoryImpl
import dev.gusrill.tdd.wishlist.persistence.WishlistDao
import dev.gusrill.tdd.wishlist.persistence.WishlistDaoImpl
import dev.gusrill.tdd.wishlist.ui.MainViewModel
import dev.gusrill.tdd.wishlist.ui.wishlist.DetailViewModel
import org.koin.androidx.viewmodel.dsl.viewModel
import org.koin.dsl.module

val appModule = module {
    single<Repository> { RepositoryImpl(get()) }
    single<WishlistDao> { WishlistDaoImpl() }
    viewModel { MainViewModel(get()) }
    viewModel { DetailViewModel(get()) }
}