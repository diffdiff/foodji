import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductsService } from '../products.service';
import { ManufacturerService } from '../manufacturer.service';
import Manufacturer from '../Manufacturer';

@Component({
  selector: 'app-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.css']
})
export class ProductEditComponent implements OnInit {

  angForm: FormGroup;
  public product: any = {};
  manufacturers: Manufacturer[];

  constructor(private route: ActivatedRoute, private router: Router, private ps: ProductsService, private fb: FormBuilder,  private ms: ManufacturerService) {
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      ProductName: ['', Validators.required],
      ProductDescription: ['', Validators.required],
      ProductPrice: ['', Validators.required],
      ProductIngredients: ['', Validators.required],
      ProductManufacturer: ['', '']
    });
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.ps.editProduct(params.id).subscribe(product => {
        this.product = product;
      });

      this.ms.getManufacturer().subscribe(data => {
        this.manufacturers = data;
      });
    });
  }

  updateProduct(ProductName, ProductDescription, ProductIngredients, ProductPrice) {
    this.route.params.subscribe(params => {
      this.ps.updateProduct(ProductName, ProductDescription, ProductIngredients, ProductPrice, params.id).subscribe(x => {
        this.router.navigate(['product/index']);
      });
    });
  }
}
