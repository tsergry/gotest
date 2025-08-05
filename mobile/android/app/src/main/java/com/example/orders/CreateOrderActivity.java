package com.example.orders;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ProgressBar;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
import com.example.orders.api.ApiService;
import com.example.orders.api.RetrofitClient;
import com.example.orders.model.Order;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class CreateOrderActivity extends AppCompatActivity {
    
    private EditText etUserId, etProductId, etQuantity;
    private Button btnCreateOrder;
    private ProgressBar progressBar;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_create_order);
        
        initViews();
        setupListeners();
    }
    
    private void initViews() {
        etUserId = findViewById(R.id.et_user_id);
        etProductId = findViewById(R.id.et_product_id);
        etQuantity = findViewById(R.id.et_quantity);
        btnCreateOrder = findViewById(R.id.btn_create_order);
        progressBar = findViewById(R.id.progress_bar);
    }
    
    private void setupListeners() {
        btnCreateOrder.setOnClickListener(v -> createOrder());
    }
    
    private void createOrder() {
        String userIdStr = etUserId.getText().toString().trim();
        String productIdStr = etProductId.getText().toString().trim();
        String quantityStr = etQuantity.getText().toString().trim();
        
        if (userIdStr.isEmpty() || productIdStr.isEmpty() || quantityStr.isEmpty()) {
            Toast.makeText(this, "Заполните все поля", Toast.LENGTH_SHORT).show();
            return;
        }
        
        try {
            int userId = Integer.parseInt(userIdStr);
            int productId = Integer.parseInt(productIdStr);
            int quantity = Integer.parseInt(quantityStr);
            
            if (quantity <= 0) {
                Toast.makeText(this, "Количество должно быть больше 0", Toast.LENGTH_SHORT).show();
                return;
            }
            
            Order order = new Order(userId, productId, quantity);
            sendOrderToServer(order);
            
        } catch (NumberFormatException e) {
            Toast.makeText(this, "Введите корректные числовые значения", Toast.LENGTH_SHORT).show();
        }
    }
    
    private void sendOrderToServer(Order order) {
        showLoading(true);
        
        ApiService apiService = RetrofitClient.getInstance().getApiService();
        Call<Order> call = apiService.createOrder(order);
        
        call.enqueue(new Callback<Order>() {
            @Override
            public void onResponse(Call<Order> call, Response<Order> response) {
                showLoading(false);
                
                if (response.isSuccessful() && response.body() != null) {
                    Toast.makeText(CreateOrderActivity.this, 
                        "Заказ создан успешно! ID: " + response.body().getId(), 
                        Toast.LENGTH_LONG).show();
                    finish();
                } else {
                    Toast.makeText(CreateOrderActivity.this, 
                        "Ошибка создания заказа", Toast.LENGTH_SHORT).show();
                }
            }
            
            @Override
            public void onFailure(Call<Order> call, Throwable t) {
                showLoading(false);
                Toast.makeText(CreateOrderActivity.this, 
                    "Ошибка сети: " + t.getMessage(), Toast.LENGTH_SHORT).show();
            }
        });
    }
    
    private void showLoading(boolean show) {
        progressBar.setVisibility(show ? View.VISIBLE : View.GONE);
        btnCreateOrder.setEnabled(!show);
    }
} 