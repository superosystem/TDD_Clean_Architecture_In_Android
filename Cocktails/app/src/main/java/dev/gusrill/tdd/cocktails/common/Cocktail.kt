package dev.gusrill.tdd.cocktails.common

import android.os.Parcelable
import kotlinx.android.parcel.Parcelize

@Parcelize
data class Cocktail(
    val idDrink: String,
    val strDrink: String,
    val strDrinkThumb: String
): Parcelable
