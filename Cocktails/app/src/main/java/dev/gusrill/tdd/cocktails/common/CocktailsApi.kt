package dev.gusrill.tdd.cocktails.common

import com.google.gson.GsonBuilder
import dev.gusrill.tdd.cocktails.BuildConfig
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Call
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET

interface CocktailsApi {
    @GET("filter.php?a=Alcoholic")
    fun getAlcoholic(): Call<CocktailsContainer>

    companion object Factory {
        fun create(): CocktailsApi {
            val gson = GsonBuilder().create()

            val client = OkHttpClient.Builder().apply {
                if (BuildConfig.DEBUG) {
                    val interceptor = HttpLoggingInterceptor()
                    interceptor.level = HttpLoggingInterceptor.Level.BODY
                    addInterceptor(interceptor)
                }
            }.build()

            val retrofit = Retrofit.Builder()
                .baseUrl("https://www.thecocktaildb.com/api/json/v1/1/")
                .client(client)
                .addConverterFactory(GsonConverterFactory.create(gson))
                .build()

            return retrofit.create(CocktailsApi::class.java)
        }
    }
}