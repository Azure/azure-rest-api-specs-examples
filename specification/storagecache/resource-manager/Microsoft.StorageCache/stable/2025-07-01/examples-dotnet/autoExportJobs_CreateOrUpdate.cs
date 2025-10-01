using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageCache;

// Generated from example definition: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/autoExportJobs_CreateOrUpdate.json
// this example is just showing the usage of "autoExportJobs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AmlFileSystemResource created on azure
// for more information of creating AmlFileSystemResource, please refer to the document of AmlFileSystemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "scgroup";
string amlFileSystemName = "fs1";
ResourceIdentifier amlFileSystemResourceId = AmlFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, amlFileSystemName);
AmlFileSystemResource amlFileSystem = client.GetAmlFileSystemResource(amlFileSystemResourceId);

// get the collection of this AutoExportJobResource
AutoExportJobCollection collection = amlFileSystem.GetAutoExportJobs();

// invoke the operation
string autoExportJobName = "job1";
AutoExportJobData data = new AutoExportJobData(new AzureLocation("eastus"))
{
    AutoExportPrefixes = { "/" },
    Tags =
    {
    ["Dept"] = "ContosoAds"
    },
};
ArmOperation<AutoExportJobResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, autoExportJobName, data);
AutoExportJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutoExportJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
