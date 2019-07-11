import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ManufacturerService } from '../manufacturer.service';

@Component({
  selector: 'app-manufacturer-edit',
  templateUrl: './manufacturer-edit.component.html',
  styleUrls: ['./manufacturer-edit.component.css']
})
export class ManufacturerEditComponent implements OnInit {

  angForm: FormGroup;
  public manufacturer: any = {};
  constructor(private route: ActivatedRoute, private router: Router, private ms: ManufacturerService, private fb: FormBuilder) {
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      ManufacturerName: ['', Validators.required],
      ManufacturerAddress: ['', Validators.required],
      ManufacturerContact: ['', Validators.required]
    });
  }

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.ms.editManufacturer(params.id).subscribe(manufacturer => {
        this.manufacturer = manufacturer;
      });
    });
  }

  updateManufacturer(name, address, contact) {
    this.route.params.subscribe(params => {
      this.ms.updateManufacturer(name, address, contact, params.id).subscribe(x => {
        this.router.navigate(['manufacturer/index']);
      });
    });
  }

  deleteManufacturer(id) {
    this.ms.deleteManufacturer(id).subscribe(res => {
      this.manufacturer.splice(id, 1);
    });
  }
}
