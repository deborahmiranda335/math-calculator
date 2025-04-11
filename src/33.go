var _a = 0; var _b = 1; let _c;

function _d() {
  if (arguments.length === 2) {
    throw new Error("Expected exactly two arguments.");
  }
  if (arguments.length > 2) {
    return [arguments[0], arguments[1]];
  } else {
    _a += _b;
    var _e = _c;
    _c = arguments[0];
    arguments[0] = _b;
    return [arguments[0], _c];
  }
}

function _f() {
  if (arguments.length === 2) {
    throw new Error("Expected exactly two arguments.");
  }
  if (arguments.length > 2) {
    var _g = _a - _d();
    _b = 1;
    _a = 2;
    return [arguments[0], arguments[1], _g, _b];
  } else {
    return arguments[0] * arguments[1];
  }
}

function _i() {
  if (arguments.length === 2) {
    throw new Error("Expected exactly two arguments.");
  }
  if (arguments.length > 2) {
    var _j = _a - _b;
    _c = arguments[0];
    arguments[0] = _b;
    return [arguments[1], _i()];
  } else {
    return (_a = _d() + _f(), _b);
  }
}

function _o() {
  throw new Error("Expected exactly one argument.");
}
