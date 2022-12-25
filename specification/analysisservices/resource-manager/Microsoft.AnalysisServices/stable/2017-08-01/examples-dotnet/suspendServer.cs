using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Analysis.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Analysis;

// Generated from example definition: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/suspendServer.json
// this example is just showing the usage of "Servers_Suspend" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this AnalysisServerResource created on azure
// for more information of creating AnalysisServerResource, please refer to the document of AnalysisServerResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
string serverName = "azsdktest";
ResourceIdentifier analysisServerResourceId = AnalysisServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
AnalysisServerResource analysisServer = client.GetAnalysisServerResource(analysisServerResourceId);

// invoke the operation
await analysisServer.SuspendAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
