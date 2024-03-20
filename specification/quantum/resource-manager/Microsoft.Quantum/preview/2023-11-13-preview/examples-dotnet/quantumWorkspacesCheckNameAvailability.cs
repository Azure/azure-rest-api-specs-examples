using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quantum;
using Azure.ResourceManager.Quantum.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/quantumWorkspacesCheckNameAvailability.json
// this example is just showing the usage of "Workspace_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string locationName = "westus2";
WorkspaceNameAvailabilityContent content = new WorkspaceNameAvailabilityContent()
{
    Name = "sample-workspace-name",
    ResourceType = new ResourceType("Microsoft.Quantum/Workspaces"),
};
WorkspaceNameAvailabilityResult result = await subscriptionResource.CheckWorkspaceNameAvailabilityAsync(locationName, content);

Console.WriteLine($"Succeeded: {result}");
