using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountSetManagementPolicy.json
// this example is just showing the usage of "ManagementPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountManagementPolicyResource created on azure
// for more information of creating StorageAccountManagementPolicyResource, please refer to the document of StorageAccountManagementPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7687";
string accountName = "sto9699";
ManagementPolicyName managementPolicyName = ManagementPolicyName.Default;
ResourceIdentifier storageAccountManagementPolicyResourceId = StorageAccountManagementPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, managementPolicyName);
StorageAccountManagementPolicyResource storageAccountManagementPolicy = client.GetStorageAccountManagementPolicyResource(storageAccountManagementPolicyResourceId);

// invoke the operation
StorageAccountManagementPolicyData data = new StorageAccountManagementPolicyData()
{
    Rules =
    {
    new ManagementPolicyRule("olcmtest1",ManagementPolicyRuleType.Lifecycle,new ManagementPolicyDefinition(new ManagementPolicyAction()
    {
    BaseBlob = new ManagementPolicyBaseBlob()
    {
    TierToCool = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 30,
    },
    TierToArchive = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 90,
    },
    Delete = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 1000,
    },
    },
    Snapshot = new ManagementPolicySnapShot()
    {
    Delete = new DateAfterCreation(30),
    },
    })
    {
    Filters = new ManagementPolicyFilter(new string[]
    {
    "blockBlob"
    })
    {
    PrefixMatch =
    {
    "olcmtestcontainer1"
    },
    },
    })
    {
    IsEnabled = true,
    },new ManagementPolicyRule("olcmtest2",ManagementPolicyRuleType.Lifecycle,new ManagementPolicyDefinition(new ManagementPolicyAction()
    {
    BaseBlob = new ManagementPolicyBaseBlob()
    {
    TierToCool = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 30,
    },
    TierToArchive = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 90,
    },
    Delete = new DateAfterModification()
    {
    DaysAfterModificationGreaterThan = 1000,
    },
    },
    })
    {
    Filters = new ManagementPolicyFilter(new string[]
    {
    "blockBlob"
    })
    {
    PrefixMatch =
    {
    "olcmtestcontainer2"
    },
    BlobIndexMatch =
    {
    new ManagementPolicyTagFilter("tag1","==","val1"),new ManagementPolicyTagFilter("tag2","==","val2")
    },
    },
    })
    {
    IsEnabled = true,
    }
    },
};
ArmOperation<StorageAccountManagementPolicyResource> lro = await storageAccountManagementPolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
StorageAccountManagementPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountManagementPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
