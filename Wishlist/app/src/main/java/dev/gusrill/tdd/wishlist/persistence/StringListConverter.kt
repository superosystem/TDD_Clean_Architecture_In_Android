package dev.gusrill.tdd.wishlist.persistence

import androidx.room.TypeConverter

object StringListConverter {
    @TypeConverter
    @JvmStatic
    fun stringListToString(list: MutableList<String>?): String? =
        list?.joinToString(separator = "|")

    @TypeConverter
    @JvmStatic
    fun stringToStringList(string: String?): MutableList<String>? =
        string?.split("|")?.toMutableList()
}