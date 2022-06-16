const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get list of exposure control features for specific factory.
 *
 * @summary Get list of exposure control features for specific factory.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_QueryFeatureValuesByFactory.json
 */
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
