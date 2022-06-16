const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get exposure control feature for specific factory.
 *
 * @summary Get exposure control feature for specific factory.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_GetFeatureValueByFactory.json
 */
async function exposureControlGetFeatureValueByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const exposureControlRequest = {
    featureName: "ADFIntegrationRuntimeSharingRbac",
    featureType: "Feature",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.exposureControl.getFeatureValueByFactory(
    resourceGroupName,
    factoryName,
    exposureControlRequest
  );
  console.log(result);
}

exposureControlGetFeatureValueByFactory().catch(console.error);
