using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceGraph.Models;
using Azure.ResourceManager.ResourceGraph;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2024-04-01/examples/GraphQueryDelete.json
// this example is just showing the usage of "GraphQuery_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGraphQueryResource created on azure
// for more information of creating ResourceGraphQueryResource, please refer to the document of ResourceGraphQueryResource
string subscriptionId = "024e2271-06fa-46b6-9079-f1ed3c7b070e";
string resourceGroupName = "my-resource-group";
string resourceName = "MyDockerVM";
ResourceIdentifier resourceGraphQueryResourceId = ResourceGraphQueryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ResourceGraphQueryResource resourceGraphQuery = client.GetResourceGraphQueryResource(resourceGraphQueryResourceId);

// invoke the operation
await resourceGraphQuery.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
