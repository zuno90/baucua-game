const d = {
  dice: "bau",
  amount: 1000,
}.toString();

const j = JSON.stringify({
  type: "BET",
  msg: {
    dice: "bau",
    amount: 1000,
  },
});

const u = JSON.parse(
  JSON.stringify({
    type: "GAMEINFO",
    msg: '{"stage":"BET TIME","no":835,"sec":17}',
  }),
);

console.log(d);
console.log(j);
console.log(u);
