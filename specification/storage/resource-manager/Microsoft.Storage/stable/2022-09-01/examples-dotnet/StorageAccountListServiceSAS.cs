using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountListServiceSAS.json
// this example is just showing the usage of "StorageAccounts_ListServiceSAS" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7439";
string accountName = "sto1299";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation
ServiceSasContent content = new ServiceSasContent("/blob/sto1299/music")
{
    Resource = ServiceSasSignedResourceType.Container,
    Permissions = StorageAccountSasPermission.L,
    SharedAccessExpiryOn = DateTimeOffset.Parse("2017-05-24T11:32:48.8457197Z"),
};
GetServiceSasResult result = await storageAccount.GetServiceSasAsync(content);

Console.WriteLine($"Succeeded: {result}");
