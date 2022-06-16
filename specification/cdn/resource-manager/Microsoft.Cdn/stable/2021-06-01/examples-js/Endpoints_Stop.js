const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops an existing running CDN endpoint.
 *
 * @summary Stops an existing running CDN endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Stop.json
 */
async function endpointsStop() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginStopAndWait(
    resourceGroupName,
    profileName,
    endpointName
  );
  console.log(result);
}

endpointsStop().catch(console.error);
