package com.example.orders.api;

import com.example.orders.model.Order;
import com.example.orders.model.Product;
import java.util.List;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;
import retrofit2.http.PUT;
import retrofit2.http.Query;

public interface ApiService {
    
    // Orders endpoints
    @GET("api/orders")
    Call<List<Order>> getOrders();
    
    @GET("api/orders/get")
    Call<Order> getOrder(@Query("id") int id);
    
    @GET("api/orders/user")
    Call<List<Order>> getOrdersByUser(@Query("user_id") int userId);
    
    @GET("api/orders/status")
    Call<List<Order>> getOrdersByStatus(@Query("status") String status);
    
    @POST("api/orders")
    Call<Order> createOrder(@Body Order order);
    
    @PUT("api/orders/update-status")
    Call<ApiResponse> updateOrderStatus(@Query("id") int id, @Body StatusUpdate statusUpdate);
    
    // Products endpoints
    @GET("api/products")
    Call<List<Product>> getProducts();
    
    @GET("api/products/get")
    Call<Product> getProduct(@Query("id") int id);
    
    @GET("api/products/category")
    Call<List<Product>> getProductsByCategory(@Query("category") String category);
    
    @POST("api/products")
    Call<Product> createProduct(@Body Product product);
    
    public static class StatusUpdate {
        private String status;
        
        public StatusUpdate(String status) {
            this.status = status;
        }
        
        public String getStatus() {
            return status;
        }
        
        public void setStatus(String status) {
            this.status = status;
        }
    }
    
    public static class ApiResponse {
        private String message;
        
        public String getMessage() {
            return message;
        }
        
        public void setMessage(String message) {
            this.message = message;
        }
    }
} 