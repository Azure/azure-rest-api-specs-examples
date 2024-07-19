using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateApi.json
// this example is just showing the usage of "Api_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiResource created on azure
// for more information of creating ApiResource, please refer to the document of ApiResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "echo-api";
ResourceIdentifier apiResourceId = ApiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId);
ApiResource api = client.GetApiResource(apiResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiPatch patch = new ApiPatch()
{
    DisplayName = "Echo API New",
    ServiceLink = "http://echoapi.cloudapp.net/api2",
    Path = "newecho",
};
ApiResource result = await api.UpdateAsync(ifMatch, patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
