Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resourcehealth_3.0.1/sdk/resourcehealth/arm-resourcehealth/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the current availability status for all the resources in the subscription. Use the nextLink property in the response to get the next page of availability statuses.
 *
 * @summary Lists the current availability status for all the resources in the subscription. Use the nextLink property in the response to get the next page of availability statuses.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2017-07-01/examples/AvailabilityStatuses_ListBySubscriptionId.json
 */
async function listHealthBySubscriptionId() {
  const subscriptionId = "subscriptionId";
  const expand = "recommendedactions";
  const options = {
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilityStatuses.listBySubscriptionId(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listHealthBySubscriptionId().catch(console.error);
```
