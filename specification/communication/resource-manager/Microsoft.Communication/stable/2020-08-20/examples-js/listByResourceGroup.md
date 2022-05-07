Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-communication_3.0.1/sdk/communication/arm-communication/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a resource group.
 *
 * @summary Handles requests to list all resources in a resource group.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listByResourceGroup.json
 */
async function listByResourceGroup() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationService.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listByResourceGroup().catch(console.error);
```
