using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/PostDeployWorkflowArtifacts.json
// this example is just showing the usage of "WebApps_DeployWorkflowArtifacts" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "testsite2";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// invoke the operation
WorkflowArtifacts workflowArtifacts = new WorkflowArtifacts
{
    AppSettings = BinaryData.FromObjectAsJson(new
    {
        eventHub_connectionString = "Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=EXAMPLE1a2b3c4d5e6fEXAMPLE=",
    }),
    Files =
    {
    ["connections.json"] = BinaryData.FromObjectAsJson(new
    {
    managedApiConnections = new object(),
    serviceProviderConnections = new
    {
    eventHub = new
    {
    displayName = "example1",
    parameterValues = new
    {
    connectionString = "@appsetting('eventHub_connectionString')",
    },
    serviceProvider = new
    {
    id = "/serviceProviders/eventHub",
    },
    },
    },
    }),
    ["test1/workflow.json"] = BinaryData.FromObjectAsJson(new
    {
    definition = new Dictionary<string, object>
    {
    ["$schema"] = "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
    ["actions"] = new object(),
    ["contentVersion"] = "1.0.0.0",
    ["outputs"] = new object(),
    ["triggers"] = new
    {
    When_events_are_available_in_Event_hub = new
    {
    type = "ServiceProvider",
    inputs = new
    {
    parameters = new
    {
    eventHubName = "test123",
    },
    serviceProviderConfiguration = new
    {
    operationId = "receiveEvents",
    connectionName = "eventHub",
    serviceProviderId = "/serviceProviders/eventHub",
    },
    },
    splitOn = "@triggerOutputs()?['body']",
    },
    }
    },
    kind = "Stateful",
    })
    },
};
await webSite.DeployWorkflowArtifactsAsync(workflowArtifacts: workflowArtifacts);

Console.WriteLine("Succeeded");
