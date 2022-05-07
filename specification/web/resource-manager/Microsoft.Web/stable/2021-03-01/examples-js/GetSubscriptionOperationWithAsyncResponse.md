Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets an operation in a subscription and given region
 *
 * @summary Description for Gets an operation in a subscription and given region
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetSubscriptionOperationWithAsyncResponse.json
 */
async function getsAnOperationInASubscriptionAndGivenRegion() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "West US";
  const operationId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab5d5";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.global.getSubscriptionOperationWithAsyncResponse(
    location,
    operationId
  );
  console.log(result);
}

getsAnOperationInASubscriptionAndGivenRegion().catch(console.error);
```
