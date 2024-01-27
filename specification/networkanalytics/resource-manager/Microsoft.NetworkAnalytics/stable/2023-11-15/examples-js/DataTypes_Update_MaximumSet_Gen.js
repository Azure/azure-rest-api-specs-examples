const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update data type resource.
 *
 * @summary Update data type resource.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataTypes_Update_MaximumSet_Gen.json
 */
async function dataTypesUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const dataTypeName = "datatypename";
  const properties = {
    properties: {
      databaseCacheRetention: 16,
      databaseRetention: 9,
      state: "STARTED",
      storageOutputRetention: 30,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataTypes.beginUpdateAndWait(
    resourceGroupName,
    dataProductName,
    dataTypeName,
    properties,
  );
  console.log(result);
}
