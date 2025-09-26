using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DisconnectedOperations.Models;
using Azure.ResourceManager.DisconnectedOperations;

// Generated from example definition: 2025-06-01-preview/DisconnectedOperations_ListDeploymentManifest_MaximumSet_Gen.json
// this example is just showing the usage of "DisconnectedOperations_ListDeploymentManifest" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DisconnectedOperationResource created on azure
// for more information of creating DisconnectedOperationResource, please refer to the document of DisconnectedOperationResource
string subscriptionId = "301DCB09-82EC-4777-A56C-6FFF26BCC814";
string resourceGroupName = "rgdisconnectedoperations";
string name = "demo-resource";
ResourceIdentifier disconnectedOperationResourceId = DisconnectedOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DisconnectedOperationResource disconnectedOperation = client.GetDisconnectedOperationResource(disconnectedOperationResourceId);

// invoke the operation
DisconnectedOperationDeploymentManifest result = await disconnectedOperation.GetDeploymentManifestAsync();

Console.WriteLine($"Succeeded: {result}");
