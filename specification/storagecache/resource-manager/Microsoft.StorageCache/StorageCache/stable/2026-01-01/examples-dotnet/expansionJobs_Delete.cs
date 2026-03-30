using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageCache.Models;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/StorageCache/stable/2026-01-01/examples/expansionJobs_Delete.json
// this example is just showing the usage of "expansionJobs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpansionJobResource created on azure
// for more information of creating ExpansionJobResource, please refer to the document of ExpansionJobResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string amlFileSystemName = "fs1";
string expansionJobName = "expansionjob1";
ResourceIdentifier expansionJobResourceId = ExpansionJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, amlFileSystemName, expansionJobName);
ExpansionJobResource expansionJob = client.GetExpansionJobResource(expansionJobResourceId);

// invoke the operation
await expansionJob.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
