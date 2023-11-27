const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates an invitation link for a user with the role
 *
 * @summary Description for Creates an invitation link for a user with the role
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/CreateUserRolesInvitationLink.json
 */
async function createAnInvitationLinkForAUserForAStaticSite() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
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
