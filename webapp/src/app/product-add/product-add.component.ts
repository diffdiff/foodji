import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductsService } from '../products.service';
import { ManufacturerService } from '../manufacturer.service';
import Manufacturer from '../Manufacturer';

@Component({
  selector: 'app-product-add',
  templateUrl: './product-add.component.html',
  styleUrls: ['./product-add.component.css']
})
export class ProductAddComponent implements OnInit {

  angForm: FormGroup;
  manufacturers: Manufacturer[];

  constructor(private route: ActivatedRoute, private router: Router, private fb: FormBuilder, private ps: ProductsService, private ms: ManufacturerService) {
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      ProductName: ['', Validators.required],
      ProductDescription: ['', Validators.required],
      ProductIngredients: ['', Validators.required],
      ProductPrice: ['', Validators.required],
      Manufacturer: ['', Validators.required]
    });
  }

  addProduct(ProductName, ProductDescription, ProductIngredients, ProductPrice, Manufacturer) {
    this.route.params.subscribe(params => {
      this.ps.addProduct(ProductName, ProductDescription, ProductIngredients, ProductPrice, Manufacturer).subscribe(x => {
        this.router.navigate(['product/index']);
      });
    });
  }

  ngOnInit() {
    this.ms.getManufacturer().subscribe(data => {
      this.manufacturers = data;
    });
  }

}
