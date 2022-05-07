Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-communication_3.0.1/sdk/communication/arm-communication/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate CommunicationService access key. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
 *
 * @summary Regenerate CommunicationService access key. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/regenerateKey.json
 */
async function regenerateKey() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = { keyType: "Primary" };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationService.regenerateKey(
    resourceGroupName,
    communicationServiceName,
    parameters
  );
  console.log(result);
}

regenerateKey().catch(console.error);
```
