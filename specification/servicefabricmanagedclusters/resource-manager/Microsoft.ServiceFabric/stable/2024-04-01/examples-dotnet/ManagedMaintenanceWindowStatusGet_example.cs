using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ManagedMaintenanceWindowStatusGet_example.json
// this example is just showing the usage of "managedMaintenanceWindowStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedClusterResource created on azure
// for more information of creating ServiceFabricManagedClusterResource, please refer to the document of ServiceFabricManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroup1";
string clusterName = "mycluster1";
ResourceIdentifier serviceFabricManagedClusterResourceId = ServiceFabricManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ServiceFabricManagedClusterResource serviceFabricManagedCluster = client.GetServiceFabricManagedClusterResource(serviceFabricManagedClusterResourceId);

// invoke the operation
ManagedMaintenanceWindowStatus result = await serviceFabricManagedCluster.GetManagedMaintenanceWindowStatusAsync();

Console.WriteLine($"Succeeded: {result}");
