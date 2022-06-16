const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists Security Center pricing configurations in the subscription.
 *
 * @summary Lists Security Center pricing configurations in the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-03-01/examples/Pricings/ListPricings_example.json
 */
async function getPricingsOnSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.pricings.list();
  console.log(result);
}

getPricingsOnSubscription().catch(console.error);
