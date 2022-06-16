const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a custom domain.
 *
 * @summary Create or update a custom domain.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_CreateOrUpdate.json
 */
async function signalRCustomDomainsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const name = "myDomain";
  const parameters = {
    customCertificate: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/customCertificates/myCert",
    },
    domainName: "example.com",
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRCustomDomains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    name,
    parameters
  );
  console.log(result);
}

signalRCustomDomainsCreateOrUpdate().catch(console.error);
