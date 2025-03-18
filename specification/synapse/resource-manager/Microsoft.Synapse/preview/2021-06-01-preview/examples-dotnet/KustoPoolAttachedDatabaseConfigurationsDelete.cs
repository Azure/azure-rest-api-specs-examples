using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsDelete.json
// this example is just showing the usage of "KustoPoolAttachedDatabaseConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseAttachedDatabaseConfigurationResource created on azure
// for more information of creating SynapseAttachedDatabaseConfigurationResource, please refer to the document of SynapseAttachedDatabaseConfigurationResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "kustorptest";
string kustoPoolName = "kustoclusterrptest4";
string attachedDatabaseConfigurationName = "attachedDatabaseConfigurations1";
ResourceIdentifier synapseAttachedDatabaseConfigurationResourceId = SynapseAttachedDatabaseConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName, attachedDatabaseConfigurationName);
SynapseAttachedDatabaseConfigurationResource synapseAttachedDatabaseConfiguration = client.GetSynapseAttachedDatabaseConfigurationResource(synapseAttachedDatabaseConfigurationResourceId);

// invoke the operation
await synapseAttachedDatabaseConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
