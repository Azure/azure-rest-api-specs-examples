using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/ApplicationPackageDelete.json
// this example is just showing the usage of "ApplicationPackage_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchApplicationPackageResource created on azure
// for more information of creating BatchApplicationPackageResource, please refer to the document of BatchApplicationPackageResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string applicationName = "app1";
string versionName = "1";
ResourceIdentifier batchApplicationPackageResourceId = BatchApplicationPackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, applicationName, versionName);
BatchApplicationPackageResource batchApplicationPackage = client.GetBatchApplicationPackageResource(batchApplicationPackageResourceId);

// invoke the operation
await batchApplicationPackage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
