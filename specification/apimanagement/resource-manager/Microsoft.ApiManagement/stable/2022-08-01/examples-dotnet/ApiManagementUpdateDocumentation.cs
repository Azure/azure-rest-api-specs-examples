using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateDocumentation.json
// this example is just showing the usage of "Documentation_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DocumentationContractResource created on azure
// for more information of creating DocumentationContractResource, please refer to the document of DocumentationContractResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string documentationId = "57d1f7558aa04f15146d9d8a";
ResourceIdentifier documentationContractResourceId = DocumentationContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, documentationId);
DocumentationContractResource documentationContract = client.GetDocumentationContractResource(documentationContractResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
DocumentationContractPatch patch = new DocumentationContractPatch()
{
    Title = "Title updated",
    Content = "content updated",
};
DocumentationContractResource result = await documentationContract.UpdateAsync(ifMatch, patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DocumentationContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
