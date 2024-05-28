CREATE INDEX IF NOT EXISTS merchant_id ON merchants (merchant_id);
CREATE INDEX IF NOT EXISTS merchant_name ON merchants (name);
CREATE INDEX IF NOT EXISTS merchant_category_SmallRestaurant ON merchants (merchant_category) WHERE merchant_category = 'SmallRestaurant';
CREATE INDEX IF NOT EXISTS merchant_category_MediumRestaurant ON merchants (merchant_category) WHERE merchant_category = 'MediumRestaurant';
CREATE INDEX IF NOT EXISTS merchant_category_LargeRestaurant ON merchants (merchant_category) WHERE merchant_category = 'LargeRestaurant';
CREATE INDEX IF NOT EXISTS merchant_category_MerchandiseRestaurant ON merchants (merchant_category) WHERE merchant_category = 'MerchandiseRestaurant';
CREATE INDEX IF NOT EXISTS merchant_category_BoothKiosk ON merchants (merchant_category) WHERE merchant_category = 'BoothKiosk';
CREATE INDEX IF NOT EXISTS merchant_category_ConvenienceStore ON merchants (merchant_category) WHERE merchant_category = 'ConvenienceStore';
CREATE INDEX IF NOT EXISTS merchant_created_at ON merchants (created_at);

CREATE INDEX IF NOT EXISTS merchant_item_id ON merchant_items (item_id);
CREATE INDEX IF NOT EXISTS merchant_item_merchant_id ON merchant_items (merchant_id);
CREATE INDEX IF NOT EXISTS merchant_item_name ON merchant_items (name);
CREATE INDEX IF NOT EXISTS merchant_item_product_category_Beverage ON merchant_items (product_category) WHERE product_category = 'Beverage';
CREATE INDEX IF NOT EXISTS merchant_item_product_category_Food ON merchant_items (product_category) WHERE product_category = 'Food';
CREATE INDEX IF NOT EXISTS merchant_item_product_category_Snack ON merchant_items (product_category) WHERE product_category = 'Snack';
CREATE INDEX IF NOT EXISTS merchant_item_product_category_Condiments ON merchant_items (product_category) WHERE product_category = 'Condiments';
CREATE INDEX IF NOT EXISTS merchant_item_product_category_Additions ON merchant_items (product_category) WHERE product_category = 'Additions';
CREATE INDEX IF NOT EXISTS merchant_item_created_at ON merchant_items (created_at);