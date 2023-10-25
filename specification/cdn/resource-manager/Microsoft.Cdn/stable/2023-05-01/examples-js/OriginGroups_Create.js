const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new origin group within the specified endpoint.
 *
 * @summary Creates a new origin group within the specified endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/OriginGroups_Create.json
 */
async function originGroupsCreate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originGroupName = "origingroup1";
  const originGroup = {
    healthProbeSettings: {
      probeIntervalInSeconds: 120,
      probePath: "/health.aspx",
      probeProtocol: "Http",
      probeRequestType: "GET",
    },
    origins: [
      {
        id: "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1",
      },
    ],
    responseBasedOriginErrorDetectionSettings: {
      responseBasedDetectedErrorTypes: "TcpErrorsOnly",
      responseBasedFailoverThresholdPercentage: 10,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.originGroups.beginCreateAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    originGroupName,
    originGroup
  );
  console.log(result);
}
