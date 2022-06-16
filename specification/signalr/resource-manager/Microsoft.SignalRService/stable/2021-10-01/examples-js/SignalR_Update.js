const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const parameters = {
    cors: { allowedOrigins: ["https://foo.com", "https://bar.com"] },
    disableAadAuth: false,
    disableLocalAuth: false,
    features: [
      { flag: "ServiceMode", properties: {}, value: "Serverless" },
      { flag: "EnableConnectivityLogs", properties: {}, value: "True" },
      { flag: "EnableMessagingLogs", properties: {}, value: "False" },
      { flag: "EnableLiveTrace", properties: {}, value: "False" },
    ],
    identity: { type: "SystemAssigned" },
    kind: "SignalR",
    location: "eastus",
    networkACLs: {
      defaultAction: "Deny",
      privateEndpoints: [
        {
          name: "mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e",
          allow: ["ServerConnection"],
        },
      ],
      publicNetwork: { allow: ["ClientConnection"] },
    },
    publicNetworkAccess: "Enabled",
    sku: { name: "Standard_S1", capacity: 1, tier: "Standard" },
    tags: { key1: "value1" },
    tls: { clientCertEnabled: false },
    upstream: {
      templates: [
        {
          auth: {
            type: "ManagedIdentity",
            managedIdentity: { resource: "api://example" },
          },
          categoryPattern: "*",
          eventPattern: "connect,disconnect",
          hubPattern: "*",
          urlTemplate: "https://example.com/chat/api/connect",
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalR.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

signalRUpdate().catch(console.error);
