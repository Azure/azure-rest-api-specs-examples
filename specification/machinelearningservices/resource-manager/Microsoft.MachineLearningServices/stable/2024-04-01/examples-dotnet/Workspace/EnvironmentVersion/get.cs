using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/EnvironmentVersion/get.json
// this example is just showing the usage of "EnvironmentVersions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningEnvironmentContainerResource created on azure
// for more information of creating MachineLearningEnvironmentContainerResource, please refer to the document of MachineLearningEnvironmentContainerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
ResourceIdentifier machineLearningEnvironmentContainerResourceId = MachineLearningEnvironmentContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
MachineLearningEnvironmentContainerResource machineLearningEnvironmentContainer = client.GetMachineLearningEnvironmentContainerResource(machineLearningEnvironmentContainerResourceId);

// get the collection of this MachineLearningEnvironmentVersionResource
MachineLearningEnvironmentVersionCollection collection = machineLearningEnvironmentContainer.GetMachineLearningEnvironmentVersions();

// invoke the operation
string version = "string";
NullableResponse<MachineLearningEnvironmentVersionResource> response = await collection.GetIfExistsAsync(version);
MachineLearningEnvironmentVersionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MachineLearningEnvironmentVersionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
