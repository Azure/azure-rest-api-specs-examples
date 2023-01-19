using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountGetEncryptionScope.json
// this example is just showing the usage of "EncryptionScopes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EncryptionScopeResource created on azure
// for more information of creating EncryptionScopeResource, please refer to the document of EncryptionScopeResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "resource-group-name";
string accountName = "{storage-account-name}";
string encryptionScopeName = "{encryption-scope-name}";
ResourceIdentifier encryptionScopeResourceId = EncryptionScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, encryptionScopeName);
EncryptionScopeResource encryptionScope = client.GetEncryptionScopeResource(encryptionScopeResourceId);

// invoke the operation
EncryptionScopeResource result = await encryptionScope.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EncryptionScopeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
