const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes a content from AzureFrontDoor.
 *
 * @summary Removes a content from AzureFrontDoor.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_PurgeContent.json
 */
async function afdEndpointsPurgeContent() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const contents = {
    contentPaths: ["/folder1"],
    domains: ["endpoint1.azureedge.net"],
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdEndpoints.beginPurgeContentAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    contents
  );
  console.log(result);
}

afdEndpointsPurgeContent().catch(console.error);
