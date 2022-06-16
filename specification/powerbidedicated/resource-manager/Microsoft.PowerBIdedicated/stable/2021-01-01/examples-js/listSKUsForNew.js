const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists eligible SKUs for PowerBI Dedicated resource provider.
 *
 * @summary Lists eligible SKUs for PowerBI Dedicated resource provider.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listSKUsForNew.json
 */
async function listEligibleSkUsForANewCapacity() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.capacities.listSkus();
  console.log(result);
}

listEligibleSkUsForANewCapacity().catch(console.error);
