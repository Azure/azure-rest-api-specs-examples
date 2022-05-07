Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-azureadexternalidentities_1.0.0/sdk/azureadexternalidentities/arm-azureadexternalidentities/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ExternalIdentitiesConfigurationClient } = require("@azure/arm-azureadexternalidentities");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initiates an async operation to delete the Azure AD B2C tenant and Azure resource. The resource deletion can only happen as the last step in [the tenant deletion process](https://aka.ms/deleteB2Ctenant).
 *
 * @summary Initiates an async operation to delete the Azure AD B2C tenant and Azure resource. The resource deletion can only happen as the last step in [the tenant deletion process](https://aka.ms/deleteB2Ctenant).
 * x-ms-original-file: specification/cpim/resource-manager/Microsoft.AzureActiveDirectory/stable/2021-04-01/examples/deleteTenant.json
 */
async function deleteTenant() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab346";
  const resourceGroupName = "rg1";
  const resourceName = "contoso.onmicrosoft.com";
  const credential = new DefaultAzureCredential();
  const client = new ExternalIdentitiesConfigurationClient(credential, subscriptionId);
  const result = await client.b2CTenants.beginDeleteAndWait(resourceGroupName, resourceName);
  console.log(result);
}

deleteTenant().catch(console.error);
```
