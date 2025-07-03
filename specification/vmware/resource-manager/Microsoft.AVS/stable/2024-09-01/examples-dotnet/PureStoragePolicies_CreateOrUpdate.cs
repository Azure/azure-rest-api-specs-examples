using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/PureStoragePolicies_CreateOrUpdate.json
// this example is just showing the usage of "PureStoragePolicy_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPureStoragePolicyResource created on azure
// for more information of creating AvsPureStoragePolicyResource, please refer to the document of AvsPureStoragePolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string storagePolicyName = "storagePolicy1";
ResourceIdentifier avsPureStoragePolicyResourceId = AvsPureStoragePolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, storagePolicyName);
AvsPureStoragePolicyResource avsPureStoragePolicy = client.GetAvsPureStoragePolicyResource(avsPureStoragePolicyResourceId);

// invoke the operation
AvsPureStoragePolicyData data = new AvsPureStoragePolicyData
{
    Properties = new AvsPureStoragePolicyProperties("storagePolicyDefinition1", "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/PureStorage.Block/storagePools/storagePool1"),
};
ArmOperation<AvsPureStoragePolicyResource> lro = await avsPureStoragePolicy.UpdateAsync(WaitUntil.Completed, data);
AvsPureStoragePolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AvsPureStoragePolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
