using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/CodeContainer/createOrUpdate.json
// this example is just showing the usage of "RegistryCodeContainers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningRegistryCodeContainerResource created on azure
// for more information of creating MachineLearningRegistryCodeContainerResource, please refer to the document of MachineLearningRegistryCodeContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string registryName = "testregistry";
string codeName = "testContainer";
ResourceIdentifier machineLearningRegistryCodeContainerResourceId = MachineLearningRegistryCodeContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, codeName);
MachineLearningRegistryCodeContainerResource machineLearningRegistryCodeContainer = client.GetMachineLearningRegistryCodeContainerResource(machineLearningRegistryCodeContainerResourceId);

// invoke the operation
MachineLearningCodeContainerData data = new MachineLearningCodeContainerData(new MachineLearningCodeContainerProperties
{
    Description = "string",
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
});
ArmOperation<MachineLearningRegistryCodeContainerResource> lro = await machineLearningRegistryCodeContainer.UpdateAsync(WaitUntil.Completed, data);
MachineLearningRegistryCodeContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningCodeContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
