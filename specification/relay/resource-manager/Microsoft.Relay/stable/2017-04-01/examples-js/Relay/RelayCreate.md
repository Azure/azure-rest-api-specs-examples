```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a WCF relay. This operation is idempotent.
 *
 * @summary Creates or updates a WCF relay. This operation is idempotent.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayCreate.json
 */
async function relayCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-9953";
  const relayName = "example-Relay-Wcf-1194";
  const parameters = {
    relayType: "NetTcp",
    requiresClientAuthorization: true,
    requiresTransportSecurity: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.wCFRelays.createOrUpdate(
    resourceGroupName,
    namespaceName,
    relayName,
    parameters
  );
  console.log(result);
}

relayCreate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.
