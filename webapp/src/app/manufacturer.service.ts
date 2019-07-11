import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import Manufacturer from './Manufacturer';

@Injectable({
  providedIn: 'root'
})
export class ManufacturerService {

  uri = 'http://localhost:3000/api/v1/manufacturers';

  constructor(private http: HttpClient) { }

  addManufacturer(name, address, contact) {
    const manufacturer = {
      name,
      address,
      contact
    };
    this.http.post(`${this.uri}`, manufacturer)
    .subscribe(res => console.log('Done'));
  }

  getManufacturer() {
    return this.http.get(`${this.uri}`)
    .pipe(
      map((array: any[]) => {
        return array.map(manufacturer => {
          return {
            ManufacturerID: manufacturer.ID,
            ManufacturerName: manufacturer.name,
            ManufacturerAddress: manufacturer.address,
            ManufacturerContact: manufacturer.contacts
          } as Manufacturer;
        });
      }));
  }

  editManufacturer(id) {
    return this
      .http
      .get(`${this.uri}/${id}`)
      .pipe(map((manufacturer: any) => {
        return {
          ManufacturerID: manufacturer.ID,
          ManufacturerName: manufacturer.name,
          ManufacturerAddress: manufacturer.address,
          ManufacturerContact: manufacturer.contacts
        } as Manufacturer;
      }));
  }

  updateManufacturer(name, address, contact, id) {
    const manufacturer = {
      name,
      address,
      contacts: contact
    };
    return this
      .http
      .put(`${this.uri}/${id}`, manufacturer);
  }

  deleteManufacturer(id) {
    return this
      .http
      .delete(`${this.uri}/${id}`);
  }
}
