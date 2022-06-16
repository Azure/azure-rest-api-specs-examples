const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 *
 * @summary Deletes an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Delete.json
 */
async function endpointsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    endpointName
  );
  console.log(result);
}

endpointsDelete().catch(console.error);
