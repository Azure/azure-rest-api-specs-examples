const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all availability sets in a resource group.
 *
 * @summary lists all availability sets in a resource group.
 * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_List_MaximumSet_Gen.json
 */
async function availabilitySetListMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.availabilitySets.list("rgcompute")) {
    resArray.push(item);
  }

  console.log(resArray);
}
