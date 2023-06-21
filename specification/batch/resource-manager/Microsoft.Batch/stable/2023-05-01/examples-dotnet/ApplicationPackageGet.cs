using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/ApplicationPackageGet.json
// this example is just showing the usage of "ApplicationPackage_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchApplicationResource created on azure
// for more information of creating BatchApplicationResource, please refer to the document of BatchApplicationResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string applicationName = "app1";
ResourceIdentifier batchApplicationResourceId = BatchApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, applicationName);
BatchApplicationResource batchApplication = client.GetBatchApplicationResource(batchApplicationResourceId);

// get the collection of this BatchApplicationPackageResource
BatchApplicationPackageCollection collection = batchApplication.GetBatchApplicationPackages();

// invoke the operation
string versionName = "1";
bool result = await collection.ExistsAsync(versionName);

Console.WriteLine($"Succeeded: {result}");
