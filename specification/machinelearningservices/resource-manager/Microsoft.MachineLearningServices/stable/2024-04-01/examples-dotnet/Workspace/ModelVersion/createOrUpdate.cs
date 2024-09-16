using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/ModelVersion/createOrUpdate.json
// this example is just showing the usage of "ModelVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningModelVersionResource created on azure
// for more information of creating MachineLearningModelVersionResource, please refer to the document of MachineLearningModelVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
string name = "string";
string version = "string";
ResourceIdentifier machineLearningModelVersionResourceId = MachineLearningModelVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name, version);
MachineLearningModelVersionResource machineLearningModelVersion = client.GetMachineLearningModelVersionResource(machineLearningModelVersionResourceId);

// invoke the operation
MachineLearningModelVersionData data = new MachineLearningModelVersionData(new MachineLearningModelVersionProperties()
{
    Flavors =
    {
    ["string"] = new MachineLearningFlavorData()
    {
    Data =
    {
    ["string"] = "string",
    },
    },
    },
    ModelType = "CustomModel",
    ModelUri = new Uri("string"),
    IsAnonymous = false,
    Description = "string",
    Tags =
    {
    ["string"] = "string",
    },
    Properties =
    {
    ["string"] = "string",
    },
});
ArmOperation<MachineLearningModelVersionResource> lro = await machineLearningModelVersion.UpdateAsync(WaitUntil.Completed, data);
MachineLearningModelVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningModelVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
