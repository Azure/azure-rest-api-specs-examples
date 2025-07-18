using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2024-11-01-preview/ManagedPrivateEndpoints_Patch.json
// this example is just showing the usage of "ManagedPrivateEndpointModel_Update" operation, for the dependent resources, they will have to be created separately.

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
ManagedPrivateEndpointModelPatch patch = new ManagedPrivateEndpointModelPatch
{
    Tags =
    {
    ["Environment"] = "Dev 2"
    },
};
ArmOperation<ManagedPrivateEndpointModelResource> lro = await managedPrivateEndpointModel.UpdateAsync(WaitUntil.Completed, patch);
ManagedPrivateEndpointModelResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedPrivateEndpointModelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
