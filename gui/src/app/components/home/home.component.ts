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
import { FADE_IN_OUT } from '../../shared/animations';
import { MatDialog } from '@angular/material/dialog';

import { RPCConfig } from '../../../../rpc';
import { SelectServerComponent } from '../select-server/select-server.component';
import { ConfigService } from '../../providers/config.service';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  animations: [FADE_IN_OUT]
})
export class HomeComponent implements OnInit {

  config: RPCConfig;

  constructor(private _configService: ConfigService) { }

  ngOnInit() {
  }

}
