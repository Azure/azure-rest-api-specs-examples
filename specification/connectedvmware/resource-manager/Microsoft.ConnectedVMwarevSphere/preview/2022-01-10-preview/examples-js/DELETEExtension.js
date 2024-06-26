const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the extension.
 *
 * @summary The operation to delete the extension.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DELETEExtension.json
 */
async function deleteAMachineExtension() {
  const subscriptionId = process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "myResourceGroup";
  const name = "myMachine";
  const extensionName = "MMA";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.machineExtensions.beginDeleteAndWait(
    resourceGroupName,
    name,
    extensionName
  );
  console.log(result);
}
