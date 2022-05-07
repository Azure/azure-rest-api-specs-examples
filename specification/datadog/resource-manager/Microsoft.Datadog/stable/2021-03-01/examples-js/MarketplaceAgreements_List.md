Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datadog_3.0.1/sdk/datadog/arm-datadog/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Datadog marketplace agreements in the subscription.
 *
 * @summary List Datadog marketplace agreements in the subscription.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/MarketplaceAgreements_List.json
 */
async function marketplaceAgreementsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.marketplaceAgreements.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

marketplaceAgreementsList().catch(console.error);
```
