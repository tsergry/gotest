package com.example.orders.adapter;

import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import androidx.annotation.NonNull;
import androidx.cardview.widget.CardView;
import androidx.recyclerview.widget.RecyclerView;
import com.example.orders.OrderDetailActivity;
import com.example.orders.R;
import com.example.orders.model.Order;
import java.util.List;
import java.util.Locale;

public class OrderAdapter extends RecyclerView.Adapter<OrderAdapter.OrderViewHolder> {
    
    private List<Order> orders;
    private Context context;

    public OrderAdapter(Context context, List<Order> orders) {
        this.context = context;
        this.orders = orders;
    }

    @NonNull
    @Override
    public OrderViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(context).inflate(R.layout.item_order, parent, false);
        return new OrderViewHolder(view);
    }

    @Override
    public void onBindViewHolder(@NonNull OrderViewHolder holder, int position) {
        Order order = orders.get(position);
        
        holder.tvOrderId.setText("Заказ #" + order.getId());
        holder.tvUserName.setText(order.getUserName());
        holder.tvProductName.setText(order.getProductName());
        holder.tvQuantity.setText("Количество: " + order.getQuantity());
        holder.tvTotalPrice.setText(String.format(Locale.getDefault(), "%.2f ₽", order.getTotalPrice()));
        holder.tvStatus.setText(order.getStatus());
        
        // Устанавливаем цвет статуса
        switch (order.getStatus().toLowerCase()) {
            case "pending":
                holder.tvStatus.setTextColor(context.getResources().getColor(android.R.color.holo_orange_dark));
                break;
            case "confirmed":
                holder.tvStatus.setTextColor(context.getResources().getColor(android.R.color.holo_blue_dark));
                break;
            case "shipped":
                holder.tvStatus.setTextColor(context.getResources().getColor(android.R.color.holo_purple));
                break;
            case "delivered":
                holder.tvStatus.setTextColor(context.getResources().getColor(android.R.color.holo_green_dark));
                break;
            case "cancelled":
                holder.tvStatus.setTextColor(context.getResources().getColor(android.R.color.holo_red_dark));
                break;
        }
        
        holder.cardView.setOnClickListener(v -> {
            Intent intent = new Intent(context, OrderDetailActivity.class);
            intent.putExtra("order_id", order.getId());
            context.startActivity(intent);
        });
    }

    @Override
    public int getItemCount() {
        return orders.size();
    }

    public void updateOrders(List<Order> newOrders) {
        this.orders = newOrders;
        notifyDataSetChanged();
    }

    static class OrderViewHolder extends RecyclerView.ViewHolder {
        CardView cardView;
        TextView tvOrderId, tvUserName, tvProductName, tvQuantity, tvTotalPrice, tvStatus;

        public OrderViewHolder(@NonNull View itemView) {
            super(itemView);
            cardView = itemView.findViewById(R.id.card_order);
            tvOrderId = itemView.findViewById(R.id.tv_order_id);
            tvUserName = itemView.findViewById(R.id.tv_user_name);
            tvProductName = itemView.findViewById(R.id.tv_product_name);
            tvQuantity = itemView.findViewById(R.id.tv_quantity);
            tvTotalPrice = itemView.findViewById(R.id.tv_total_price);
            tvStatus = itemView.findViewById(R.id.tv_status);
        }
    }
} 