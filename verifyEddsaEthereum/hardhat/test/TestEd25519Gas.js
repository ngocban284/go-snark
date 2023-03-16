const { ethers } = require("hardhat");
const { expect } = require("chai");
let ed25519, Ed25519;
let libEd25519, LibEd25519;

beforeEach(async function () {
  LibEd25519 = await ethers.getContractFactory("Ed25519");
  libEd25519 = await LibEd25519.deploy();
  await libEd25519.deployed();

  Ed25519 = await ethers.getContractFactory("TestEd25519", {
    libraries: {
      Ed25519: libEd25519.address,
    },
  });
  ed25519 = await Ed25519.deploy();
  await ed25519.deployed();
});

describe("Ed25519", function () {
  it("Contract deployed", async function () {
    expect(ed25519.address).to.not.equal(0);
  });

  it("Verify EdDSA gas", async function () {
    let tx = await ed25519.testVerify();
    let receipt = await tx.wait();
    // print tx gas price
    console.log("Gas price: " + receipt.effectiveGasPrice.toString());
    console.log("Gas used: " + receipt.gasUsed.toString());
    // amount of ether spent on gas
    console.log(
      "Ether spent on gas: " +
        receipt.effectiveGasPrice.mul(receipt.gasUsed).toString() / 1e18
    );
  });
});
