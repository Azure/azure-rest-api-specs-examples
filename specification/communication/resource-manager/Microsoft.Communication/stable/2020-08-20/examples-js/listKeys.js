const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the access keys of the CommunicationService resource.
 *
 * @summary Get the access keys of the CommunicationService resource.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listKeys.json
 */
async function listKeys() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationService.listKeys(
    resourceGroupName,
    communicationServiceName
  );
  console.log(result);
}

listKeys().catch(console.error);
