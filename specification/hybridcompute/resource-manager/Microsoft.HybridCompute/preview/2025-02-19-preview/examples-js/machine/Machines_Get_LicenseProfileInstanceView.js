const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves information about the model view or the instance view of a hybrid machine.
 *
 * @summary Retrieves information about the model view or the instance view of a hybrid machine.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/machine/Machines_Get_LicenseProfileInstanceView.json
 */
async function getMachineWithLicenseProfileInstanceView() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const machineName = "myMachine";
  const expand = "instanceView";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machines.get(resourceGroupName, machineName, options);
  console.log(result);
}
