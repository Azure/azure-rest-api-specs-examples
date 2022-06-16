const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a shared private link resource
 *
 * @summary Create or update a shared private link resource
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRSharedPrivateLinkResources_CreateOrUpdate.json
 */
async function signalRSharedPrivateLinkResourcesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const sharedPrivateLinkResourceName = "upstream";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const parameters = {
    groupId: "sites",
    privateLinkResourceId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Web/sites/myWebApp",
    requestMessage: "Please approve",
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRSharedPrivateLinkResources.beginCreateOrUpdateAndWait(
    sharedPrivateLinkResourceName,
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

signalRSharedPrivateLinkResourcesCreateOrUpdate().catch(console.error);
