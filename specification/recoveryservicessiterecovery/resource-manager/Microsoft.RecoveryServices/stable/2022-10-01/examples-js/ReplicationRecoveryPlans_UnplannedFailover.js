const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to start the unplanned failover of a recovery plan.
 *
 * @summary The operation to start the unplanned failover of a recovery plan.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationRecoveryPlans_UnplannedFailover.json
 */
async function executeUnplannedFailoverOfTheRecoveryPlan() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const recoveryPlanName = "RPtest1";
  const input = {
    properties: {
      failoverDirection: "PrimaryToRecovery",
      providerSpecificDetails: [{ instanceType: "HyperVReplicaAzure" }],
      sourceSiteOperations: "Required",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationRecoveryPlans.beginUnplannedFailoverAndWait(
    resourceName,
    resourceGroupName,
    recoveryPlanName,
    input
  );
  console.log(result);
}

executeUnplannedFailoverOfTheRecoveryPlan().catch(console.error);
