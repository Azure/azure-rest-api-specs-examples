const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates an ownership identifier for a domain or updates identifier details for an existing identifier
 *
 * @summary Description for Creates an ownership identifier for a domain or updates identifier details for an existing identifier
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/UpdateAppServiceDomainOwnershipIdentifier.json
 */
async function updateAppServiceDomainOwnershipIdentifier() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const domainName = "example.com";
  const name = "SampleOwnershipId";
  const domainOwnershipIdentifier = {
    ownershipId: "SampleOwnershipId",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.domains.updateOwnershipIdentifier(
    resourceGroupName,
    domainName,
    name,
    domainOwnershipIdentifier
  );
  console.log(result);
}

updateAppServiceDomainOwnershipIdentifier().catch(console.error);
