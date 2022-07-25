const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detaches all followers of a database owned by this Kusto Pool.
 *
 * @summary Detaches all followers of a database owned by this Kusto Pool.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesDetach.json
 */
async function kustoPoolDetachFollowerDatabases() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const resourceGroupName = "kustorptest";
  const followerDatabaseToRemove = {
    attachedDatabaseConfigurationName: "myAttachedDatabaseConfiguration",
    kustoPoolResourceId:
      "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/leader4",
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPools.beginDetachFollowerDatabasesAndWait(
    workspaceName,
    kustoPoolName,
    resourceGroupName,
    followerDatabaseToRemove
  );
  console.log(result);
}

kustoPoolDetachFollowerDatabases().catch(console.error);
