using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiCenter;
using Azure.ResourceManager.ApiCenter.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Services_Get.json
// this example is just showing the usage of "Services_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterServiceResource created on azure
// for more information of creating ApiCenterServiceResource, please refer to the document of ApiCenterServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
ResourceIdentifier apiCenterServiceResourceId = ApiCenterServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiCenterServiceResource apiCenterService = client.GetApiCenterServiceResource(apiCenterServiceResourceId);

// invoke the operation
ApiCenterServiceResource result = await apiCenterService.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiCenterServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
