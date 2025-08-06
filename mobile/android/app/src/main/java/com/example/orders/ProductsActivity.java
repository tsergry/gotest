package com.example.orders;

import android.content.Intent;
import android.os.Bundle;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;
import com.example.orders.adapter.ProductAdapter;
import com.example.orders.api.ApiService;
import com.example.orders.api.RetrofitClient;
import com.example.orders.model.Product;
import com.google.android.material.floatingactionbutton.FloatingActionButton;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import java.util.ArrayList;
import java.util.List;

public class ProductsActivity extends AppCompatActivity implements ProductAdapter.OnProductClickListener {
    
    private RecyclerView recyclerView;
    private ProductAdapter productAdapter;
    private SwipeRefreshLayout swipeRefreshLayout;
    private List<Product> products = new ArrayList<>();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_products);
        
        initViews();
        setupRecyclerView();
        loadProducts();
        
        // Настройка SwipeRefreshLayout
        swipeRefreshLayout.setOnRefreshListener(this::loadProducts);
        
        // Настройка FloatingActionButton
        FloatingActionButton fab = findViewById(R.id.fab_add_product);
        fab.setOnClickListener(v -> {
            // TODO: Добавить создание продукта
            Toast.makeText(this, "Создание продукта будет добавлено позже", Toast.LENGTH_SHORT).show();
        });
    }
    
    private void initViews() {
        recyclerView = findViewById(R.id.recycler_view);
        swipeRefreshLayout = findViewById(R.id.swipe_refresh_layout);
    }
    
    private void setupRecyclerView() {
        productAdapter = new ProductAdapter(this, products);
        productAdapter.setOnProductClickListener(this);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));
        recyclerView.setAdapter(productAdapter);
    }
    
    private void loadProducts() {
        ApiService apiService = RetrofitClient.getInstance().getApiService();
        
        Call<List<Product>> call = apiService.getProducts();
        call.enqueue(new Callback<List<Product>>() {
            @Override
            public void onResponse(Call<List<Product>> call, Response<List<Product>> response) {
                swipeRefreshLayout.setRefreshing(false);
                
                if (response.isSuccessful() && response.body() != null) {
                    products.clear();
                    products.addAll(response.body());
                    productAdapter.notifyDataSetChanged();
                } else {
                    Toast.makeText(ProductsActivity.this, "Ошибка загрузки продуктов", Toast.LENGTH_SHORT).show();
                }
            }
            
            @Override
            public void onFailure(Call<List<Product>> call, Throwable t) {
                swipeRefreshLayout.setRefreshing(false);
                Toast.makeText(ProductsActivity.this, "Ошибка сети: " + t.getMessage(), Toast.LENGTH_SHORT).show();
            }
        });
    }
    
    @Override
    public void onProductClick(Product product) {
        // Показываем детали продукта
        Toast.makeText(this, "Выбран продукт: " + product.getName(), Toast.LENGTH_SHORT).show();
        // TODO: Добавить детальную активность для продукта
    }
    
    @Override
    protected void onResume() {
        super.onResume();
        loadProducts(); // Обновляем список при возвращении в активность
    }
} 