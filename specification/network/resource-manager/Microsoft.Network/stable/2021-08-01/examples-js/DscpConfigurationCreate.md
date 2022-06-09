```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DSCP Configuration.
 *
 * @summary Creates or updates a DSCP Configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DscpConfigurationCreate.json
 */
async function createDscpConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const dscpConfigurationName = "mydscpconfig";
  const parameters = {
    location: "eastus",
    qosDefinitionCollection: [
      {
        destinationIpRanges: [{ endIP: "127.0.10.2", startIP: "127.0.10.1" }],
        destinationPortRanges: [{ end: 15, start: 15 }],
        markings: [1],
        sourceIpRanges: [{ endIP: "127.0.0.2", startIP: "127.0.0.1" }],
        sourcePortRanges: [
          { end: 11, start: 10 },
          { end: 21, start: 20 },
        ],
        protocol: "Tcp",
      },
      {
        destinationIpRanges: [{ endIP: "12.0.10.2", startIP: "12.0.10.1" }],
        destinationPortRanges: [{ end: 52, start: 51 }],
        markings: [2],
        sourceIpRanges: [{ endIP: "12.0.0.2", startIP: "12.0.0.1" }],
        sourcePortRanges: [{ end: 12, start: 11 }],
        protocol: "Udp",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.dscpConfigurationOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    dscpConfigurationName,
    parameters
  );
  console.log(result);
}

createDscpConfiguration().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
