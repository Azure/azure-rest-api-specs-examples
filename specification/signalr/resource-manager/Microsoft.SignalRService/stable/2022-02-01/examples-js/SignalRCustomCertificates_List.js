const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all custom certificates.
 *
 * @summary List all custom certificates.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_List.json
 */
async function signalRCustomCertificatesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.signalRCustomCertificates.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

signalRCustomCertificatesList().catch(console.error);
