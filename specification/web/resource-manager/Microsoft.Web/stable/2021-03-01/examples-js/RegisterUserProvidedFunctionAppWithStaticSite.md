Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Register a user provided function app with a static site
 *
 * @summary Description for Register a user provided function app with a static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/RegisterUserProvidedFunctionAppWithStaticSite.json
 */
async function registerAUserProvidedFunctionAppWithAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const functionAppName = "testFunctionApp";
  const isForced = true;
  const staticSiteUserProvidedFunctionEnvelope = {
    functionAppRegion: "West US 2",
    functionAppResourceId:
      "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp",
  };
  const options = {
    isForced,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.beginRegisterUserProvidedFunctionAppWithStaticSiteAndWait(
    resourceGroupName,
    name,
    functionAppName,
    staticSiteUserProvidedFunctionEnvelope,
    options
  );
  console.log(result);
}

registerAUserProvidedFunctionAppWithAStaticSite().catch(console.error);
```
