const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Confluent marketplace agreements in the subscription.
 *
 * @summary List Confluent marketplace agreements in the subscription.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/MarketplaceAgreements_List.json
 */
async function marketplaceAgreementsList() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.marketplaceAgreements.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
