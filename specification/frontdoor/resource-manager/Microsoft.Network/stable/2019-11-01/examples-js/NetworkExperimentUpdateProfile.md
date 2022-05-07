Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-frontdoor_5.0.1/sdk/frontdoor/arm-frontdoor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an NetworkExperimentProfiles
 *
 * @summary Updates an NetworkExperimentProfiles
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentUpdateProfile.json
 */
async function updatesAnExperiment() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const parameters = {
    enabledState: "Enabled",
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.networkExperimentProfiles.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    parameters
  );
  console.log(result);
}

updatesAnExperiment().catch(console.error);
```
