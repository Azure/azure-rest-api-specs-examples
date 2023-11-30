const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate sas token for storage account.
 *
 * @summary Generate sas token for storage account.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_GenerateStorageAccountSasToken_MaximumSet_Gen.json
 */
async function dataProductsGenerateStorageAccountSasTokenMaximumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const body = {
    expiryTimeStamp: new Date("2023-08-24T05:34:58.151Z"),
    ipAddress: "1.1.1.1",
    startTimeStamp: new Date("2023-08-24T05:34:58.151Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataProducts.generateStorageAccountSasToken(
    resourceGroupName,
    dataProductName,
    body
  );
  console.log(result);
}
