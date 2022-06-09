```javascript
const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the cluster name is valid and is not already in use.
 *
 * @summary Checks that the cluster name is valid and is not already in use.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersCheckNameAvailability.json
 */
async function kustoClustersCheckNameAvailability() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const location = "westus";
  const clusterName = {
    name: "kustoCluster",
    type: "Microsoft.Kusto/clusters",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.checkNameAvailability(location, clusterName);
  console.log(result);
}

kustoClustersCheckNameAvailability().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-kusto_7.1.1/sdk/kusto/arm-kusto/README.md) on how to add the SDK to your project and authenticate.
