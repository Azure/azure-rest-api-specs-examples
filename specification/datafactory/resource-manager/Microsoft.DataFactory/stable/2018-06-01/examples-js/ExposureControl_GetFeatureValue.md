Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.6.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get exposure control feature for specific location.
 *
 * @summary Get exposure control feature for specific location.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ExposureControl_GetFeatureValue.json
 */
async function exposureControlGetFeatureValue() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const locationId = "WestEurope";
  const exposureControlRequest = {
    featureName: "ADFIntegrationRuntimeSharingRbac",
    featureType: "Feature",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.exposureControl.getFeatureValue(locationId, exposureControlRequest);
  console.log(result);
}

exposureControlGetFeatureValue().catch(console.error);
```
