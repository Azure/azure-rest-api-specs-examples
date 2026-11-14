const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the managed compute capacities for a subscription. Returns available capacity
 * per accelerator type, including deployment size information.
 *
 * @summary gets the managed compute capacities for a subscription. Returns available capacity
 * per accelerator type, including deployment size information.
 * x-ms-original-file: 2026-11-01/ListManagedComputeCapacitiesWithLocation.json
 */
async function listManagedComputeCapacitiesWithLocation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedComputeCapacities.list("MaaP", {
    acceleratorType: "H100_80GB",
    location: "eastus",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
