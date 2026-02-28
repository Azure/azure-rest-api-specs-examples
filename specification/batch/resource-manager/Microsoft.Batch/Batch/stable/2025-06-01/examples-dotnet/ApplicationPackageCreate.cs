using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/ApplicationPackageCreate.json
// this example is just showing the usage of "ApplicationPackage_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchApplicationResource created on azure
// for more information of creating BatchApplicationResource, please refer to the document of BatchApplicationResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string applicationName = "app1";
ResourceIdentifier batchApplicationResourceId = BatchApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, applicationName);
BatchApplicationResource batchApplication = client.GetBatchApplicationResource(batchApplicationResourceId);

// get the collection of this BatchApplicationPackageResource
BatchApplicationPackageCollection collection = batchApplication.GetBatchApplicationPackages();

// invoke the operation
string versionName = "1";
BatchApplicationPackageData data = new BatchApplicationPackageData();
ArmOperation<BatchApplicationPackageResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, versionName, data);
BatchApplicationPackageResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchApplicationPackageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
