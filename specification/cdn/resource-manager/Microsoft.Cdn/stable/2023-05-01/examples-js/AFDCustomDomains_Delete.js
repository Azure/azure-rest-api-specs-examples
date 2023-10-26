const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
 *
 * @summary Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDCustomDomains_Delete.json
 */
async function afdCustomDomainsDelete() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const customDomainName = "domain1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdCustomDomains.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    customDomainName
  );
  console.log(result);
}
