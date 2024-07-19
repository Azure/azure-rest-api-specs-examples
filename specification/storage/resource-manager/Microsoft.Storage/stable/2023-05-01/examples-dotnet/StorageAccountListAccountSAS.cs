using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListAccountSAS.json
// this example is just showing the usage of "StorageAccounts_ListAccountSas" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7985";
string accountName = "sto8588";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation
AccountSasContent content = new AccountSasContent(StorageAccountSasSignedService.B, StorageAccountSasSignedResourceType.S, StorageAccountSasPermission.R, DateTimeOffset.Parse("2017-05-24T11:42:03.1567373Z"))
{
    Protocols = StorageAccountHttpProtocol.HttpsHttp,
    SharedAccessStartOn = DateTimeOffset.Parse("2017-05-24T10:42:03.1567373Z"),
    KeyToSign = "key1",
};
GetAccountSasResult result = await storageAccount.GetAccountSasAsync(content);

Console.WriteLine($"Succeeded: {result}");
