using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Subscription.Models;
using Azure.ResourceManager.Subscription;

// Generated from example definition: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/createAlias.json
// this example is just showing the usage of "Alias_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionAliasResource created on azure
// for more information of creating SubscriptionAliasResource, please refer to the document of SubscriptionAliasResource
string aliasName = "aliasForNewSub";
ResourceIdentifier subscriptionAliasResourceId = SubscriptionAliasResource.CreateResourceIdentifier(aliasName);
SubscriptionAliasResource subscriptionAlias = client.GetSubscriptionAliasResource(subscriptionAliasResourceId);

// invoke the operation
SubscriptionAliasCreateOrUpdateContent content = new SubscriptionAliasCreateOrUpdateContent
{
    DisplayName = "Test Subscription",
    Workload = SubscriptionWorkload.Production,
    BillingScope = "/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH",
    SubscriptionId = null,
    AdditionalProperties = new SubscriptionAliasAdditionalProperties
    {
        ManagementGroupId = null,
        SubscriptionTenantId = Guid.Parse("66f6e4d6-07dc-4aea-94ea-e12d3026a3c8"),
        SubscriptionOwnerId = "f09b39eb-c496-482c-9ab9-afd799572f4c",
        Tags =
        {
        ["tag1"] = "Messi",
        ["tag2"] = "Ronaldo",
        ["tag3"] = "Lebron"
        },
    },
};
ArmOperation<SubscriptionAliasResource> lro = await subscriptionAlias.UpdateAsync(WaitUntil.Completed, content);
SubscriptionAliasResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SubscriptionAliasData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
