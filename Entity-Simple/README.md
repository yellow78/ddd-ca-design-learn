# Entity 概念與實作

# 當一個物件的生命期開始被你追蹤時，你就需要為它建立身份標誌 (identity)，而這個物件我們就稱為 Entity

# Entity 是一個物件，或者說是一個 Class，跟 Clean Architecture 中的 Entity Layer 是不同的概念

# Entity 的特徵
	1. Entity 最重要的是他的 ID。
	2. 兩個 Entity 不論其他狀態，只要 ID 相同就是相同的物件。
	3. 除了 ID，他們其他的狀態是可變的 (mutable)。
	4. 他們可能擁有很長的壽命，甚至不會被刪除。
	5. 一個 Entity 是可變的、長壽的，所以通常會有複雜的生命週期變化，如一套 CRUD 的操作。

# 範例：高鐵車票對號座, 在賣票系統中需要追蹤每站位置的使用狀況, 在此會為Entity

# 範例：高鐵車票自由座, 在賣票系統中有票即可上車, 不追蹤位置的使用狀況, 在此就不會為Entity

# entity 實作：Entity-Simple/internal/train/entity/seat.go 

# 二次封裝UUID : Entity-Simple/internal/core/uuid.go

# entity 不依賴外部套件, 因為這是領域核心邏輯