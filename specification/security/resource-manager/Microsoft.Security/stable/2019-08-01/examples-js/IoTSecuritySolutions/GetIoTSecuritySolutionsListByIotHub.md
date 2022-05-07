Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get the list of IoT Security solutions by subscription.
 *
 * @summary Use this method to get the list of IoT Security solutions by subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHub.json
 */
async function listIoTSecuritySolutionsByIoTHub() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const filter =
    'properties.iotHubs/any(i eq "/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")';
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotSecuritySolution.listBySubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listIoTSecuritySolutionsByIoTHub().catch(console.error);
```
