const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

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
