using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/AddManagementGroupSubscription.json
// this example is just showing the usage of "ManagementGroupSubscriptions_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string groupId = "Group";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(groupId);
ManagementGroupResource managementGroup = client.GetManagementGroupResource(managementGroupResourceId);

// get the collection of this ManagementGroupSubscriptionResource
ManagementGroupSubscriptionCollection collection = managementGroup.GetManagementGroupSubscriptions();

// invoke the operation
string subscriptionId = "728bcbe4-8d56-4510-86c2-4921b8beefbc";
string cacheControl = "no-cache";
ArmOperation<ManagementGroupSubscriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, subscriptionId, cacheControl: cacheControl);
ManagementGroupSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagementGroupSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
