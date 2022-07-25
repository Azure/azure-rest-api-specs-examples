const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get requirements state for a data connector type.
 *
 * @summary Get requirements state for a data connector type.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/dataConnectors/CheckRequirementsIoT.json
 */
async function checkRequirementsForIoT() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const dataConnectorsCheckRequirements = {
    kind: "IOT",
    subscriptionId: "c0688291-89d7-4bed-87a2-a7b1bff43f4c",
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

checkRequirementsForIoT().catch(console.error);
