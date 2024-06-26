const { MicrosoftNetworkAnalytics } = require("@azure/arm-networkanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove role from the data product.
 *
 * @summary Remove role from the data product.
 * x-ms-original-file: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_RemoveUserRole_MinimumSet_Gen.json
 */
async function dataProductsRemoveUserRoleMaximumSetGenGeneratedByMinimumSetRuleMinimumSetGen() {
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
    roleAssignmentId: "00000000-0000-0000-0000-00000000000",
    roleId: "00000000-0000-0000-0000-00000000000",
    userName: "UserName",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftNetworkAnalytics(credential, subscriptionId);
  const result = await client.dataProducts.removeUserRole(resourceGroupName, dataProductName, body);
  console.log(result);
}
