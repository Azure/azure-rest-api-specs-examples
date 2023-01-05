const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Azure credits and Microsoft Azure consumption commitments for a billing account or a billing profile. Microsoft Azure consumption commitments are only supported for the billing account scope.
 *
 * @summary Lists all Azure credits and Microsoft Azure consumption commitments for a billing account or a billing profile. Microsoft Azure consumption commitments are only supported for the billing account scope.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByBillingProfile.json
 */
async function lotsListByBillingProfile() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const billingAccountId = "1234:5678";
  const billingProfileId = "2468";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.lotsOperations.listByBillingProfile(
    billingAccountId,
    billingProfileId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
