const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a new partner registration with the specified parameters.
 *
 * @summary Creates a new partner registration with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerRegistrations_CreateOrUpdate.json
 */
async function partnerRegistrationsCreateOrUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const partnerRegistrationName = "examplePartnerRegistrationName1";
  const partnerRegistrationInfo = {
    location: "global",
    tags: { key1: "value1", key2: "Value2", key3: "Value3" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.partnerRegistrations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    partnerRegistrationName,
    partnerRegistrationInfo,
  );
  console.log(result);
}
