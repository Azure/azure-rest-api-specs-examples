Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates or updates an App Service Plan.
 *
 * @summary Description for Creates or updates an App Service Plan.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateAppServicePlan.json
 */
async function createOrUpdateAppServicePlan() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "testsf6141";
  const appServicePlan = {
    kind: "app",
    location: "East US",
    sku: { name: "P1", capacity: 1, family: "P", size: "P1", tier: "Premium" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServicePlans.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    appServicePlan
  );
  console.log(result);
}

createOrUpdateAppServicePlan().catch(console.error);
```
