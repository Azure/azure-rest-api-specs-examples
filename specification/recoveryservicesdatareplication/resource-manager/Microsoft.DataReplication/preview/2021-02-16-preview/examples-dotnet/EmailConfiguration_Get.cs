using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesDataReplication;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_Get.json
// this example is just showing the usage of "EmailConfiguration_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationEmailConfigurationResource created on azure
// for more information of creating DataReplicationEmailConfigurationResource, please refer to the document of DataReplicationEmailConfigurationResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string vaultName = "4";
string emailConfigurationName = "0";
ResourceIdentifier dataReplicationEmailConfigurationResourceId = DataReplicationEmailConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, emailConfigurationName);
DataReplicationEmailConfigurationResource dataReplicationEmailConfiguration = client.GetDataReplicationEmailConfigurationResource(dataReplicationEmailConfigurationResourceId);

// invoke the operation
DataReplicationEmailConfigurationResource result = await dataReplicationEmailConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataReplicationEmailConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
