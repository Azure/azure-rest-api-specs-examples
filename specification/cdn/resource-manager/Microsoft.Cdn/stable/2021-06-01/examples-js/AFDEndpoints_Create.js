const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 *
 * @summary Creates a new AzureFrontDoor endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_Create.json
 */
async function afdEndpointsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const endpoint = {
    autoGeneratedDomainNameLabelScope: "TenantReuse",
    enabledState: "Enabled",
    location: "CentralUs",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdEndpoints.beginCreateAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    endpoint
  );
  console.log(result);
}

afdEndpointsCreate().catch(console.error);
