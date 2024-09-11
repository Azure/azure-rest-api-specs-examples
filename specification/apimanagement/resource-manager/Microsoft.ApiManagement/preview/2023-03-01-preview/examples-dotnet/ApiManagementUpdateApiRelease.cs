using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementUpdateApiRelease.json
// this example is just showing the usage of "ApiRelease_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiReleaseResource created on azure
// for more information of creating ApiReleaseResource, please refer to the document of ApiReleaseResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "a1";
string releaseId = "testrev";
ResourceIdentifier apiReleaseResourceId = ApiReleaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, releaseId);
ApiReleaseResource apiRelease = client.GetApiReleaseResource(apiReleaseResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiReleaseData data = new ApiReleaseData()
{
    ApiId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/a1"),
    Notes = "yahooagain",
};
ApiReleaseResource result = await apiRelease.UpdateAsync(ifMatch, data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiReleaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
