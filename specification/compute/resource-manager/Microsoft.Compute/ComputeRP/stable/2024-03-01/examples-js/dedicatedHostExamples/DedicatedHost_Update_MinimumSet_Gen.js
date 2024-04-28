const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a dedicated host .
 *
 * @summary Update a dedicated host .
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHost_Update_MinimumSet_Gen.json
 */
async function dedicatedHostUpdateMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const hostGroupName = "aa";
  const hostName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.beginUpdateAndWait(
    resourceGroupName,
    hostGroupName,
    hostName,
    parameters,
  );
  console.log(result);
}
