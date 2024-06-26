const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update GuestAgent.
 *
 * @summary Create Or Update GuestAgent.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateGuestAgent.json
 */
async function createGuestAgent() {
  const subscriptionId =
    process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "testrg";
  const virtualMachineName = "ContosoVm";
  const name = "default";
  const body = {
    credentials: { password: "<password>", username: "tempuser" },
    httpProxyConfig: { httpsProxy: "http://192.1.2.3:8080" },
    provisioningAction: "install",
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.guestAgents.beginCreateAndWait(
    resourceGroupName,
    virtualMachineName,
    name,
    options
  );
  console.log(result);
}
