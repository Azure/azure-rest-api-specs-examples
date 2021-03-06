const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts an existing CDN endpoint that is on a stopped state.
 *
 * @summary Starts an existing CDN endpoint that is on a stopped state.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Start.json
 */
async function endpointsStart() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginStartAndWait(
    resourceGroupName,
    profileName,
    endpointName
  );
  console.log(result);
}

endpointsStart().catch(console.error);
