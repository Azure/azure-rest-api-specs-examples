using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Kusto.Models;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoAttachedDatabaseConfigurationsGet.json
// this example is just showing the usage of "AttachedDatabaseConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoAttachedDatabaseConfigurationResource created on azure
// for more information of creating KustoAttachedDatabaseConfigurationResource, please refer to the document of KustoAttachedDatabaseConfigurationResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster2";
string attachedDatabaseConfigurationName = "attachedDatabaseConfigurationsTest";
ResourceIdentifier kustoAttachedDatabaseConfigurationResourceId = KustoAttachedDatabaseConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, attachedDatabaseConfigurationName);
KustoAttachedDatabaseConfigurationResource kustoAttachedDatabaseConfiguration = client.GetKustoAttachedDatabaseConfigurationResource(kustoAttachedDatabaseConfigurationResourceId);

// invoke the operation
KustoAttachedDatabaseConfigurationResource result = await kustoAttachedDatabaseConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KustoAttachedDatabaseConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
