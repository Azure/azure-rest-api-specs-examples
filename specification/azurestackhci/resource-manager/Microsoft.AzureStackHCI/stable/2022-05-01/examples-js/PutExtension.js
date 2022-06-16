const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Extension for HCI cluster.
 *
 * @summary Create Extension for HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/PutExtension.json
 */
async function createArcExtension() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const extensionName = "MicrosoftMonitoringAgent";
  const extension = {
    typePropertiesExtensionParametersType: "MicrosoftMonitoringAgent",
    protectedSettings: { workspaceKey: "xx" },
    publisher: "Microsoft.Compute",
    settings: { workspaceId: "xx" },
    typeHandlerVersion: "1.10",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.extensions.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    arcSettingName,
    extensionName,
    extension
  );
  console.log(result);
}

createArcExtension().catch(console.error);
