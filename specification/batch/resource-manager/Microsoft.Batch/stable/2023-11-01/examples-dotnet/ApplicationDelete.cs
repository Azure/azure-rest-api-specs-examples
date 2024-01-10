using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/ApplicationDelete.json
// this example is just showing the usage of "Application_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await batchApplication.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
