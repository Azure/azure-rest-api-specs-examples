using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/CreateOrUpdateBigDataPool.json
// this example is just showing the usage of "BigDataPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceResource created on azure
// for more information of creating SynapseWorkspaceResource, please refer to the document of SynapseWorkspaceResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
ResourceIdentifier synapseWorkspaceResourceId = SynapseWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceResource synapseWorkspace = client.GetSynapseWorkspaceResource(synapseWorkspaceResourceId);

// get the collection of this SynapseBigDataPoolInfoResource
SynapseBigDataPoolInfoCollection collection = synapseWorkspace.GetSynapseBigDataPoolInfos();

// invoke the operation
string bigDataPoolName = "ExamplePool";
SynapseBigDataPoolInfoData info = new SynapseBigDataPoolInfoData(new AzureLocation("West US 2"))
{
    AutoScale = new BigDataPoolAutoScaleProperties()
    {
        MinNodeCount = 3,
        IsEnabled = true,
        MaxNodeCount = 50,
    },
    AutoPause = new BigDataPoolAutoPauseProperties()
    {
        DelayInMinutes = 15,
        IsEnabled = true,
    },
    IsAutotuneEnabled = false,
    SparkEventsFolder = "/events",
    NodeCount = 4,
    LibraryRequirements = new BigDataPoolLibraryRequirements()
    {
        Content = "",
        Filename = "requirements.txt",
    },
    SparkVersion = "2.4",
    DefaultSparkLogFolder = "/logs",
    NodeSize = BigDataPoolNodeSize.Medium,
    NodeSizeFamily = BigDataPoolNodeSizeFamily.MemoryOptimized,
    Tags =
    {
    ["key"] = "value",
    },
};
ArmOperation<SynapseBigDataPoolInfoResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, bigDataPoolName, info);
SynapseBigDataPoolInfoResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseBigDataPoolInfoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
