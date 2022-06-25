const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a partner registration with the specified parameters.
 *
 * @summary Deletes a partner registration with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerRegistrations_Delete.json
 */
async function partnerRegistrationsDelete() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const partnerRegistrationName = "examplePartnerRegistrationName1";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.partnerRegistrations.beginDeleteAndWait(
    resourceGroupName,
    partnerRegistrationName
  );
  console.log(result);
}

partnerRegistrationsDelete().catch(console.error);
