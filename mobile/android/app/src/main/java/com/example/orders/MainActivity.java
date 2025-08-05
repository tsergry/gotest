package com.example.orders;

import android.content.Intent;
import android.os.Bundle;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;
import com.example.orders.adapter.OrderAdapter;
import com.example.orders.api.ApiService;
import com.example.orders.api.RetrofitClient;
import com.example.orders.model.Order;
import com.google.android.material.floatingactionbutton.FloatingActionButton;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import java.util.ArrayList;
import java.util.List;

public class MainActivity extends AppCompatActivity {
    
    private RecyclerView recyclerView;
    private OrderAdapter orderAdapter;
    private SwipeRefreshLayout swipeRefreshLayout;
    private List<Order> orders = new ArrayList<>();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        
        initViews();
        setupRecyclerView();
        loadOrders();
        
        // Настройка SwipeRefreshLayout
        swipeRefreshLayout.setOnRefreshListener(this::loadOrders);
        
        // Настройка FloatingActionButton
        FloatingActionButton fab = findViewById(R.id.fab_add_order);
        fab.setOnClickListener(v -> {
            Intent intent = new Intent(MainActivity.this, CreateOrderActivity.class);
            startActivity(intent);
        });
    }
    
    private void initViews() {
        recyclerView = findViewById(R.id.recycler_view);
        swipeRefreshLayout = findViewById(R.id.swipe_refresh_layout);
    }
    
    private void setupRecyclerView() {
        orderAdapter = new OrderAdapter(this, orders);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));
        recyclerView.setAdapter(orderAdapter);
    }
    
    private void loadOrders() {
        ApiService apiService = RetrofitClient.getInstance().getApiService();
        
        Call<List<Order>> call = apiService.getOrders();
        call.enqueue(new Callback<List<Order>>() {
            @Override
            public void onResponse(Call<List<Order>> call, Response<List<Order>> response) {
                swipeRefreshLayout.setRefreshing(false);
                
                if (response.isSuccessful() && response.body() != null) {
                    orders.clear();
                    orders.addAll(response.body());
                    orderAdapter.notifyDataSetChanged();
                } else {
                    Toast.makeText(MainActivity.this, "Ошибка загрузки заказов", Toast.LENGTH_SHORT).show();
                }
            }
            
            @Override
            public void onFailure(Call<List<Order>> call, Throwable t) {
                swipeRefreshLayout.setRefreshing(false);
                Toast.makeText(MainActivity.this, "Ошибка сети: " + t.getMessage(), Toast.LENGTH_SHORT).show();
            }
        });
    }
    
    @Override
    protected void onResume() {
        super.onResume();
        loadOrders(); // Обновляем список при возвращении в активность
    }
} 