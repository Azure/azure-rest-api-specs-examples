using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountCreateDnsEndpointTypeToStandard.json
// this example is just showing the usage of "StorageAccounts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9101";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StorageAccountResource
StorageAccountCollection collection = resourceGroupResource.GetStorageAccounts();

// invoke the operation
string accountName = "sto4445";
StorageAccountCreateOrUpdateContent content = new StorageAccountCreateOrUpdateContent(new StorageSku(StorageSkuName.StandardGrs), StorageKind.Storage, new AzureLocation("eastus"))
{
    ExtendedLocation = new ExtendedLocation()
    {
        Name = "losangeles001",
    },
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2",
    },
    SasPolicy = new StorageAccountSasPolicy("1.15:59:59", ExpirationAction.Log),
    KeyExpirationPeriodInDays = 20,
    Encryption = new StorageAccountEncryption()
    {
        Services = new StorageAccountEncryptionServices()
        {
            Blob = new StorageEncryptionService()
            {
                IsEnabled = true,
                KeyType = StorageEncryptionKeyType.Account,
            },
            File = new StorageEncryptionService()
            {
                IsEnabled = true,
                KeyType = StorageEncryptionKeyType.Account,
            },
        },
        KeySource = StorageAccountKeySource.Storage,
        RequireInfrastructureEncryption = false,
    },
    IsSftpEnabled = true,
    IsHnsEnabled = true,
    RoutingPreference = new StorageRoutingPreference()
    {
        RoutingChoice = StorageRoutingChoice.MicrosoftRouting,
        IsMicrosoftEndpointsPublished = true,
        IsInternetEndpointsPublished = true,
    },
    AllowBlobPublicAccess = false,
    MinimumTlsVersion = StorageMinimumTlsVersion.Tls1_2,
    AllowSharedKeyAccess = true,
    IsDefaultToOAuthAuthentication = false,
    DnsEndpointType = StorageDnsEndpointType.Standard,
};
ArmOperation<StorageAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, content);
StorageAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
