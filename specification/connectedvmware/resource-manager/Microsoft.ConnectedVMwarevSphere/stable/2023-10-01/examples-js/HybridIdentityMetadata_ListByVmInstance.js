const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of HybridIdentityMetadata of the given vm.
 *
 * @summary Returns the list of HybridIdentityMetadata of the given vm.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/HybridIdentityMetadata_ListByVmInstance.json
 */
async function hybridIdentityMetadataListByVM() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential);
  const resArray = new Array();
  for await (let item of client.vmInstanceHybridIdentityMetadataOperations.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
