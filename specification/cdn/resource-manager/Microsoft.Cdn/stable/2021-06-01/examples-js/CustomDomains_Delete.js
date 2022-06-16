const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing custom domain within an endpoint.
 *
 * @summary Deletes an existing custom domain within an endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_Delete.json
 */
async function customDomainsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const customDomainName = "www-someDomain-net";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    customDomainName
  );
  console.log(result);
}

customDomainsDelete().catch(console.error);
