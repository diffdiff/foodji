import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProductAddComponent } from './product-add/product-add.component';
import { ProductEditComponent } from './product-edit/product-edit.component';
import { ProductGetComponent } from './product-get/product-get.component';
import { ManufacturerAddComponent } from './manufacturer-add/manufacturer-add.component';
import { ManufacturerEditComponent } from './manufacturer-edit/manufacturer-edit.component';
import { ManufacturerGetComponent } from './manufacturer-get/manufacturer-get.component';

const routes: Routes = [
  {
    path: 'product/create',
    component: ProductAddComponent
  },
  {
    path: 'product/edit/:id',
    component: ProductEditComponent
  },
  {
    path: 'product/index',
    component: ProductGetComponent
  },
  {
    path: 'manufacturer/create',
    component: ManufacturerAddComponent
  },
  {
    path: 'manufacturer/edit/:id',
    component: ManufacturerEditComponent
  },
  {
    path: 'manufacturer/index',
    component: ManufacturerGetComponent
  }
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
