const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update a replication policy.
 *
 * @summary The operation to update a replication policy.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationPolicies_Update.json
 */
async function updatesThePolicy() {
  const subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
  const resourceName = "vault1";
  const resourceGroupName = "resourceGroupPS1";
  const policyName = "protectionprofile1";
  const input = {
    properties: {
      replicationProviderSettings: { instanceType: "HyperVReplicaAzure" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationPolicies.beginUpdateAndWait(
    resourceName,
    resourceGroupName,
    policyName,
    input
  );
  console.log(result);
}

updatesThePolicy().catch(console.error);
