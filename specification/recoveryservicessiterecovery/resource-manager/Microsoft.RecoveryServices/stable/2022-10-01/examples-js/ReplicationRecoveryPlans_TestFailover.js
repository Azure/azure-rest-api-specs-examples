const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to start the test failover of a recovery plan.
 *
 * @summary The operation to start the test failover of a recovery plan.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationRecoveryPlans_TestFailover.json
 */
async function executeTestFailoverOfTheRecoveryPlan() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const recoveryPlanName = "RPtest1";
  const input = {
    properties: {
      failoverDirection: "PrimaryToRecovery",
      networkId:
        "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai",
      networkType: "VmNetworkAsInput",
      providerSpecificDetails: [{ instanceType: "HyperVReplicaAzure" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationRecoveryPlans.beginTestFailoverAndWait(
    resourceName,
    resourceGroupName,
    recoveryPlanName,
    input
  );
  console.log(result);
}

executeTestFailoverOfTheRecoveryPlan().catch(console.error);
