using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/SolutionConfigurations_Get.json
// this example is just showing the usage of "SolutionConfiguration_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this PublicCloudConnectorSolutionConfigurationResource
string resourceUri = "ymuj";
PublicCloudConnectorSolutionConfigurationCollection collection = client.GetPublicCloudConnectorSolutionConfigurations(new ResourceIdentifier(resourceUri));

// invoke the operation
string solutionConfiguration = "tks";
NullableResponse<PublicCloudConnectorSolutionConfigurationResource> response = await collection.GetIfExistsAsync(solutionConfiguration);
PublicCloudConnectorSolutionConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PublicCloudConnectorSolutionConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
