using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetUserSubscription.json
// this example is just showing the usage of "UserSubscription_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementUserResource created on azure
// for more information of creating ApiManagementUserResource, please refer to the document of ApiManagementUserResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string userId = "1";
ResourceIdentifier apiManagementUserResourceId = ApiManagementUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, userId);
ApiManagementUserResource apiManagementUser = client.GetApiManagementUserResource(apiManagementUserResourceId);

// get the collection of this ApiManagementUserSubscriptionResource
ApiManagementUserSubscriptionCollection collection = apiManagementUser.GetApiManagementUserSubscriptions();

// invoke the operation
string sid = "5fa9b096f3df14003c070001";
NullableResponse<ApiManagementUserSubscriptionResource> response = await collection.GetIfExistsAsync(sid);
ApiManagementUserSubscriptionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubscriptionContractData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
