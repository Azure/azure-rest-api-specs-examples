using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppConfiguration;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresCreateReplica.json
// this example is just showing the usage of "Replicas_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppConfigurationReplicaResource created on azure
// for more information of creating AppConfigurationReplicaResource, please refer to the document of AppConfigurationReplicaResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
string resourceGroupName = "myResourceGroup";
string configStoreName = "contoso";
string replicaName = "myReplicaEus";
ResourceIdentifier appConfigurationReplicaResourceId = AppConfigurationReplicaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configStoreName, replicaName);
AppConfigurationReplicaResource appConfigurationReplica = client.GetAppConfigurationReplicaResource(appConfigurationReplicaResourceId);

// invoke the operation
AppConfigurationReplicaData data = new AppConfigurationReplicaData()
{
    Location = new AzureLocation("eastus"),
};
ArmOperation<AppConfigurationReplicaResource> lro = await appConfigurationReplica.UpdateAsync(WaitUntil.Completed, data);
AppConfigurationReplicaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppConfigurationReplicaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
