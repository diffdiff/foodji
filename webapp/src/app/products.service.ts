import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import Product from './Product';

@Injectable({
  providedIn: 'root'
})
export class ProductsService {

  uri = 'http://localhost:3000/api/v1/products';

  constructor(private http: HttpClient) { }

  addProduct(productName, productDescription, productIngredients, productPrice, productManufacturer) {
    const product = {
      name: productName,
      description: productDescription,
      ingredients: productIngredients,
      price: Number(productPrice),
      manufacturer_id: Number(productManufacturer)
    };
    return this.http.post(`${this.uri}`, product);
  }

  getProducts() {
    return this
      .http
      .get(`${this.uri}`)
      .pipe(
        map((array: any[]) => {
          return array.map(product => {
            return {
              ProductID: product.ID,
              ProductName: product.name,
              ProductDescription: product.description,
              ProductIngredients: product.ingredients,
              ProductPrice: product.price,
              Manufacturer: product.manufacturer_id
            } as Product;
          });
        }));
  }

  editProduct(id) {
    return this
      .http
      .get(`${this.uri}/${id}`)
      .pipe(map((product: any) => {
        return {
          ProductID: product.ID,
          ProductName: product.name,
          ProductDescription: product.description,
          ProductIngredients: product.ingredients,
          ProductPrice: product.price,
          Manufacturer: product.manufacturer_id
        } as Product;
      }));
  }


  updateProduct(ProductName, ProductDescription, ProductIngredients, ProductPrice, id) {
    const product = {
      name: ProductName,
      description: ProductDescription,
      ingredients: ProductIngredients,
      price: Number(ProductPrice)
    };
    return this
      .http
      .put(`${this.uri}/${id}`, product);
  }

  deleteProduct(id) {
    return this
      .http
      .delete(`${this.uri}/${id}`);
  }

}
