using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/RemoveManagementGroupSubscription.json
// this example is just showing the usage of "ManagementGroupSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupSubscriptionResource created on azure
// for more information of creating ManagementGroupSubscriptionResource, please refer to the document of ManagementGroupSubscriptionResource
string groupId = "Group";
string subscriptionId = "728bcbe4-8d56-4510-86c2-4921b8beefbc";
ResourceIdentifier managementGroupSubscriptionResourceId = ManagementGroupSubscriptionResource.CreateResourceIdentifier(groupId, subscriptionId);
ManagementGroupSubscriptionResource managementGroupSubscription = client.GetManagementGroupSubscriptionResource(managementGroupSubscriptionResourceId);

// invoke the operation
string cacheControl = "no-cache";
await managementGroupSubscription.DeleteAsync(WaitUntil.Completed, cacheControl: cacheControl);

Console.WriteLine("Succeeded");
