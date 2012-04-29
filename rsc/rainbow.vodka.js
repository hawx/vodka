Rainbow.extend('vodka', [
  {
      'name': 'string',
      'pattern': /('|")[\s\S]*?('|")/gm
  }, {
    'name': 'comment',
    'pattern': /;.*$/gm
  }, {
    'name': 'integer',
    'pattern': /\b\d+\b/g
  }, {
    'name': 'constant.language',
    'pattern': /true|false|nil/g
  }, {
      'name': 'support',
      'pattern': /:\w+/g
  }
], true);
