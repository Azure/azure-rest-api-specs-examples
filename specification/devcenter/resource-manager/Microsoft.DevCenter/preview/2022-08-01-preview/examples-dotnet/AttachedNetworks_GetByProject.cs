using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_GetByProject.json
// this example is just showing the usage of "AttachedNetworks_GetByProject" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProjectResource created on azure
// for more information of creating ProjectResource, please refer to the document of ProjectResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "{projectName}";
ResourceIdentifier projectResourceId = ProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
ProjectResource project = client.GetProjectResource(projectResourceId);

// get the collection of this ProjectAttachedNetworkConnectionResource
ProjectAttachedNetworkConnectionCollection collection = project.GetProjectAttachedNetworkConnections();

// invoke the operation
string attachedNetworkConnectionName = "network-uswest3";
bool result = await collection.ExistsAsync(attachedNetworkConnectionName);

Console.WriteLine($"Succeeded: {result}");
