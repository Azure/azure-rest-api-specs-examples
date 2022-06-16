const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new CommunicationService or update an existing CommunicationService.
 *
 * @summary Create a new CommunicationService or update an existing CommunicationService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/createOrUpdate.json
 */
async function createOrUpdateResource() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = {
    dataLocation: "United States",
    location: "Global",
  };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationService.beginCreateOrUpdateAndWait(
    resourceGroupName,
    communicationServiceName,
    options
  );
  console.log(result);
}

createOrUpdateResource().catch(console.error);
