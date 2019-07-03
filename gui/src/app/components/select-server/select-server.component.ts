/*
  Sliver Implant Framework
  Copyright (C) 2019  Bishop Fox
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { ConfigService } from '../../providers/config.service';
import { RPCConfig } from '../../../../rpc';
import { FADE_IN_OUT } from '../../shared/animations';

@Component({
  selector: 'app-select-server',
  templateUrl: './select-server.component.html',
  styleUrls: ['./select-server.component.scss'],
  animations: [FADE_IN_OUT]
})
export class SelectServerComponent implements OnInit {

  configs: RPCConfig[];
  selectedConfig: RPCConfig;
  connecting = false;

  selectConfigForm = new FormControl('', [Validators.required]);

  constructor(private _router: Router,
              private _configService: ConfigService) { }

  ngOnInit() {
    this.fetchConfigs();
  }

  onSelectedConfig(config: RPCConfig) {
    console.log(config);
    this.connecting = true;
    this._configService.setActiveConfig(config).then(() => {
      this._router.navigate(['/home']);
    }).catch((err) => {
      console.log(err);
    });
  }

  async fetchConfigs() {
    this.configs = await this._configService.listConfigs();
  }

}
