using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/Workspace/OnlineEndpoint/update.json
// this example is just showing the usage of "OnlineEndpoints_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningOnlineEndpointResource created on azure
// for more information of creating MachineLearningOnlineEndpointResource, please refer to the document of MachineLearningOnlineEndpointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string endpointName = "testEndpointName";
ResourceIdentifier machineLearningOnlineEndpointResourceId = MachineLearningOnlineEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, endpointName);
MachineLearningOnlineEndpointResource machineLearningOnlineEndpoint = client.GetMachineLearningOnlineEndpointResource(machineLearningOnlineEndpointResourceId);

// invoke the operation
MachineLearningResourcePatchWithIdentity body = new MachineLearningResourcePatchWithIdentity()
{
    Identity = new MachineLearningPartialManagedServiceIdentity()
    {
        ManagedServiceIdentityType = "SystemAssigned",
        UserAssignedIdentities =
        {
        ["string"] = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
        }),
        },
    },
    Tags =
    {
    },
};
ArmOperation<MachineLearningOnlineEndpointResource> lro = await machineLearningOnlineEndpoint.UpdateAsync(WaitUntil.Completed, body);
MachineLearningOnlineEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningOnlineEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
