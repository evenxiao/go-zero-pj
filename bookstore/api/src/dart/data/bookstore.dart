// --bookstore--

class addReq{
	
	/// 
	final {string} book;
	
	/// 
	final {int64} price;
	
	addReq({ 
		this.book,
		this.price,
	});
	factory addReq.fromJson(Map<String,dynamic> m) {
		return addReq(
			book: 