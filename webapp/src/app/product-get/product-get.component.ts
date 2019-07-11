import { Component, OnInit } from '@angular/core';
import Product from '../Product';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductsService } from '../products.service';

@Component({
  selector: 'app-product-get',
  templateUrl: './product-get.component.html',
  styleUrls: ['./product-get.component.css']
})
export class ProductGetComponent implements OnInit {
  products: Product[];
  constructor(private route: ActivatedRoute, private router: Router, private ps: ProductsService) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => {
      return false;
    };
  }

  ngOnInit() {
    this.getProducts();
  }

  getProducts() {
    this.ps
      .getProducts()
      .subscribe((data: Product[]) => {
        this.products = data;
      });
  }

  deleteProduct(id) {
    this.ps.deleteProduct(id).subscribe(res => {
      this.products = this.products.filter(p => p.ProductID !== id);
    });
  }
}
