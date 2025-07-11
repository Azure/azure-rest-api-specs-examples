using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/PutRaiPolicy.json
// this example is just showing the usage of "RaiPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RaiPolicyResource created on azure
// for more information of creating RaiPolicyResource, please refer to the document of RaiPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string raiPolicyName = "raiPolicyName";
ResourceIdentifier raiPolicyResourceId = RaiPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, raiPolicyName);
RaiPolicyResource raiPolicy = client.GetRaiPolicyResource(raiPolicyResourceId);

// invoke the operation
RaiPolicyData data = new RaiPolicyData
{
    Properties = new RaiPolicyProperties
    {
        Mode = RaiPolicyMode.AsynchronousFilter,
        BasePolicyName = "Microsoft.Default",
        ContentFilters = {new RaiPolicyContentFilter
        {
        Name = "Hate",
        Enabled = false,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = false,
        Source = RaiPolicyContentSource.Prompt,
        }, new RaiPolicyContentFilter
        {
        Name = "Hate",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Sexual",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }, new RaiPolicyContentFilter
        {
        Name = "Sexual",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Selfharm",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.High,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }, new RaiPolicyContentFilter
        {
        Name = "Selfharm",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Violence",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }, new RaiPolicyContentFilter
        {
        Name = "Violence",
        Enabled = true,
        SeverityThreshold = RaiPolicyContentLevel.Medium,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Jailbreak",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }, new RaiPolicyContentFilter
        {
        Name = "Protected Material Text",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Protected Material Code",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Completion,
        }, new RaiPolicyContentFilter
        {
        Name = "Profanity",
        Enabled = true,
        Blocking = true,
        Source = RaiPolicyContentSource.Prompt,
        }},
    },
};
ArmOperation<RaiPolicyResource> lro = await raiPolicy.UpdateAsync(WaitUntil.Completed, data);
RaiPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RaiPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
