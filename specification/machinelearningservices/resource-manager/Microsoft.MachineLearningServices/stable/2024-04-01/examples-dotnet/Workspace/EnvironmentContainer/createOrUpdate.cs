using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/EnvironmentContainer/createOrUpdate.json
// this example is just showing the usage of "EnvironmentContainers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg123";
string workspaceName = "testworkspace";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningEnvironmentContainerResource
MachineLearningEnvironmentContainerCollection collection = machineLearningWorkspace.GetMachineLearningEnvironmentContainers();

// invoke the operation
string name = "testEnvironment";
MachineLearningEnvironmentContainerData data = new MachineLearningEnvironmentContainerData(new MachineLearningEnvironmentContainerProperties
{
    Description = "string",
    Tags =
    {
    ["additionalProp1"] = "string",
    ["additionalProp2"] = "string",
    ["additionalProp3"] = "string"
    },
    Properties =
    {
    ["additionalProp1"] = "string",
    ["additionalProp2"] = "string",
    ["additionalProp3"] = "string"
    },
});
ArmOperation<MachineLearningEnvironmentContainerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
MachineLearningEnvironmentContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningEnvironmentContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
