const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing origin group within a profile.
 *
 * @summary Updates an existing origin group within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOriginGroups_Update.json
 */
async function afdOriginGroupsUpdate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const originGroupUpdateProperties = {
    healthProbeSettings: {
      probeIntervalInSeconds: 10,
      probePath: "/path2",
      probeProtocol: "NotSet",
      probeRequestType: "NotSet",
    },
    loadBalancingSettings: {
      additionalLatencyInMilliseconds: 1000,
      sampleSize: 3,
      successfulSamplesRequired: 3,
    },
    trafficRestorationTimeToHealedOrNewEndpointsInMinutes: 5,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOriginGroups.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    originGroupName,
    originGroupUpdateProperties,
  );
  console.log(result);
}
