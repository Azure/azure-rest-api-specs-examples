using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/GetProject.json
// this example is just showing the usage of "Projects_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesAccountResource created on azure
// for more information of creating CognitiveServicesAccountResource, please refer to the document of CognitiveServicesAccountResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myResourceGroup";
string accountName = "myAccount";
ResourceIdentifier cognitiveServicesAccountResourceId = CognitiveServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CognitiveServicesAccountResource cognitiveServicesAccount = client.GetCognitiveServicesAccountResource(cognitiveServicesAccountResourceId);

// get the collection of this CognitiveServicesProjectResource
CognitiveServicesProjectCollection collection = cognitiveServicesAccount.GetCognitiveServicesProjects();

// invoke the operation
string projectName = "myProject";
NullableResponse<CognitiveServicesProjectResource> response = await collection.GetIfExistsAsync(projectName);
CognitiveServicesProjectResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CognitiveServicesProjectData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
