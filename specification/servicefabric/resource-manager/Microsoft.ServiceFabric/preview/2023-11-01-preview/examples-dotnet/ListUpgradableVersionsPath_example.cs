using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabric.Models;
using Azure.ResourceManager.ServiceFabric;

// Generated from example definition: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/preview/2023-11-01-preview/examples/ListUpgradableVersionsPath_example.json
// this example is just showing the usage of "Clusters_ListUpgradableVersions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricClusterResource created on azure
// for more information of creating ServiceFabricClusterResource, please refer to the document of ServiceFabricClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
ResourceIdentifier serviceFabricClusterResourceId = ServiceFabricClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ServiceFabricClusterResource serviceFabricCluster = client.GetServiceFabricClusterResource(serviceFabricClusterResourceId);

// invoke the operation
UpgradableVersionsDescription versionsDescription = new UpgradableVersionsDescription("7.2.432.9590");
UpgradableVersionPathResult result = await serviceFabricCluster.GetUpgradableVersionsAsync(versionsDescription: versionsDescription);

Console.WriteLine($"Succeeded: {result}");
