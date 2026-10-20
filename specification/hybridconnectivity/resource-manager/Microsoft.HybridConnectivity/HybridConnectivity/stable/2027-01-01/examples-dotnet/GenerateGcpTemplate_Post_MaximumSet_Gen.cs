using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2027-01-01/GenerateGcpTemplate_Post_MaximumSet_Gen.json
// this example is just showing the usage of "GenerateGcpTemplate_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5ACC4579-DB34-4C2F-8F8C-25061168F342";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
GenerateGcpTemplateContent content = new GenerateGcpTemplateContent("lvrdneuerxztxiixlzswcbm")
{
    SolutionTypes = {new PublicCloudConnectorSolutionTypeSettings("hjyownzpfxwiufmd")
    {
    SolutionSettings = new PublicCloudConnectorSolutionSettings(),
    }},
    GcpCloudProfile = new GcpCloudProfile
    {
        ProjectProperties = new GcpProjectProperties("mjubieitixhpm", "mjubieitixhpmid"),
        OrganizationProperties = new GcpOrganizationProperties("vqlzghfdinlamurmg", "mjubieitixhpm", "mjubieitixhpmid")
        {
            ExcludedProjectNumbers = { "sepdnfxmhcrubtklwllxfbhju" },
            ExcludedFolderIds = { "xxl" },
        },
    },
    GcpTemplateFormat = GcpTemplateFormat.Terraform,
};
GenerateGcpTemplateResult result = await subscriptionResource.PostGenerateGcpTemplateAsync(content);

Console.WriteLine($"Succeeded: {result}");
