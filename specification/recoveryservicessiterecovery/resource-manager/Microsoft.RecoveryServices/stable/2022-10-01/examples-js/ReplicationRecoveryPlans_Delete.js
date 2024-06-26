const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a recovery plan.
 *
 * @summary Delete a recovery plan.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationRecoveryPlans_Delete.json
 */
async function deletesTheSpecifiedRecoveryPlan() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const recoveryPlanName = "RPtest1";
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationRecoveryPlans.beginDeleteAndWait(
    resourceName,
    resourceGroupName,
    recoveryPlanName
  );
  console.log(result);
}

deletesTheSpecifiedRecoveryPlan().catch(console.error);
