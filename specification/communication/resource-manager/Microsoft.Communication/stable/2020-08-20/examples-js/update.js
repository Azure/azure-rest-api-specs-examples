const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an existing CommunicationService.
 *
 * @summary Operation to update an existing CommunicationService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/update.json
 */
async function updateResource() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = {
    tags: { newTag: "newVal" },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationService.update(
    resourceGroupName,
    communicationServiceName,
    options
  );
  console.log(result);
}

updateResource().catch(console.error);
