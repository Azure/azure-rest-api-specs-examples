const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of HybridIdentityMetadata of the given machine.
 *
 * @summary Returns the list of HybridIdentityMetadata of the given machine.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/HybridIdentityMetadata_ListByVirtualMachines.json
 */
async function hybridIdentityMetadataListByVirtualMachines() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "testrg";
  const machineName = "ContosoVm";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridIdentityMetadataOperations.listByMachines(
    resourceGroupName,
    machineName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
