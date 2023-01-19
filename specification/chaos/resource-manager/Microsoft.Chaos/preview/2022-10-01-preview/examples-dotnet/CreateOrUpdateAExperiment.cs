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

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/CreateOrUpdateAExperiment.json
// this example is just showing the usage of "Experiments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
ExperimentData data = new ExperimentData(new AzureLocation("eastus2euap"), new Step[]
{
new Step("step1",new Branch[]
{
new Branch("branch1",new Models.Action[]
{
new ContinuousAction("urn:csci:microsoft:virtualMachine:shutdown/1.0",XmlConvert.ToTimeSpan("PT10M"),new KeyValuePair[]
{
new KeyValuePair("abruptShutdown","false")
},"selector1")
})
})
}, new Selector[]
{
new Selector(SelectorType.List,"selector1",new TargetReference[]
{
new TargetReference("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-VirtualMachine")
{
ReferenceType = TargetReferenceType.ChaosTarget,
}
})
})
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
};
ArmOperation<ExperimentResource> lro = await experiment.UpdateAsync(WaitUntil.Completed, data);
ExperimentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExperimentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
