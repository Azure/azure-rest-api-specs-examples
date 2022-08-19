const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a devcenter resource
 *
 * @summary Creates or updates a devcenter resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/DevCenters_CreateWithUserIdentity.json
 */
async function devCentersCreateWithUserIdentity() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const body = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/identityGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/testidentity1":
          {},
      },
    },
    location: "centralus",
    tags: { costCode: "12345" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devCenters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    devCenterName,
    body
  );
  console.log(result);
}

devCentersCreateWithUserIdentity().catch(console.error);
