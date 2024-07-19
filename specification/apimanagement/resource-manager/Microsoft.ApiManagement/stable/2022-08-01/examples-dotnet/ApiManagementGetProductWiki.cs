using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetProductWiki.json
// this example is just showing the usage of "ProductWiki_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceProductWikiResource created on azure
// for more information of creating ServiceProductWikiResource, please refer to the document of ServiceProductWikiResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "57d1f7558aa04f15146d9d8a";
ResourceIdentifier serviceProductWikiResourceId = ServiceProductWikiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId);
ServiceProductWikiResource serviceProductWiki = client.GetServiceProductWikiResource(serviceProductWikiResourceId);

// invoke the operation
ServiceProductWikiResource result = await serviceProductWiki.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WikiContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
