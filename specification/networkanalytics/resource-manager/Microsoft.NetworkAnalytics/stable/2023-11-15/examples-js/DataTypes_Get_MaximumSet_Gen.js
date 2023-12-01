const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve data type resource.
 *
 * @summary Retrieve data type resource.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataTypes_Get_MaximumSet_Gen.json
 */
async function dataTypesGetMaximumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const dataTypeName = "datatypename";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataTypes.get(resourceGroupName, dataProductName, dataTypeName);
  console.log(result);
}
