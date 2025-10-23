using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/StorageAccountGetManagementPolicy.json
// this example is just showing the usage of "ManagementPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountManagementPolicyResource created on azure
// for more information of creating StorageAccountManagementPolicyResource, please refer to the document of StorageAccountManagementPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
ManagementPolicyName managementPolicyName = ManagementPolicyName.Default;
ResourceIdentifier storageAccountManagementPolicyResourceId = StorageAccountManagementPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, managementPolicyName);
StorageAccountManagementPolicyResource storageAccountManagementPolicy = client.GetStorageAccountManagementPolicyResource(storageAccountManagementPolicyResourceId);

// invoke the operation
StorageAccountManagementPolicyResource result = await storageAccountManagementPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountManagementPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
