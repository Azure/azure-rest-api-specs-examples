using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/PutRaiPolicy.json
// this example is just showing the usage of "RaiPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CognitiveServicesAccountResource created on azure
// for more information of creating CognitiveServicesAccountResource, please refer to the document of CognitiveServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
ResourceIdentifier cognitiveServicesAccountResourceId = CognitiveServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CognitiveServicesAccountResource cognitiveServicesAccount = client.GetCognitiveServicesAccountResource(cognitiveServicesAccountResourceId);

// get the collection of this RaiPolicyResource
RaiPolicyCollection collection = cognitiveServicesAccount.GetRaiPolicies();

// invoke the operation
string raiPolicyName = "raiPolicyName";
RaiPolicyData data = new RaiPolicyData()
{
    Properties = new RaiPolicyProperties()
    {
        Mode = RaiPolicyMode.AsynchronousFilter,
        BasePolicyName = "Microsoft.Default",
        ContentFilters =
        {
        new RaiPolicyContentFilter()
        {
        Name = "hate",
        Enabled = false,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = false,
        Source = RaiPolicyContentSource.Prompt,
        },new RaiPolicyContentFilter()
        {
        Name = "hate",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "sexual",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        },new RaiPolicyContentFilter()
        {
        Name = "sexual",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "selfharm",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        },new RaiPolicyContentFilter()
        {
        Name = "selfharm",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "violence",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        },new RaiPolicyContentFilter()
        {
        Name = "violence",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "jailbreak",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        },new RaiPolicyContentFilter()
        {
        Name = "protected_material_text",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "protected_material_code",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        },new RaiPolicyContentFilter()
        {
        Name = "profanity",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }
        },
    },
};
ArmOperation<RaiPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, raiPolicyName, data);
RaiPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RaiPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
