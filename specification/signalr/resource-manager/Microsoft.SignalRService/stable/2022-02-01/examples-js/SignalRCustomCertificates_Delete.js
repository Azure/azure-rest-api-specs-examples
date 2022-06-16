const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a custom certificate.
 *
 * @summary Delete a custom certificate.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_Delete.json
 */
async function signalRCustomCertificatesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const certificateName = "myCert";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRCustomCertificates.delete(
    resourceGroupName,
    resourceName,
    certificateName
  );
  console.log(result);
}

signalRCustomCertificatesDelete().catch(console.error);
