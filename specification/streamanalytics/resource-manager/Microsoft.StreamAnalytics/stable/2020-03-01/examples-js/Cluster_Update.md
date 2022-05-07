Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing cluster. This can be used to partially update (ie. update one or two properties) a cluster without affecting the rest of the cluster definition.
 *
 * @summary Updates an existing cluster. This can be used to partially update (ie. update one or two properties) a cluster without affecting the rest of the cluster definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Cluster_Update.json
 */
async function updateACluster() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "sjrg";
  const clusterName = "testcluster";
  const cluster = { location: "Central US", sku: { capacity: 96 } };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAndWait(resourceGroupName, clusterName, cluster);
  console.log(result);
}

updateACluster().catch(console.error);
```
