using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/ProjectCapabilityHost/get.json
// this example is just showing the usage of "ProjectCapabilityHosts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesProjectResource created on azure
// for more information of creating CognitiveServicesProjectResource, please refer to the document of CognitiveServicesProjectResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string accountName = "account-1";
string projectName = "project-1";
ResourceIdentifier cognitiveServicesProjectResourceId = CognitiveServicesProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, projectName);
CognitiveServicesProjectResource cognitiveServicesProject = client.GetCognitiveServicesProjectResource(cognitiveServicesProjectResourceId);

// get the collection of this CognitiveServicesProjectCapabilityHostResource
CognitiveServicesProjectCapabilityHostCollection collection = cognitiveServicesProject.GetCognitiveServicesProjectCapabilityHosts();

// invoke the operation
string capabilityHostName = "capabilityHostName";
NullableResponse<CognitiveServicesProjectCapabilityHostResource> response = await collection.GetIfExistsAsync(capabilityHostName);
CognitiveServicesProjectCapabilityHostResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CognitiveServicesCapabilityHostData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
