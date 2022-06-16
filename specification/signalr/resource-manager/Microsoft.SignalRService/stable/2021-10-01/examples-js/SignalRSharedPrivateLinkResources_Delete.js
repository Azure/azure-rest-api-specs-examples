const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRSharedPrivateLinkResourcesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const sharedPrivateLinkResourceName = "upstream";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRSharedPrivateLinkResources.beginDeleteAndWait(
    sharedPrivateLinkResourceName,
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

signalRSharedPrivateLinkResourcesDelete().catch(console.error);
