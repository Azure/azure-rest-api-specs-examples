using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/ServerlessEndpoint/update.json
// this example is just showing the usage of "ServerlessEndpoints_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningServerlessEndpointResource created on azure
// for more information of creating MachineLearningServerlessEndpointResource, please refer to the document of MachineLearningServerlessEndpointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningServerlessEndpointResourceId = MachineLearningServerlessEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningServerlessEndpointResource machineLearningServerlessEndpoint = client.GetMachineLearningServerlessEndpointResource(machineLearningServerlessEndpointResourceId);

// invoke the operation
MachineLearningServerlessEndpointPatch patch = new MachineLearningServerlessEndpointPatch
{
    Sku = new MachineLearningSkuPatch
    {
        Name = "string",
        Tier = MachineLearningSkuTier.Premium,
        Size = "string",
        Family = "string",
        Capacity = 1,
    },
    Identity = new MachineLearningPartialManagedServiceIdentity
    {
        ManagedServiceIdentityType = "None",
        UserAssignedIdentities =
        {
        ["string"] = BinaryData.FromObjectAsJson(new object())
        },
    },
    Tags = { },
};
ArmOperation<MachineLearningServerlessEndpointResource> lro = await machineLearningServerlessEndpoint.UpdateAsync(WaitUntil.Completed, patch);
MachineLearningServerlessEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningServerlessEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
