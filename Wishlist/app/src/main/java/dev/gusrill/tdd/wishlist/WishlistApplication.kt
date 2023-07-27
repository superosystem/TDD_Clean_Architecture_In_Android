package dev.gusrill.tdd.wishlist

import android.app.Application
import dev.gusrill.tdd.wishlist.di.appModule
import org.koin.android.ext.koin.androidContext
import org.koin.core.context.GlobalContext
import org.koin.core.context.startKoin

class WishlistApplication : Application() {
    override fun onCreate() {
        super.onCreate()
        if (GlobalContext.getOrNull() == null) {
            startKoin {
                // declare Android context
                androidContext(this@WishlistApplication)
                // declare used modules
                modules(appModule)
            }
        }
    }
}
