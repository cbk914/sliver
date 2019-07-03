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

import { trigger, state, style, animate, transition } from '@angular/animations';

export const FADE_IN_OUT = trigger('fadeInOut', [
  transition(':enter', [   // :enter is alias to 'void => *'
    style({opacity: 0}),
    animate(400, style({opacity: 1}))
  ]),
  transition(':leave', [   // :leave is alias to '* => void'
    animate(400, style({opacity: 0}))
  ])
]);
