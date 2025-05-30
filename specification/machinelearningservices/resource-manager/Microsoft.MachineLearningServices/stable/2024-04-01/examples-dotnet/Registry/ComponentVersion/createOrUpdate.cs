using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/ComponentVersion/createOrUpdate.json
// this example is just showing the usage of "RegistryComponentVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearninRegistryComponentVersionResource created on azure
// for more information of creating MachineLearninRegistryComponentVersionResource, please refer to the document of MachineLearninRegistryComponentVersionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string registryName = "my-aml-registry";
string componentName = "string";
string version = "string";
ResourceIdentifier machineLearninRegistryComponentVersionResourceId = MachineLearninRegistryComponentVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, componentName, version);
MachineLearninRegistryComponentVersionResource machineLearninRegistryComponentVersion = client.GetMachineLearninRegistryComponentVersionResource(machineLearninRegistryComponentVersionResourceId);

// invoke the operation
MachineLearningComponentVersionData data = new MachineLearningComponentVersionData(new MachineLearningComponentVersionProperties
{
    ComponentSpec = BinaryData.FromObjectAsJson(new Dictionary<string, object>
    {
        ["8ced901b-d826-477d-bfef-329da9672513"] = null
    }),
    IsAnonymous = false,
    Description = "string",
    Tags =
    {
    ["string"] = "string"
    },
    Properties =
    {
    ["string"] = "string"
    },
});
ArmOperation<MachineLearninRegistryComponentVersionResource> lro = await machineLearninRegistryComponentVersion.UpdateAsync(WaitUntil.Completed, data);
MachineLearninRegistryComponentVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningComponentVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
