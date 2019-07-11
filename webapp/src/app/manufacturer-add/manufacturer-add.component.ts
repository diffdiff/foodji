import { Component, OnInit } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ManufacturerService } from '../manufacturer.service';
@Component({
  selector: 'app-manufacturer-add',
  templateUrl: './manufacturer-add.component.html',
  styleUrls: ['./manufacturer-add.component.css']
})
export class ManufacturerAddComponent implements OnInit {

  angForm: FormGroup;

  constructor(private route: ActivatedRoute, private router: Router, private fb: FormBuilder, private ms: ManufacturerService) {
    this.createForm();
  }

  createForm() {
    this.angForm = this.fb.group({
      ManufacturerName: ['', Validators.required ],
      ManufacturerAddress: ['', Validators.required ],
      ManufacturerContact: ['', Validators.required ]
    });
  }



  addManufacturer(name, address, contact) {
    this.route.params.subscribe(params => {
      this.ms.addManufacturer(name, address, contact);
      this.router.navigate(['manufacturer/index']);
    });
  }

  ngOnInit() {
  }
}
