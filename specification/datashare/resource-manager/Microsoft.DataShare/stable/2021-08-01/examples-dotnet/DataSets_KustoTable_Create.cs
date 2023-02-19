using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/DataSets_KustoTable_Create.json
// this example is just showing the usage of "DataSets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataShareResource created on azure
// for more information of creating DataShareResource, please refer to the document of DataShareResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
ResourceIdentifier dataShareResourceId = DataShareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName);
DataShareResource dataShare = client.GetDataShareResource(dataShareResourceId);

// get the collection of this ShareDataSetResource
ShareDataSetCollection collection = dataShare.GetShareDataSets();

// invoke the operation
string dataSetName = "Dataset1";
ShareDataSetData data = new KustoTableDataSet(new ResourceIdentifier("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Kusto/clusters/Cluster1/databases/Database1"), new TableLevelSharingProperties()
{
    ExternalTablesToExclude =
    {
    "test11","test12"
    },
    ExternalTablesToInclude =
    {
    "test9","test10"
    },
    MaterializedViewsToExclude =
    {
    "test7","test8"
    },
    MaterializedViewsToInclude =
    {
    "test5","test6"
    },
    TablesToExclude =
    {
    "test3","test4"
    },
    TablesToInclude =
    {
    "test1","test2"
    },
});
ArmOperation<ShareDataSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataSetName, data);
ShareDataSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ShareDataSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
