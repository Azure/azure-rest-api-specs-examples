```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Deletes a custom domain.
 *
 * @summary Description for Deletes a custom domain.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteStaticSiteCustomDomain.json
 */
async function deleteACustomDomainForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const domainName = "custom.domain.net";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.beginDeleteStaticSiteCustomDomainAndWait(
    resourceGroupName,
    name,
    domainName
  );
  console.log(result);
}

deleteACustomDomainForAStaticSite().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.
