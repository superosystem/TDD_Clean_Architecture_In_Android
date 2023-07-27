package dev.gusrill.tdd.wishlist.persistence

import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.TypeConverters
import dev.gusrill.tdd.wishlist.domain.Wishlist

@Database(entities = [Wishlist::class], version = 1)
@TypeConverters(StringListConverter::class)
abstract class WishlistRoom : RoomDatabase()