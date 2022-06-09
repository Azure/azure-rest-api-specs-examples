```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of all possible traffic between resources for the subscription
 *
 * @summary Gets the list of all possible traffic between resources for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/GetAllowedConnectionsSubscription_example.json
 */
async function getAllowedConnectionsOnASubscription() {
  const subscriptionId = "3eeab341-f466-499c-a8be-85427e154bad";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.allowedConnections.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllowedConnectionsOnASubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
