using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotasSubscriptions/DeleteGroupQuotaSubscriptions.json
// this example is just showing the usage of "GroupQuotaSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GroupQuotaSubscriptionResource created on azure
// for more information of creating GroupQuotaSubscriptionResource, please refer to the document of GroupQuotaSubscriptionResource
string managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
string groupQuotaName = "groupquota1";
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier groupQuotaSubscriptionResourceId = GroupQuotaSubscriptionResource.CreateResourceIdentifier(managementGroupId, groupQuotaName, subscriptionId);
GroupQuotaSubscriptionResource groupQuotaSubscription = client.GetGroupQuotaSubscriptionResource(groupQuotaSubscriptionResourceId);

// invoke the operation
await groupQuotaSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
