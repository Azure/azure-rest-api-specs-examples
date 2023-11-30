const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create data product resource.
 *
 * @summary Create data product resource.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Create_MinimumSet_Gen.json
 */
async function dataProductsCreateMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const resource = {
    location: "eastus",
    properties: {
      majorVersion: "1.0.0",
      product: "MCC",
      publisher: "Microsoft",
    },
    tags: { userSpecifiedKeyName: "userSpecifiedKeyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataProducts.beginCreateAndWait(
    resourceGroupName,
    dataProductName,
    resource
  );
  console.log(result);
}
