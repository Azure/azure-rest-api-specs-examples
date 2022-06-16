const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the software inventory of all virtual machines in the subscriptions.
 *
 * @summary Gets the software inventory of all virtual machines in the subscriptions.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/ListBySubscriptionSoftwareInventories_example.json
 */
async function getsTheSoftwareInventoryOfAllVirtualMachinesInTheSubscriptions() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.softwareInventories.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsTheSoftwareInventoryOfAllVirtualMachinesInTheSubscriptions().catch(console.error);
