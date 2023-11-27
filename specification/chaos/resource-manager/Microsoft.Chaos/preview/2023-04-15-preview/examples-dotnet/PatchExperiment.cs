using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;
using Azure.ResourceManager.Chaos.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/PatchExperiment.json
// this example is just showing the usage of "Experiments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExperimentResource created on azure
// for more information of creating ExperimentResource, please refer to the document of ExperimentResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string experimentName = "exampleExperiment";
ResourceIdentifier experimentResourceId = ExperimentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, experimentName);
ExperimentResource experiment = client.GetExperimentResource(experimentResourceId);

// invoke the operation
ExperimentPatch patch = new ExperimentPatch()
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.ManagedIdentity/userAssignedIdentity/exampleUMI")] = new UserAssignedIdentity(),
        },
    },
};
ExperimentResource result = await experiment.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExperimentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
