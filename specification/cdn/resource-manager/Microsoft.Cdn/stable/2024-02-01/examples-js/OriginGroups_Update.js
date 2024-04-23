const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing origin group within an endpoint.
 *
 * @summary Updates an existing origin group within an endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/OriginGroups_Update.json
 */
async function originGroupsUpdate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originGroupName = "originGroup1";
  const originGroupUpdateProperties = {
    healthProbeSettings: {
      probeIntervalInSeconds: 120,
      probePath: "/health.aspx",
      probeProtocol: "Http",
      probeRequestType: "GET",
    },
    origins: [
      {
        id: "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.originGroups.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    originGroupName,
    originGroupUpdateProperties,
  );
  console.log(result);
}
