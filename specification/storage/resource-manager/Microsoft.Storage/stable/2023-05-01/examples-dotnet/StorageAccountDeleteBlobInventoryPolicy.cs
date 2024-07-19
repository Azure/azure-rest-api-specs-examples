using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountDeleteBlobInventoryPolicy.json
// this example is just showing the usage of "BlobInventoryPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlobInventoryPolicyResource created on azure
// for more information of creating BlobInventoryPolicyResource, please refer to the document of BlobInventoryPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
BlobInventoryPolicyName blobInventoryPolicyName = BlobInventoryPolicyName.Default;
ResourceIdentifier blobInventoryPolicyResourceId = BlobInventoryPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, blobInventoryPolicyName);
BlobInventoryPolicyResource blobInventoryPolicy = client.GetBlobInventoryPolicyResource(blobInventoryPolicyResourceId);

// invoke the operation
await blobInventoryPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
