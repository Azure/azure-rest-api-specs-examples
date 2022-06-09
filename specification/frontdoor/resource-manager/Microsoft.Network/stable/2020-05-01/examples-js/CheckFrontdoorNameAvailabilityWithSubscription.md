```javascript
const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the availability of a Front Door subdomain.
 *
 * @summary Check the availability of a Front Door subdomain.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/CheckFrontdoorNameAvailabilityWithSubscription.json
 */
async function checkNameAvailabilityWithSubscription() {
  const subscriptionId = "subid";
  const checkFrontDoorNameAvailabilityInput = {
    name: "sampleName",
    type: "Microsoft.Network/frontDoors/frontendEndpoints",
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.frontDoorNameAvailabilityWithSubscription.check(
    checkFrontDoorNameAvailabilityInput
  );
  console.log(result);
}

checkNameAvailabilityWithSubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-frontdoor_5.0.1/sdk/frontdoor/arm-frontdoor/README.md) on how to add the SDK to your project and authenticate.
