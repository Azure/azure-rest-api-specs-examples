using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Chaos.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Chaos;

// Generated from example definition: 2025-01-01/Experiments_Update.json
// this example is just showing the usage of "Experiment_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosExperimentResource created on azure
// for more information of creating ChaosExperimentResource, please refer to the document of ChaosExperimentResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string experimentName = "exampleExperiment";
ResourceIdentifier chaosExperimentResourceId = ChaosExperimentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, experimentName);
ChaosExperimentResource chaosExperiment = client.GetChaosExperimentResource(chaosExperimentResourceId);

// invoke the operation
ChaosExperimentPatch patch = new ChaosExperimentPatch
{
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2"
    },
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.ManagedIdentity/userAssignedIdentity/exampleUMI")] = new UserAssignedIdentity()
        },
    },
};
ArmOperation<ChaosExperimentResource> lro = await chaosExperiment.UpdateAsync(WaitUntil.Completed, patch);
ChaosExperimentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ChaosExperimentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
