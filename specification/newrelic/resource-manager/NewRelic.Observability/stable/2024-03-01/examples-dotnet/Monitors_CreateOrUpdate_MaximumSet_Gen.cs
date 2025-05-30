using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-03-01/examples/Monitors_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Monitors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgNewRelic";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NewRelicMonitorResource
NewRelicMonitorResourceCollection collection = resourceGroupResource.GetNewRelicMonitorResources();

// invoke the operation
string monitorName = "cdlymktqw";
NewRelicMonitorResourceData data = new NewRelicMonitorResourceData(new AzureLocation("k"))
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key8903")] = new UserAssignedIdentity()
        },
    },
    NewRelicAccountProperties = new NewRelicAccountProperties
    {
        UserId = "vcscxlncofcuduadesd",
        AccountInfo = new NewRelicObservabilityAccountInfo
        {
            AccountId = "xhqmg",
            IngestionKey = "wltnimmhqt",
            Region = new AzureLocation("ljcf"),
        },
        OrganizationId = "k",
        SingleSignOnProperties = new NewRelicSingleSignOnProperties
        {
            SingleSignOnState = NewRelicSingleSignOnState.Initial,
            EnterpriseAppId = "kwiwfz",
            SingleSignOnUri = new Uri("kvseueuljsxmfwpqctz"),
            ProvisioningState = NewRelicProvisioningState.Accepted,
        },
    },
    UserInfo = new NewRelicObservabilityUserInfo
    {
        FirstName = "vdftzcggirefejajwahhwhyibutramdaotvnuf",
        LastName = "bcsztgqovdlmzfkjdrngidwzqsevagexzzilnlc",
        EmailAddress = "%6%@4-g.N1.3F-kI1.Ue-.lJso",
        PhoneNumber = "krf",
        Country = "hslqnwdanrconqyekwbnttaetv",
    },
    PlanData = new NewRelicPlanDetails
    {
        UsageType = NewRelicObservabilityUsageType.Payg,
        NewRelicPlanBillingCycle = "Yearly",
        PlanDetails = "tbbiaga",
        EffectiveOn = DateTimeOffset.Parse("2022-12-05T14:11:37.786Z"),
    },
    OrgCreationSource = NewRelicObservabilityOrgCreationSource.Liftr,
    AccountCreationSource = NewRelicObservabilityAccountCreationSource.Liftr,
    SubscriptionState = "Suspended",
    SaaSAzureSubscriptionStatus = "Subscribed",
    Tags =
    {
    ["key6976"] = "oaxfhf"
    },
};
ArmOperation<NewRelicMonitorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, monitorName, data);
NewRelicMonitorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NewRelicMonitorResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
