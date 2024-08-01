using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Containers.Models;
using Azure.ResourceManager.HDInsight.Containers;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/CreateRangerCluster.json
// this example is just showing the usage of "Clusters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightClusterPoolResource created on azure
// for more information of creating HDInsightClusterPoolResource, please refer to the document of HDInsightClusterPoolResource
string subscriptionId = "10e32bab-26da-4cc4-a441-52b318f824e6";
string resourceGroupName = "hiloResourcegroup";
string clusterPoolName = "clusterpool1";
ResourceIdentifier hdInsightClusterPoolResourceId = HDInsightClusterPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterPoolName);
HDInsightClusterPoolResource hdInsightClusterPool = client.GetHDInsightClusterPoolResource(hdInsightClusterPoolResourceId);

// get the collection of this HDInsightClusterResource
HDInsightClusterCollection collection = hdInsightClusterPool.GetHDInsightClusters();

// invoke the operation
string clusterName = "cluster1";
HDInsightClusterData data = new HDInsightClusterData(new AzureLocation("West US 2"))
{
    Properties = new HDInsightClusterProperties("ranger", new ClusterComputeProfile(new ClusterComputeNodeProfile[]
{
new ClusterComputeNodeProfile("head","Standard_D3_v2",2)
})
    {
        AvailabilityZones =
        {
        "1","2","3"
        },
    }, new ClusterProfile("0.0.1", "2.2.3", new AuthorizationProfile()
    {
        UserIds =
        {
        "testuser1","testuser2"
        },
    })
    {
        IdentityList =
        {
        new HDInsightManagedIdentitySpec(HDInsightManagedIdentityType.Cluster,new ResourceIdentifier("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),"de91f1d8-767f-460a-ac11-3cf103f74b34","40491351-c240-4042-91e0-f644a1d2b441")
        },
        RangerProfile = new RangerProfile(new RangerAdminSpec(new string[]
{
"testuser1@contoso.com","testuser2@contoso.com"
}, new RangerAdminSpecDatabase("testsqlserver.database.windows.net", "testdb")
{
    PasswordSecretRef = "https://testkv.vault.azure.net/secrets/mysecret/5df6584d9c25418c8d900240aa6c3452",
    Username = "admin",
}), new RangerUsersyncSpec()
{
    IsEnabled = true,
    Groups =
{
"0a53828f-36c9-44c3-be3d-99a7fce977ad","13be6971-79db-4f33-9d41-b25589ca25ac"
},
    Mode = RangerUsersyncMode.Automatic,
    Users =
{
"testuser1@contoso.com","testuser2@contoso.com"
},
})
        {
            RangerAuditStorageAccount = "https://teststorage.blob.core.windows.net/testblob",
        },
    }),
};
ArmOperation<HDInsightClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
HDInsightClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HDInsightClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
