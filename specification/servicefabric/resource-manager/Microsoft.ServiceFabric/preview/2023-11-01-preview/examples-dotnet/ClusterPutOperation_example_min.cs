using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceFabric;
using Azure.ResourceManager.ServiceFabric.Models;

// Generated from example definition: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview/examples/ClusterPutOperation_example_min.json
// this example is just showing the usage of "Clusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ServiceFabricClusterResource
ServiceFabricClusterCollection collection = resourceGroupResource.GetServiceFabricClusters();

// invoke the operation
string clusterName = "myCluster";
ServiceFabricClusterData data = new ServiceFabricClusterData(new AzureLocation("eastus"))
{
    DiagnosticsStorageAccountConfig = new DiagnosticsStorageAccountConfig("diag", "StorageAccountKey1", new Uri("https://diag.blob.core.windows.net/"), new Uri("https://diag.queue.core.windows.net/"), new Uri("https://diag.table.core.windows.net/")),
    FabricSettings =
    {
    new SettingsSectionDescription("UpgradeService",new SettingsParameterDescription[]
    {
    new SettingsParameterDescription("AppPollIntervalInSeconds","60")
    })
    },
    ManagementEndpoint = new Uri("http://myCluster.eastus.cloudapp.azure.com:19080"),
    NodeTypes =
    {
    new ClusterNodeTypeDescription("nt1vm",19000,19007,true,5)
    {
    DurabilityLevel = ClusterDurabilityLevel.Bronze,
    ApplicationPorts = new ClusterEndpointRangeDescription(20000,30000),
    EphemeralPorts = new ClusterEndpointRangeDescription(49000,64000),
    }
    },
    ReliabilityLevel = ClusterReliabilityLevel.Silver,
    UpgradeMode = ClusterUpgradeMode.Automatic,
    Tags =
    {
    },
};
ArmOperation<ServiceFabricClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
ServiceFabricClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
