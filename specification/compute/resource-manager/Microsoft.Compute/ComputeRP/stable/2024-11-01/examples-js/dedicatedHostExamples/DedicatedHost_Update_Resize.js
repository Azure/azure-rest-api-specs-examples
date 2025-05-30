const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update a dedicated host .
 *
 * @summary Update a dedicated host .
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/dedicatedHostExamples/DedicatedHost_Update_Resize.json
 */
async function dedicatedHostUpdateResize() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const hostGroupName = "aaaaaaaaa";
  const hostName = "aaaaaaaaaaaaaaaaaaaaa";
  const parameters = { sku: { name: "DSv3-Type1" } };
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
