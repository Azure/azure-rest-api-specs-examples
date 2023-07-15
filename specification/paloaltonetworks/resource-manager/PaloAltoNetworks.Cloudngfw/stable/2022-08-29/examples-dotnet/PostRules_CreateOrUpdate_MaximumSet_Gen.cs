using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PostRules_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "PostRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostRulestackRuleResource created on azure
// for more information of creating PostRulestackRuleResource, please refer to the document of PostRulestackRuleResource
string globalRulestackName = "lrs1";
string priority = "1";
ResourceIdentifier postRulestackRuleResourceId = PostRulestackRuleResource.CreateResourceIdentifier(globalRulestackName, priority);
PostRulestackRuleResource postRulestackRule = client.GetPostRulestackRuleResource(postRulestackRuleResourceId);

// invoke the operation
PostRulestackRuleData data = new PostRulestackRuleData("postRule1")
{
    ETag = new ETag("c18e6eef-ba3e-49ee-8a85-2b36c863a9d0"),
    Description = "description of post rule",
    RuleState = RulestackStateType.Disabled,
    Source = new SourceAddressInfo()
    {
        Cidrs =
        {
        "1.0.0.1/10"
        },
        Countries =
        {
        "India"
        },
        Feeds =
        {
        "feed"
        },
        PrefixLists =
        {
        "PL1"
        },
    },
    NegateSource = FirewallBooleanType.True,
    Destination = new DestinationAddressInfo()
    {
        Cidrs =
        {
        "1.0.0.1/10"
        },
        Countries =
        {
        "India"
        },
        Feeds =
        {
        "feed"
        },
        PrefixLists =
        {
        "PL1"
        },
        FqdnLists =
        {
        "FQDN1"
        },
    },
    NegateDestination = FirewallBooleanType.True,
    Applications =
    {
    "app1"
    },
    Category = new EdlMatchCategory(new string[]
{
"https://microsoft.com"
}, new string[]
{
"feed"
}),
    Protocol = "HTTP",
    ProtocolPortList =
    {
    "80"
    },
    InboundInspectionCertificate = "cert1",
    AuditComment = "example comment",
    ActionType = RulestackActionType.Allow,
    EnableLogging = RulestackStateType.Disabled,
    DecryptionRuleType = DecryptionRuleType.SslOutboundInspection,
    Tags =
    {
    new RulestackTagInfo("keyName","value")
    },
};
ArmOperation<PostRulestackRuleResource> lro = await postRulestackRule.UpdateAsync(WaitUntil.Completed, data);
PostRulestackRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PostRulestackRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
