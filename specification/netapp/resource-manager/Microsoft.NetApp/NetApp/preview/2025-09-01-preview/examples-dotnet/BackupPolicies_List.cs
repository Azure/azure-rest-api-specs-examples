using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/BackupPolicies_List.json
// this example is just showing the usage of "BackupPolicies_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppAccountResource created on azure
// for more information of creating NetAppAccountResource, please refer to the document of NetAppAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
ResourceIdentifier netAppAccountResourceId = NetAppAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
NetAppAccountResource netAppAccount = client.GetNetAppAccountResource(netAppAccountResourceId);

// get the collection of this NetAppBackupPolicyResource
NetAppBackupPolicyCollection collection = netAppAccount.GetNetAppBackupPolicies();

// invoke the operation and iterate over the result
await foreach (NetAppBackupPolicyResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppBackupPolicyData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
