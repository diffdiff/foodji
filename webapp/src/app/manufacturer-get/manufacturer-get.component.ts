import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ManufacturerService } from '../manufacturer.service';
import Manufacturer from '../Manufacturer';

@Component({
  selector: 'app-manufacturer-get',
  templateUrl: './manufacturer-get.component.html',
  styleUrls: ['./manufacturer-get.component.css']
})
export class ManufacturerGetComponent implements OnInit {

  Manufacturers: Manufacturer[];
  constructor(private route: ActivatedRoute, private router: Router, private ms: ManufacturerService) {
    this.router.routeReuseStrategy.shouldReuseRoute = () => {
      return false;
    };
  }

  ngOnInit() {
    this.getManufacturer();
  }

  getManufacturer() {
    this.ms.getManufacturer().subscribe(data => {
      this.Manufacturers = data;
    });
  }

  deleteManufacturer(id) {
    this.ms.deleteManufacturer(id).subscribe(res => {
      this.Manufacturers = this.Manufacturers.filter(m => m.ManufacturerID !== id);
    });
  }
}
