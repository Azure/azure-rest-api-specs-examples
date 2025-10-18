using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: 2025-06-01-preview/faultSimulation/NodeTypeStopFaultSimulation_example.json
// this example is just showing the usage of "NodeTypes_StopFaultSimulation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedNodeTypeResource created on azure
// for more information of creating ServiceFabricManagedNodeTypeResource, please refer to the document of ServiceFabricManagedNodeTypeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string nodeTypeName = "BE";
ResourceIdentifier serviceFabricManagedNodeTypeResourceId = ServiceFabricManagedNodeTypeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, nodeTypeName);
ServiceFabricManagedNodeTypeResource serviceFabricManagedNodeType = client.GetServiceFabricManagedNodeTypeResource(serviceFabricManagedNodeTypeResourceId);

// invoke the operation
FaultSimulationIdContent content = new FaultSimulationIdContent("1bb61ba9-8a41-4d73-b5f0-7fc93b1edfe3");
ArmOperation<FaultSimulation> lro = await serviceFabricManagedNodeType.StopFaultSimulationAsync(WaitUntil.Completed, content);
FaultSimulation result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
