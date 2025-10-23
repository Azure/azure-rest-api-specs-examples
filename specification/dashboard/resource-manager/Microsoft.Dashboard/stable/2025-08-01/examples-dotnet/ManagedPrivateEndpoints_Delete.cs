using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2025-08-01/ManagedPrivateEndpoints_Delete.json
// this example is just showing the usage of "ManagedPrivateEndpointModel_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedPrivateEndpointModelResource created on azure
// for more information of creating ManagedPrivateEndpointModelResource, please refer to the document of ManagedPrivateEndpointModelResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string workspaceName = "myWorkspace";
string managedPrivateEndpointName = "myMPEName";
ResourceIdentifier managedPrivateEndpointModelResourceId = ManagedPrivateEndpointModelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, managedPrivateEndpointName);
ManagedPrivateEndpointModelResource managedPrivateEndpointModel = client.GetManagedPrivateEndpointModelResource(managedPrivateEndpointModelResourceId);

// invoke the operation
await managedPrivateEndpointModel.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
