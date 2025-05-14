using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationJobs_Get.json
// this example is just showing the usage of "ReplicationJobs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SiteRecoveryJobResource
string resourceName = "vault1";
SiteRecoveryJobCollection collection = resourceGroupResource.GetSiteRecoveryJobs(resourceName);

// invoke the operation
string jobName = "58776d0b-3141-48b2-a377-9ad863eb160d";
NullableResponse<SiteRecoveryJobResource> response = await collection.GetIfExistsAsync(jobName);
SiteRecoveryJobResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SiteRecoveryJobData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
