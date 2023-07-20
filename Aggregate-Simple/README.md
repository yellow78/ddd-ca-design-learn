# Aggregate (聚合) 概念與實作

# Aggregate 聚合多個Value object或Entity, 當Aggregate只包含一個Entity, 那這個Entity可視為Aggregate Root

# Agregate是一群相關聯的物件的組合，用於確保更改的一致性可以把它作為一個狀態變更的單位
	例如：當座位Ａ台中-新竹車票訂單取消, 而座位Ａ被改為台中-台北車票訂單, 會有許多狀態變更如：站別、票價、折扣等等

# Aggregate 都需要有一個 Entity 作為他的 Aggregate Root，任何的改變與事件傳遞，都要先通過 Aggregate Root，再傳到裡面的元件

# 固定規定 (Invariant)
	不論是何種狀態變更，都必須先經過 Aggregate Root Order 才能進入跟裡面的物件做互動，而 Aggregate Root 不但可以很好地處理這項需求，也可以同時遵守「資料變化時必須保持一致的規則」

# Aggregate 的設計原則
	* Aggregate Root 的 Entity 必須要在 Bounded Context 中有唯一標示性，它的 ID 不能與其他 Aggregate Root 重複。
	* Aggregate Root 負責檢查邊界內所有固定規則。
	* Aggregate 外的物件不能引用除了 Aggregate Root 之外的任合內部物件。
	* 只有 Aggregate Root 才能直接透過資料庫查詢來獲取。內部的其他物件都要透過 Aggregate Root 才能取得。
	* Aggregate 內部物件可以引用其他 Aggregate Root，但僅能引用該 Aggregate Root 的 ID。
	* 刪除操作必須一次刪除 Aggregate 邊界內的所有物件
	* 提交對 Aggregate 內部的任何物件的修改時，整個 Aggregate 的固定規則都要被滿足，意指一次就要存整個 Aggregate 進去。

# 依上述設計原則, 彙整以下
	1. 在一致性邊界內保護不變條件
		* Aggregate 完成了一個更改操作後，為了要儲存這次成功的操作，你需要將整個 Aggregate 一起存進資料庫
		  才能保證邊界內的不變規則被一起被完成, 一個 Transaction 會對應一個 Aggregate 的更改操作，
		  Aggregate 邊界同時也是 Transaction 的邊界
		
	2. 設計小 Aggregate
		* Aggregate 中除了 Aggregate Root 外，裡面的 Entity 數量應該越少越好，最好只剩下 Value Object。
		  這樣一來，你也可以少考慮很多狀態變化的組合，而且也讓你在操作資料庫時，少做一些 JOIN

	3. 通過 ID 引用其他 Aggregate
		* 一個 Aggregate Root 或是內部的物件想要引用外部的 Aggregate，
		我們會直接讓他們引用外部的 Aggregate Root ID 而非整個物件。這樣可以帶來幾個好處：
		1. 減少記憶體消耗。
		2. 不需要對於另一個 Aggregate 的一致性負責。
		3. 有需要時再透過 Aggregate Root ID 去 Repository 撈出 Aggregate，可提升系統的擴展性

	4. 在邊界外使用最終一致性
		* Strong Consistency (強一致性):
		  或是又被稱作 Immediate Consistency, 所有並行的處理程序以相同的順序處理所有的訪問
		
		* Eventual Consistency (最終一致性):
		  強調不在乎次序，只要最終能夠完成任務即可
		Aggregate中送出 Domain Event後繼續自己的持久化操作，而有訂閱這個 Domain Event 的 Aggregate
		就可以自己處理這個 Event 而不用管原先的 Aggregate 的狀態

# Aggregate 與 Repository 的關係：
	1. 一個 Aggregate 最好對到一個 Repository
	2. 從 Repository 一次可以拿出一整個 Aggregate Object
	3. 寫入 Repository 時也是一次寫入一整個 Aggregate Object

* Aggregate 物件與資料庫裡面的資料是不同的東西，或許他們的資料很接近，又或許 Aggregate 是透過資料庫裡的資料組出來的，
  但將兩者分離可以讓你更專注在 Aggregate 的領域知識上，甚至你可以寫完整個系統的業務邏輯後再去選擇資料庫或是儲存的細節
