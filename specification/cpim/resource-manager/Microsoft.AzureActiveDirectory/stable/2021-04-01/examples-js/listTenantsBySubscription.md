Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-azureadexternalidentities_1.0.0/sdk/azureadexternalidentities/arm-azureadexternalidentities/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ExternalIdentitiesConfigurationClient } = require("@azure/arm-azureadexternalidentities");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the Azure AD B2C tenant resources in a subscription.
 *
 * @summary Get all the Azure AD B2C tenant resources in a subscription.
 * x-ms-original-file: specification/cpim/resource-manager/Microsoft.AzureActiveDirectory/stable/2021-04-01/examples/listTenantsBySubscription.json
 */
async function b2CTenantsListBySubscription() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new ExternalIdentitiesConfigurationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.b2CTenants.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

b2CTenantsListBySubscription().catch(console.error);
```
