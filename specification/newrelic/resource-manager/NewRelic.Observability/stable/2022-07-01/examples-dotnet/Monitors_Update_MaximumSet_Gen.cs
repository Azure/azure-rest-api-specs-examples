using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NewRelicObservability;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Monitors_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NewRelicMonitorResource created on azure
// for more information of creating NewRelicMonitorResource, please refer to the document of NewRelicMonitorResource
string subscriptionId = "hfmjmpyqgezxkp";
string resourceGroupName = "rgNewRelic";
string monitorName = "cdlymktqw";
ResourceIdentifier newRelicMonitorResourceId = NewRelicMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
NewRelicMonitorResource newRelicMonitorResource = client.GetNewRelicMonitorResource(newRelicMonitorResourceId);

// invoke the operation
NewRelicMonitorResourcePatch patch = new NewRelicMonitorResourcePatch()
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key8903")] = new UserAssignedIdentity(),
        },
    },
    Tags =
    {
    ["key164"] = "jqakdrrmmyzytqu",
    },
    NewRelicAccountProperties = new NewRelicAccountProperties()
    {
        UserId = "vcscxlncofcuduadesd",
        AccountInfo = new NewRelicObservabilityAccountInfo()
        {
            AccountId = "xhqmg",
            IngestionKey = "wltnimmhqt",
            Region = new AzureLocation("ljcf"),
        },
        OrganizationId = "k",
        SingleSignOnProperties = new NewRelicSingleSignOnProperties()
        {
            SingleSignOnState = NewRelicSingleSignOnState.Initial,
            EnterpriseAppId = "kwiwfz",
            SingleSignOnUri = new Uri("kvseueuljsxmfwpqctz"),
            ProvisioningState = NewRelicProvisioningState.Accepted,
        },
    },
    UserInfo = new NewRelicObservabilityUserInfo()
    {
        FirstName = "vdftzcggirefejajwahhwhyibutramdaotvnuf",
        LastName = "bcsztgqovdlmzfkjdrngidwzqsevagexzzilnlc",
        EmailAddress = "%6%@4-g.N1.3F-kI1.Ue-.lJso",
        PhoneNumber = "krf",
        Country = "hslqnwdanrconqyekwbnttaetv",
    },
    PlanData = new NewRelicPlanDetails()
    {
        UsageType = NewRelicObservabilityUsageType.Payg,
        BillingCycle = NewRelicObservabilityBillingCycle.Yearly,
        PlanDetails = "tbbiaga",
        EffectiveOn = DateTimeOffset.Parse("2022-12-05T14:11:37.786Z"),
    },
    OrgCreationSource = NewRelicObservabilityOrgCreationSource.Liftr,
    AccountCreationSource = NewRelicObservabilityAccountCreationSource.Liftr,
};
NewRelicMonitorResource result = await newRelicMonitorResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NewRelicMonitorResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
