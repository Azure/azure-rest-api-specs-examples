const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the artifacts for the logic app
 *
 * @summary Creates or updates the artifacts for the logic app
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/LogicApps_PostDeployWorkflowArtifacts.json
 */
async function deploysWorkflowArtifacts() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "testrg123";
  const containerAppName = "testapp2";
  const logicAppName = "testapp2";
  const workflowArtifacts = {
    appSettings: {
      eventHub_connectionString:
        "Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=EXAMPLE1a2b3c4d5e6fEXAMPLE=",
    },
    files: {
      "connections.json": {
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
      "test1/workflow.json": {
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
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.logicApps.deployWorkflowArtifacts(
    resourceGroupName,
    containerAppName,
    logicAppName,
    options,
  );
  console.log(result);
}
