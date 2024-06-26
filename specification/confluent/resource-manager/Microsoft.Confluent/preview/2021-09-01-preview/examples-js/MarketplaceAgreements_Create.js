const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Confluent Marketplace agreement in the subscription.
 *
 * @summary Create Confluent Marketplace agreement in the subscription.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/MarketplaceAgreements_Create.json
 */
async function marketplaceAgreementsCreate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.marketplaceAgreements.create();
  console.log(result);
}

marketplaceAgreementsCreate().catch(console.error);
