using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-04-01-preview/RetentionPolicies_CreateOrReplace_MaximumSet_Gen.json
// this example is just showing the usage of "RetentionPolicy_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskRetentionPolicyResource created on azure
// for more information of creating DurableTaskRetentionPolicyResource, please refer to the document of DurableTaskRetentionPolicyResource
string subscriptionId = "194D3C1E-462F-4738-9025-092A628C06EB";
string resourceGroupName = "rgdurabletask";
string schedulerName = "testscheduler";
ResourceIdentifier durableTaskRetentionPolicyResourceId = DurableTaskRetentionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName);
DurableTaskRetentionPolicyResource durableTaskRetentionPolicy = client.GetDurableTaskRetentionPolicyResource(durableTaskRetentionPolicyResourceId);

// invoke the operation
DurableTaskRetentionPolicyData data = new DurableTaskRetentionPolicyData
{
    Properties = new DurableTaskRetentionPolicyProperties
    {
        RetentionPolicies = {new DurableTaskRetentionPolicyDetails(30), new DurableTaskRetentionPolicyDetails(10)
        {
        OrchestrationState = DurableTaskPurgeableOrchestrationState.Failed,
        }},
    },
};
ArmOperation<DurableTaskRetentionPolicyResource> lro = await durableTaskRetentionPolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
DurableTaskRetentionPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DurableTaskRetentionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
