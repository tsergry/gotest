package com.example.orders;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.ProgressBar;
import android.widget.TextView;
import android.widget.Toast;
import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;
import com.example.orders.api.ApiService;
import com.example.orders.api.RetrofitClient;
import com.example.orders.model.Order;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;
import java.text.SimpleDateFormat;
import java.util.Locale;

public class OrderDetailActivity extends AppCompatActivity {
    
    private TextView tvOrderId, tvUserName, tvUserEmail, tvProductName, tvQuantity, 
                    tvUnitPrice, tvTotalPrice, tvStatus, tvOrderDate, tvCreatedAt;
    private Button btnUpdateStatus, btnDeleteOrder;
    private ProgressBar progressBar;
    private int orderId;
    private Order currentOrder;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_order_detail);
        
        orderId = getIntent().getIntExtra("order_id", -1);
        if (orderId == -1) {
            Toast.makeText(this, "Ошибка: ID заказа не найден", Toast.LENGTH_SHORT).show();
            finish();
            return;
        }
        
        initViews();
        loadOrderDetails();
    }
    
    private void initViews() {
        tvOrderId = findViewById(R.id.tv_order_id);
        tvUserName = findViewById(R.id.tv_user_name);
        tvUserEmail = findViewById(R.id.tv_user_email);
        tvProductName = findViewById(R.id.tv_product_name);
        tvQuantity = findViewById(R.id.tv_quantity);
        tvUnitPrice = findViewById(R.id.tv_unit_price);
        tvTotalPrice = findViewById(R.id.tv_total_price);
        tvStatus = findViewById(R.id.tv_status);
        tvOrderDate = findViewById(R.id.tv_order_date);
        tvCreatedAt = findViewById(R.id.tv_created_at);
        btnUpdateStatus = findViewById(R.id.btn_update_status);
        btnDeleteOrder = findViewById(R.id.btn_delete_order);
        progressBar = findViewById(R.id.progress_bar);
        
        btnUpdateStatus.setOnClickListener(v -> showStatusUpdateDialog());
        btnDeleteOrder.setOnClickListener(v -> showDeleteConfirmDialog());
    }
    
    private void loadOrderDetails() {
        showLoading(true);
        
        ApiService apiService = RetrofitClient.getInstance().getApiService();
        Call<Order> call = apiService.getOrder(orderId);
        
        call.enqueue(new Callback<Order>() {
            @Override
            public void onResponse(Call<Order> call, Response<Order> response) {
                showLoading(false);
                
                if (response.isSuccessful() && response.body() != null) {
                    currentOrder = response.body();
                    displayOrderDetails(currentOrder);
                } else {
                    Toast.makeText(OrderDetailActivity.this, 
                        "Ошибка загрузки заказа", Toast.LENGTH_SHORT).show();
                    finish();
                }
            }
            
            @Override
            public void onFailure(Call<Order> call, Throwable t) {
                showLoading(false);
                Toast.makeText(OrderDetailActivity.this, 
                    "Ошибка сети: " + t.getMessage(), Toast.LENGTH_SHORT).show();
                finish();
            }
        });
    }
    
    private void displayOrderDetails(Order order) {
        SimpleDateFormat dateFormat = new SimpleDateFormat("dd.MM.yyyy HH:mm", Locale.getDefault());
        
        tvOrderId.setText("Заказ #" + order.getId());
        tvUserName.setText(order.getUserName());
        tvUserEmail.setText(order.getUserEmail());
        tvProductName.setText(order.getProductName());
        tvQuantity.setText(String.valueOf(order.getQuantity()));
        tvUnitPrice.setText(String.format(Locale.getDefault(), "%.2f ₽", order.getUnitPrice()));
        tvTotalPrice.setText(String.format(Locale.getDefault(), "%.2f ₽", order.getTotalPrice()));
        tvStatus.setText(order.getStatus());
        tvOrderDate.setText(dateFormat.format(order.getOrderDate()));
        tvCreatedAt.setText(dateFormat.format(order.getCreatedAt()));
        
        // Устанавливаем цвет статуса
        switch (order.getStatus().toLowerCase()) {
            case "pending":
                tvStatus.setTextColor(getResources().getColor(android.R.color.holo_orange_dark));
                break;
            case "confirmed":
                tvStatus.setTextColor(getResources().getColor(android.R.color.holo_blue_dark));
                break;
            case "shipped":
                tvStatus.setTextColor(getResources().getColor(android.R.color.holo_purple));
                break;
            case "delivered":
                tvStatus.setTextColor(getResources().getColor(android.R.color.holo_green_dark));
                break;
            case "cancelled":
                tvStatus.setTextColor(getResources().getColor(android.R.color.holo_red_dark));
                break;
        }
    }
    
    private void showStatusUpdateDialog() {
        String[] statuses = {"pending", "confirmed", "shipped", "delivered", "cancelled"};
        
        AlertDialog.Builder builder = new AlertDialog.Builder(this);
        builder.setTitle("Обновить статус заказа")
               .setItems(statuses, (dialog, which) -> {
                   String newStatus = statuses[which];
                   updateOrderStatus(newStatus);
               })
               .setNegativeButton("Отмена", null)
               .show();
    }
    
    private void updateOrderStatus(String newStatus) {
        showLoading(true);
        
        ApiService apiService = RetrofitClient.getInstance().getApiService();
        ApiService.StatusUpdate statusUpdate = new ApiService.StatusUpdate(newStatus);
        Call<ApiService.ApiResponse> call = apiService.updateOrderStatus(orderId, statusUpdate);
        
        call.enqueue(new Callback<ApiService.ApiResponse>() {
            @Override
            public void onResponse(Call<ApiService.ApiResponse> call, Response<ApiService.ApiResponse> response) {
                showLoading(false);
                
                if (response.isSuccessful()) {
                    Toast.makeText(OrderDetailActivity.this, 
                        "Статус обновлен: " + newStatus, Toast.LENGTH_SHORT).show();
                    loadOrderDetails(); // Перезагружаем данные
                } else {
                    Toast.makeText(OrderDetailActivity.this, 
                        "Ошибка обновления статуса", Toast.LENGTH_SHORT).show();
                }
            }
            
            @Override
            public void onFailure(Call<ApiService.ApiResponse> call, Throwable t) {
                showLoading(false);
                Toast.makeText(OrderDetailActivity.this, 
                    "Ошибка сети: " + t.getMessage(), Toast.LENGTH_SHORT).show();
            }
        });
    }
    
    private void showDeleteConfirmDialog() {
        new AlertDialog.Builder(this)
                .setTitle("Удаление заказа")
                .setMessage("Вы уверены, что хотите удалить этот заказ?")
                .setPositiveButton("Удалить", (dialog, which) -> deleteOrder())
                .setNegativeButton("Отмена", null)
                .show();
    }
    
    private void deleteOrder() {
        showLoading(true);
        
        // Здесь нужно добавить метод удаления в API
        Toast.makeText(this, "Функция удаления будет добавлена позже", Toast.LENGTH_SHORT).show();
        showLoading(false);
    }
    
    private void showLoading(boolean show) {
        progressBar.setVisibility(show ? View.VISIBLE : View.GONE);
        btnUpdateStatus.setEnabled(!show);
        btnDeleteOrder.setEnabled(!show);
    }
} 