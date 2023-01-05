const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of HybridIdentityMetadata of the given vm.
 *
 * @summary Returns the list of HybridIdentityMetadata of the given vm.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/HybridIdentityMetadata_ListByVm.json
 */
async function hybridIdentityMetadataListByVM() {
  const subscriptionId =
    process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "testrg";
  const virtualMachineName = "ContosoVm";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.hybridIdentityMetadataOperations.listByVm(
    resourceGroupName,
    virtualMachineName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
