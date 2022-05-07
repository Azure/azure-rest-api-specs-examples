Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates an invitation link for a user with the role
 *
 * @summary Description for Creates an invitation link for a user with the role
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateUserRolesInvitationLink.json
 */
async function createAnInvitationLinkForAUserForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const staticSiteUserRolesInvitationEnvelope = {
    domain: "happy-sea-15afae3e.azurestaticwebsites.net",
    numHoursToExpiration: 1,
    provider: "aad",
    roles: "admin,contributor",
    userDetails: "username",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.createUserRolesInvitationLink(
    resourceGroupName,
    name,
    staticSiteUserRolesInvitationEnvelope
  );
  console.log(result);
}

createAnInvitationLinkForAUserForAStaticSite().catch(console.error);
```
