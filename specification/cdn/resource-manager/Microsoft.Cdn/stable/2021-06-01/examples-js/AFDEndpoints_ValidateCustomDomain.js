const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates the custom domain mapping to ensure it maps to the correct CDN endpoint in DNS.
 *
 * @summary Validates the custom domain mapping to ensure it maps to the correct CDN endpoint in DNS.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ValidateCustomDomain.json
 */
async function endpointsValidateCustomDomain() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const customDomainProperties = {
    hostName: "www.someDomain.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdEndpoints.validateCustomDomain(
    resourceGroupName,
    profileName,
    endpointName,
    customDomainProperties
  );
  console.log(result);
}

endpointsValidateCustomDomain().catch(console.error);
