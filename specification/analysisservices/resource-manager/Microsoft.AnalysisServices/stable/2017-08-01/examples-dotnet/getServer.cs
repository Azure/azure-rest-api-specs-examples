using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Analysis.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Analysis;

// Generated from example definition: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/getServer.json
// this example is just showing the usage of "Servers_GetDetails" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AnalysisServerResource
AnalysisServerCollection collection = resourceGroupResource.GetAnalysisServers();

// invoke the operation
string serverName = "azsdktest";
bool result = await collection.ExistsAsync(serverName);

Console.WriteLine($"Succeeded: {result}");
