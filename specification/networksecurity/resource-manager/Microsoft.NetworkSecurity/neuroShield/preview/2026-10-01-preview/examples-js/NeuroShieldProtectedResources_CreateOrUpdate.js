const { NetworkSecurityManagementClient } = require("@azure/arm-networksecurity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a protected resource under a NeuroShield resource.
 *
 * @summary creates or updates a protected resource under a NeuroShield resource.
 * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_CreateOrUpdate.json
 */
async function createOrUpdateProtectedResourceWithAUserAssignedTargetIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkSecurityManagementClient(credential, subscriptionId);
  const result = await client.neuroShieldProtectedResources.createOrUpdate(
    "myResourceGroup",
    "myNeuroShield",
    "vm-01",
    {
      properties: {
        targetResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01",
        targetIdentity: {
          type: "UserAssigned",
          userAssignedIdentities: {
            additionalProperties: {
              "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myIdentityRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/neuronIdentity":
                {
                  principalId: "11111111-1111-1111-1111-111111111111",
                  clientId: "22222222-2222-2222-2222-222222222222",
                },
            },
          },
        },
      },
    },
  );
  console.log(result);
}
