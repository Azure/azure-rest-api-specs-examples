const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update data product resource.
 *
 * @summary Update data product resource.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Update_MaximumSet_Gen.json
 */
async function dataProductsUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const properties = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subid/resourceGroups/resourceGroupName/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
      },
    },
    properties: {
      currentMinorVersion: "1.0.1",
      owners: ["abc@micros.com", "def@micros.com"],
      privateLinksEnabled: "Disabled",
      purviewAccount: "testpurview",
      purviewCollection: "134567890",
    },
    tags: { userSpecifiedKeyName: "userSpecifiedKeyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataProducts.beginUpdateAndWait(
    resourceGroupName,
    dataProductName,
    properties
  );
  console.log(result);
}
