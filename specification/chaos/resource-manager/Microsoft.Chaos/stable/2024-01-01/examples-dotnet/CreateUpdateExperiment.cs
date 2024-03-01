using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;
using Azure.ResourceManager.Chaos.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/CreateUpdateExperiment.json
// this example is just showing the usage of "Experiments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ChaosExperimentResource
ChaosExperimentCollection collection = resourceGroupResource.GetChaosExperiments();

// invoke the operation
string experimentName = "exampleExperiment";
ChaosExperimentData data = new ChaosExperimentData(new AzureLocation("eastus2euap"), new ChaosExperimentStep[]
{
new ChaosExperimentStep("step1",new ChaosExperimentBranch[]
{
new ChaosExperimentBranch("branch1",new ChaosExperimentAction[]
{
new ChaosContinuousAction("urn:csci:microsoft:virtualMachine:shutdown/1.0",XmlConvert.ToTimeSpan("PT10M"),new ChaosKeyValuePair[]
{
new ChaosKeyValuePair("abruptShutdown","false")
},"selector1")
})
})
}, new ChaosTargetSelector[]
{
new ChaosTargetListSelector("selector1",new ChaosTargetReference[]
{
new ChaosTargetReference(ChaosTargetReferenceType.ChaosTarget,new ResourceIdentifier("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine"))
})
})
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
};
ArmOperation<ChaosExperimentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, experimentName, data);
ChaosExperimentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ChaosExperimentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
