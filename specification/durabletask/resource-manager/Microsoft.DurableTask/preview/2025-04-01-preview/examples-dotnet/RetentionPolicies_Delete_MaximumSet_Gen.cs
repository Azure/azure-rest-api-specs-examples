using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-04-01-preview/RetentionPolicies_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "RetentionPolicy_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskRetentionPolicyResource created on azure
// for more information of creating DurableTaskRetentionPolicyResource, please refer to the document of DurableTaskRetentionPolicyResource
string subscriptionId = "194D3C1E-462F-4738-9025-092A628C06EB";
string resourceGroupName = "rgdurabletask";
string schedulerName = "testcheduler";
ResourceIdentifier durableTaskRetentionPolicyResourceId = DurableTaskRetentionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName);
DurableTaskRetentionPolicyResource durableTaskRetentionPolicy = client.GetDurableTaskRetentionPolicyResource(durableTaskRetentionPolicyResourceId);

// invoke the operation
await durableTaskRetentionPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
