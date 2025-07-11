const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Creates the artifacts for web site, or a deployment slot.
 *
 * @summary Description for Creates the artifacts for web site, or a deployment slot.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/PostDeployWorkflowArtifacts.json
 */
async function deploysWorkflowArtifacts() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "testsite2";
  const workflowArtifacts = {
    appSettings: {
      eventHub_connectionString:
        "Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=EXAMPLE1a2b3c4d5e6fEXAMPLE=",
    },
    files: {
      connectionsJson: {
        managedApiConnections: {},
        serviceProviderConnections: {
          eventHub: {
            displayName: "example1",
            parameterValues: {
              connectionString: "@appsetting('eventHub_connectionString')",
            },
            serviceProvider: { id: "/serviceProviders/eventHub" },
          },
        },
      },
      "test1/workflowJson": {
        definition: {
          $schema:
            "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
          actions: {},
          contentVersion: "1.0.0.0",
          outputs: {},
          triggers: {
            When_events_are_available_in_Event_hub: {
              type: "ServiceProvider",
              inputs: {
                parameters: { eventHubName: "test123" },
                serviceProviderConfiguration: {
                  operationId: "receiveEvents",
                  connectionName: "eventHub",
                  serviceProviderId: "/serviceProviders/eventHub",
                },
              },
              splitOn: "@triggerOutputs()?['body']",
            },
          },
        },
        kind: "Stateful",
      },
    },
  };
  const options = {
    workflowArtifacts,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.deployWorkflowArtifacts(resourceGroupName, name, options);
  console.log(result);
}
