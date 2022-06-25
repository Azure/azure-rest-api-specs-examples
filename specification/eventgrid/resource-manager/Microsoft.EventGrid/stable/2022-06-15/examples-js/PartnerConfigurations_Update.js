const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously updates a partner configuration with the specified parameters.
 *
 * @summary Synchronously updates a partner configuration with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerConfigurations_Update.json
 */
async function partnerConfigurationsUpdate() {
  const subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
  const resourceGroupName = "examplerg";
  const partnerConfigurationUpdateParameters = {
    defaultMaximumExpirationTimeInDays: 100,
    tags: { tag1: "value11", tag2: "value22" },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.partnerConfigurations.beginUpdateAndWait(
    resourceGroupName,
    partnerConfigurationUpdateParameters
  );
  console.log(result);
}

partnerConfigurationsUpdate().catch(console.error);
