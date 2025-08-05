package com.example.orders.model;

import com.google.gson.annotations.SerializedName;
import java.util.Date;

public class Order {
    @SerializedName("id")
    private int id;
    
    @SerializedName("user_id")
    private int userId;
    
    @SerializedName("product_id")
    private int productId;
    
    @SerializedName("quantity")
    private int quantity;
    
    @SerializedName("total_price")
    private double totalPrice;
    
    @SerializedName("status")
    private String status;
    
    @SerializedName("order_date")
    private Date orderDate;
    
    @SerializedName("created_at")
    private Date createdAt;
    
    @SerializedName("updated_at")
    private Date updatedAt;
    
    @SerializedName("user_name")
    private String userName;
    
    @SerializedName("user_email")
    private String userEmail;
    
    @SerializedName("product_name")
    private String productName;
    
    @SerializedName("unit_price")
    private double unitPrice;

    // Constructors
    public Order() {}

    public Order(int userId, int productId, int quantity) {
        this.userId = userId;
        this.productId = productId;
        this.quantity = quantity;
        this.status = "pending";
    }

    // Getters and Setters
    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getUserId() {
        return userId;
    }

    public void setUserId(int userId) {
        this.userId = userId;
    }

    public int getProductId() {
        return productId;
    }

    public void setProductId(int productId) {
        this.productId = productId;
    }

    public int getQuantity() {
        return quantity;
    }

    public void setQuantity(int quantity) {
        this.quantity = quantity;
    }

    public double getTotalPrice() {
        return totalPrice;
    }

    public void setTotalPrice(double totalPrice) {
        this.totalPrice = totalPrice;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Date getOrderDate() {
        return orderDate;
    }

    public void setOrderDate(Date orderDate) {
        this.orderDate = orderDate;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public Date getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(Date updatedAt) {
        this.updatedAt = updatedAt;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getUserEmail() {
        return userEmail;
    }

    public void setUserEmail(String userEmail) {
        this.userEmail = userEmail;
    }

    public String getProductName() {
        return productName;
    }

    public void setProductName(String productName) {
        this.productName = productName;
    }

    public double getUnitPrice() {
        return unitPrice;
    }

    public void setUnitPrice(double unitPrice) {
        this.unitPrice = unitPrice;
    }

    @Override
    public String toString() {
        return "Order{" +
                "id=" + id +
                ", userId=" + userId +
                ", productId=" + productId +
                ", quantity=" + quantity +
                ", totalPrice=" + totalPrice +
                ", status='" + status + '\'' +
                ", userName='" + userName + '\'' +
                ", productName='" + productName + '\'' +
                '}';
    }
} 