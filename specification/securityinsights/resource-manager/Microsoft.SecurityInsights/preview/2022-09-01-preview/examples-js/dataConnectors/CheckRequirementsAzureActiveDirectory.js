const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get requirements state for a data connector type.
 *
 * @summary Get requirements state for a data connector type.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CheckRequirementsAzureActiveDirectory.json
 */
async function checkRequirementsForAad() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const dataConnectorsCheckRequirements = {
    kind: "AzureActiveDirectory",
    tenantId: "2070ecc9-b4d5-4ae4-adaa-936fa1954fa8",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectorsCheckRequirementsOperations.post(
    resourceGroupName,
    workspaceName,
    dataConnectorsCheckRequirements
  );
  console.log(result);
}
