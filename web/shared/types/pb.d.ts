interface Product {
  id: string;
  name: string;
  photo: string;
  price: number;
  barcode: string;
}

interface Order {
  created: string
  id: string
  user: string
  rows: string[]
  total: number
  expand: {
    rows: OrderRow[]
  }
}

interface OrderRow {
  product: string
  amount: number
  price: number
  expand: {
    product: Product
  }
}
