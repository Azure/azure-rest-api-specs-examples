const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate CommunicationService access key. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
 *
 * @summary Regenerate CommunicationService access key. PrimaryKey and SecondaryKey cannot be regenerated at the same time.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/regenerateKey.json
 */
async function regenerateKey() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = { keyType: "Primary" };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationServices.beginRegenerateKeyAndWait(
    resourceGroupName,
    communicationServiceName,
    parameters
  );
  console.log(result);
}

regenerateKey().catch(console.error);
