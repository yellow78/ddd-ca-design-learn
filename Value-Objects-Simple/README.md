# Value Object 概念與實作

# Value Object與Entity相近, 當物件在系統上不在乎生命週期變化

# Value Objects 特徵
	1. 描述性 : 度量或描述了領域中的某項概念
	2. 不變性 :  Value Object 在創建後就不能再改變了, 可以被替換掉
	3. 概念整體性 : 將相關屬性組成一個「概念整體 (Conceptual Whole)」
	4. 替換性 : 當度量與描述改變時，可以用另一個 Value Object 替換
	5. 相等性 : Value Object 沒有身份標識的概念，Value Object 身上的屬性的值都相等，那我們就會說這兩個 Value Object 是相等的
	6. 無副作用 : Value Object的狀態不變性, 在程式上好管理

# 範例：高鐵從台北到台中的票, 在賣票系統中座位價格會依距離或是訂票日期因素而改變, 在此會為Entity, 但起點與終點則是固定的, 在此會為Value Object
