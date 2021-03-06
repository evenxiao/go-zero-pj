// Code generated by goctl. DO NOT EDIT.
package com.xhb.logic.http.packet.bookstore.model;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import com.xhb.core.packet.HttpData;

public class AddReq extends HttpData {

	private String book = "";
	private long price;

	public AddReq() {
	}

	public AddReq(String book, long price) {
		this.book = book;
		this.price = price;
	}

	@NotNull 
	public String getBook() {
		return this.book;
	}

	public void setBook(@NotNull String book) {
		this.book = book;
	}

	public long getPrice() {
		return this.price;
	}

	public void setPrice(long price) {
		this.price = price;
	}

}
