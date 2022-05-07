Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function listsThePrefixesReceivedOverTheSpecifiedPeeringUnderTheGivenSubscriptionAndResourceGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const prefix = "1.1.1.0/24";
  const asPath = "123 456";
  const originAsValidationState = "Valid";
  const rpkiValidationState = "Valid";
  const options = {
    prefix: prefix,
    asPath: asPath,
    originAsValidationState: originAsValidationState,
    rpkiValidationState: rpkiValidationState,
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.receivedRoutes.listByPeering(
    resourceGroupName,
    peeringName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsThePrefixesReceivedOverTheSpecifiedPeeringUnderTheGivenSubscriptionAndResourceGroup().catch(
  console.error
);
```
