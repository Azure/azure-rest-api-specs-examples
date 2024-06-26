const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the virtual machine template resource.
 *
 * @summary API to update certain properties of the virtual machine template resource.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/UpdateVirtualMachineTemplate.json
 */
async function updateVirtualMachineTemplate() {
  const subscriptionId =
    process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "testrg";
  const virtualMachineTemplateName = "WebFrontEndTemplate";
  const body = { tags: { tag1: "value1", tag2: "value2" } };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.virtualMachineTemplates.update(
    resourceGroupName,
    virtualMachineTemplateName,
    options
  );
  console.log(result);
}
