const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Redeem invite code to join a redeemable lab
 *
 * @summary Redeem invite code to join a redeemable lab
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/RedeemCode.json
 */
async function redeemCode() {
  const parameters = {
    firstName: "test",
    lastName: "user",
    redeemCode: "exampleRedeemCode",
  };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const result = await client.redeemInvitationCode(parameters);
  console.log(result);
}

redeemCode().catch(console.error);
