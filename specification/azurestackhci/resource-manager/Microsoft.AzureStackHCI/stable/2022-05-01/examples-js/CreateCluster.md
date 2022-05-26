Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-azurestackhci_3.0.0/sdk/azurestackhci/arm-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create an HCI cluster.
 *
 * @summary Create an HCI cluster.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateCluster.json
 */
async function createCluster() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const cluster = {
    aadClientId: "24a6e53d-04e5-44d2-b7cc-1b732a847dfc",
    aadTenantId: "7e589cc1-a8b6-4dff-91bd-5ec0fa18db94",
    cloudManagementEndpoint: "https://98294836-31be-4668-aeae-698667faf99b.waconazure.com",
    location: "East US",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusters.create(resourceGroupName, clusterName, cluster);
  console.log(result);
}

createCluster().catch(console.error);
```
