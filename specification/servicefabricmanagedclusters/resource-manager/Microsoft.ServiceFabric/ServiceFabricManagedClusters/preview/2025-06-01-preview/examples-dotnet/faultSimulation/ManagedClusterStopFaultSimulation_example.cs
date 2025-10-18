using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: 2025-06-01-preview/faultSimulation/ManagedClusterStopFaultSimulation_example.json
// this example is just showing the usage of "ManagedClusters_StopFaultSimulation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedClusterResource created on azure
// for more information of creating ServiceFabricManagedClusterResource, please refer to the document of ServiceFabricManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
ResourceIdentifier serviceFabricManagedClusterResourceId = ServiceFabricManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ServiceFabricManagedClusterResource serviceFabricManagedCluster = client.GetServiceFabricManagedClusterResource(serviceFabricManagedClusterResourceId);

// invoke the operation
FaultSimulationIdContent content = new FaultSimulationIdContent("1bb61ba9-8a41-4d73-b5f0-7fc93b1edfe3");
ArmOperation<FaultSimulation> lro = await serviceFabricManagedCluster.StopFaultSimulationAsync(WaitUntil.Completed, content);
FaultSimulation result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
