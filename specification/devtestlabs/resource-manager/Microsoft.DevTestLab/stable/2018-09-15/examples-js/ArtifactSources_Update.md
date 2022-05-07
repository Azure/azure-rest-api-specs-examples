Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Allows modifying tags of artifact sources. All other properties will be ignored.
 *
 * @summary Allows modifying tags of artifact sources. All other properties will be ignored.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Update.json
 */
async function artifactSourcesUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{artifactSourceName}";
  const artifactSource = {
    tags: { tagName1: "tagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.artifactSources.update(
    resourceGroupName,
    labName,
    name,
    artifactSource
  );
  console.log(result);
}

artifactSourcesUpdate().catch(console.error);
```
