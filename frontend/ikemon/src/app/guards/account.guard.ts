import { CanActivateFn, Router } from '@angular/router';
import { UserService } from '../services/user.service';

export const accountGuard: CanActivateFn = (route, state) => {
  let userService = new UserService();
  let user = userService.GetUser();
  if (!user) {
    let router = new Router()
    router.navigate(["/"])
  }
  return !!user
};
