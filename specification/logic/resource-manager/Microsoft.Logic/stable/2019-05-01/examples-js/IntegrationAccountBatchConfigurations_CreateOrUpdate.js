const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a batch configuration for an integration account.
 *
 * @summary Create or update a batch configuration for an integration account.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_CreateOrUpdate.json
 */
async function createOrUpdateABatchConfiguration() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const integrationAccountName = "testIntegrationAccount";
  const batchConfigurationName = "testBatchConfiguration";
  const batchConfiguration = {
    location: "westus",
    properties: {
      batchGroupName: "DEFAULT",
      releaseCriteria: {
        batchSize: 234567,
        messageCount: 10,
        recurrence: {
          frequency: "Minute",
          interval: 1,
          startTime: "2017-03-24T11:43:00",
          timeZone: "India Standard Time",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.integrationAccountBatchConfigurations.createOrUpdate(
    resourceGroupName,
    integrationAccountName,
    batchConfigurationName,
    batchConfiguration
  );
  console.log(result);
}
