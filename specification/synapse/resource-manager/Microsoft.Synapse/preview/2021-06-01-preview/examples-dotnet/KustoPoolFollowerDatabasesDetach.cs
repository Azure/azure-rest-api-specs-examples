using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesDetach.json
// this example is just showing the usage of "KustoPools_DetachFollowerDatabases" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseKustoPoolResource created on azure
// for more information of creating SynapseKustoPoolResource, please refer to the document of SynapseKustoPoolResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "kustorptest";
string kustoPoolName = "kustoclusterrptest4";
ResourceIdentifier synapseKustoPoolResourceId = SynapseKustoPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName);
SynapseKustoPoolResource synapseKustoPool = client.GetSynapseKustoPoolResource(synapseKustoPoolResourceId);

// invoke the operation
SynapseFollowerDatabaseDefinition followerDatabaseToRemove = new SynapseFollowerDatabaseDefinition(new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/leader4"), "myAttachedDatabaseConfiguration");
await synapseKustoPool.DetachFollowerDatabasesAsync(WaitUntil.Completed, followerDatabaseToRemove);

Console.WriteLine("Succeeded");
