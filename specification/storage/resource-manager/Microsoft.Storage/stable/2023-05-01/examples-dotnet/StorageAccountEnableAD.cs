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

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountEnableAD.json
// this example is just showing the usage of "StorageAccounts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9407";
string accountName = "sto8596";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation
StorageAccountPatch patch = new StorageAccountPatch()
{
    AzureFilesIdentityBasedAuthentication = new FilesIdentityBasedAuthentication(DirectoryServiceOption.AD)
    {
        ActiveDirectoryProperties = new StorageActiveDirectoryProperties("adtest.com", Guid.Parse("aebfc118-9fa9-4732-a21f-d98e41a77ae1"))
        {
            NetBiosDomainName = "adtest.com",
            ForestName = "adtest.com",
            DomainSid = "S-1-5-21-2400535526-2334094090-2402026252",
            AzureStorageSid = "S-1-5-21-2400535526-2334094090-2402026252-0012",
            SamAccountName = "sam12498",
            AccountType = ActiveDirectoryAccountType.User,
        },
    },
};
StorageAccountResource result = await storageAccount.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
