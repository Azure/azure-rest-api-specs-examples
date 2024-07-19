using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateSubscription.json
// this example is just showing the usage of "Subscription_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementSubscriptionResource created on azure
// for more information of creating ApiManagementSubscriptionResource, please refer to the document of ApiManagementSubscriptionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string sid = "testsub";
ResourceIdentifier apiManagementSubscriptionResourceId = ApiManagementSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, sid);
ApiManagementSubscriptionResource apiManagementSubscription = client.GetApiManagementSubscriptionResource(apiManagementSubscriptionResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementSubscriptionPatch patch = new ApiManagementSubscriptionPatch()
{
    DisplayName = "testsub",
};
ApiManagementSubscriptionResource result = await apiManagementSubscription.UpdateAsync(ifMatch, patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SubscriptionContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
