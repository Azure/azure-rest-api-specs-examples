const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to create an ASR replication protection intent item.
 *
 * @summary The operation to create an ASR replication protection intent item.
 * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionIntents_Create.json
 */
async function createProtectionIntentResource() {
  const subscriptionId =
    process.env["RECOVERYSERVICESSITERECOVERY_SUBSCRIPTION_ID"] ||
    "509099b2-9d2c-4636-b43e-bd5cafb6be69";
  const resourceGroupName =
    process.env["RECOVERYSERVICESSITERECOVERY_RESOURCE_GROUP"] || "resourceGroupPS1";
  const resourceName = "vault1";
  const intentObjectName = "vm1";
  const input = {
    properties: {
      providerSpecificDetails: {
        fabricObjectId:
          "/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne/providers/Microsoft.Compute/virtualMachines/vmPpgAv5",
        instanceType: "A2A",
        primaryLocation: "eastUs2",
        recoveryAvailabilityType: "Single",
        recoveryLocation: "westus2",
        recoveryResourceGroupId:
          "/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne-asr",
        recoverySubscriptionId: "ed5bcdf6-d61e-47bd-8ea9-f2bd379a2640",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectionIntents.create(
    resourceGroupName,
    resourceName,
    intentObjectName,
    input,
  );
  console.log(result);
}
