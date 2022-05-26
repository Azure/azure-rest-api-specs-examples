Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-azurestackhci_3.0.0/sdk/azurestackhci/arm-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get particular Arc Extension of HCI Cluster.
 *
 * @summary Get particular Arc Extension of HCI Cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/GetExtension.json
 */
async function getArcSettingsExtension() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const extensionName = "MicrosoftMonitoringAgent";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.extensions.get(
    resourceGroupName,
    clusterName,
    arcSettingName,
    extensionName
  );
  console.log(result);
}

getArcSettingsExtension().catch(console.error);
```
