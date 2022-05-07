Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a cloud account connector from a subscription
 *
 * @summary Delete a cloud account connector from a subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/DeleteConnectorSubscription_example.json
 */
async function deleteACloudAccountConnectorFromASubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const connectorName = "aws_dev1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.connectors.delete(connectorName);
  console.log(result);
}

deleteACloudAccountConnectorFromASubscription().catch(console.error);
```
