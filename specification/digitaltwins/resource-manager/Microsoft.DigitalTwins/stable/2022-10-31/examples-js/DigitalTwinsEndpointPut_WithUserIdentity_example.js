const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update DigitalTwinsInstance endpoint.
 *
 * @summary Create or update DigitalTwinsInstance endpoint.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsEndpointPut_WithUserIdentity_example.json
 */
async function putADigitalTwinsEndpointResourceWithUserAssignedIdentity() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const endpointName = "myServiceBus";
  const endpointDescription = {
    properties: {
      authenticationType: "IdentityBased",
      endpointType: "ServiceBus",
      endpointUri: "sb://mysb.servicebus.windows.net/",
      entityPath: "mysbtopic",
      identity: {
        type: "UserAssigned",
        userAssignedIdentity:
          "/subscriptions/50016170-c839-41ba-a724-51e9df440b9e/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.digitalTwinsEndpoint.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    endpointName,
    endpointDescription
  );
  console.log(result);
}

putADigitalTwinsEndpointResourceWithUserAssignedIdentity().catch(console.error);
