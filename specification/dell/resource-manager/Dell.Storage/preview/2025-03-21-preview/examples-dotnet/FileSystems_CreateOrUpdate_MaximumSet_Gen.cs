using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Dell.Storage.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Dell.Storage;

// Generated from example definition: 2025-03-21-preview/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "FileSystemResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "4B6E265D-57CF-4A9D-8B35-3CC68ED9D208";
string resourceGroupName = "rgDell";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DellFileSystemResource
DellFileSystemCollection collection = resourceGroupResource.GetDellFileSystems();

// invoke the operation
string filesystemName = "abcd";
DellFileSystemData data = new DellFileSystemData(new AzureLocation("cvbmsqftppe"))
{
    Properties = new DellFileSystemProperties(
    new DellFileSystemMarketplaceDetails("eekvwfndjoxijeasksnt", "bcganbkmvznyqfnvhjuag", "planeName")
    {
        MarketplaceSubscriptionId = "mvjcxwndudbylynme",
        PublisherId = "trdzykoeskmcwpo",
        PrivateOfferId = "privateOfferId",
    },
    new ResourceIdentifier("rqkpvczbtqcxiaivtbuixblb"),
    "10.0.0.1/24",
    new DellFileSystemUserDetails("jwogfgznmjabdbcjcljjlkxdpc"),
    "fhewkj",
    new DellFileSystemEncryptionProperties(DellFileSystemEncryptionType.CustomerManagedKeysCmk)
    {
        KeyUri = "https://contoso.com/keyurl/keyVersion",
        EncryptionIdentityProperties = new DellFileSystemEncryptionIdentityProperties
        {
            IdentityType = DellFileSystemEncryptionIdentityType.UserAssigned,
            IdentityResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
        },
    })
    {
        SmartConnectFqdn = "fqdn",
        OneFsUri = new Uri("oneFsUrl"),
    },
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key7644")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key7594"] = "sfkwapubiurgedzveido"
    },
};
ArmOperation<DellFileSystemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, filesystemName, data);
DellFileSystemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DellFileSystemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
