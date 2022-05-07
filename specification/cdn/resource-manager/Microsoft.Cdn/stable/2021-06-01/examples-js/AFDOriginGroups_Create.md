Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.0/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new origin group within the specified profile.
 *
 * @summary Creates a new origin group within the specified profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Create.json
 */
async function afdOriginGroupsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const originGroup = {
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
  const result = await client.afdOriginGroups.beginCreateAndWait(
    resourceGroupName,
    profileName,
    originGroupName,
    originGroup
  );
  console.log(result);
}

afdOriginGroupsCreate().catch(console.error);
```
