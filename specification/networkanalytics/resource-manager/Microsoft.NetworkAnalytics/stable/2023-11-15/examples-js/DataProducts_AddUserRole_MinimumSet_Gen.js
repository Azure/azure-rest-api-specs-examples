const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Assign role to the data product.
 *
 * @summary Assign role to the data product.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_AddUserRole_MinimumSet_Gen.json
 */
async function dataProductsAddUserRoleMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
  const subscriptionId =
    process.env["NETWORKANALYTICS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName =
    process.env["NETWORKANALYTICS_RESOURCE_GROUP"] || "aoiresourceGroupName";
  const dataProductName = "dataproduct01";
  const body = {
    dataTypeScope: ["scope"],
    principalId: "00000000-0000-0000-0000-00000000000",
    principalType: "User",
    role: "Reader",
    roleId: "00000000-0000-0000-0000-00000000000",
    userName: "userName",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataProducts.addUserRole(resourceGroupName, dataProductName, body);
  console.log(result);
}
