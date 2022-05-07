Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Detach the user provided function app from the static site build
 *
 * @summary Description for Detach the user provided function app from the static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DetachUserProvidedFunctionAppFromStaticSiteBuild.json
 */
async function detachTheUserProvidedFunctionAppFromTheStaticSiteBuild() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const environmentName = "12";
  const functionAppName = "testFunctionApp";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.detachUserProvidedFunctionAppFromStaticSiteBuild(
    resourceGroupName,
    name,
    environmentName,
    functionAppName
  );
  console.log(result);
}

detachTheUserProvidedFunctionAppFromTheStaticSiteBuild().catch(console.error);
```
