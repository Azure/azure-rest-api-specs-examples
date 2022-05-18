Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function exposureControlQueryFeatureValuesByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const exposureControlBatchRequest = {
    exposureControlRequests: [
      {
        featureName: "ADFIntegrationRuntimeSharingRbac",
        featureType: "Feature",
      },
      { featureName: "ADFSampleFeature", featureType: "Feature" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.exposureControl.queryFeatureValuesByFactory(
    resourceGroupName,
    factoryName,
    exposureControlBatchRequest
  );
  console.log(result);
}

exposureControlQueryFeatureValuesByFactory().catch(console.error);
```
