using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateApiWiki.json
// this example is just showing the usage of "ApiWiki_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceApiWikiResource created on azure
// for more information of creating ServiceApiWikiResource, please refer to the document of ServiceApiWikiResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d1f7558aa04f15146d9d8a";
ResourceIdentifier serviceApiWikiResourceId = ServiceApiWikiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId);
ServiceApiWikiResource serviceApiWiki = client.GetServiceApiWikiResource(serviceApiWikiResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
WikiUpdateContract wikiUpdateContract = new WikiUpdateContract
{
    Documents = {new WikiDocumentationContract
    {
    DocumentationId = "docId1",
    }},
};
ServiceApiWikiResource result = await serviceApiWiki.UpdateAsync(ifMatch, wikiUpdateContract);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WikiContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
