using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectionContainerMappings_Purge.json
// this example is just showing the usage of "ReplicationProtectionContainerMappings_Purge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProtectionContainerMappingResource created on azure
// for more information of creating ProtectionContainerMappingResource, please refer to the document of ProtectionContainerMappingResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "cloud1";
string protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
string mappingName = "cloud1protectionprofile1";
ResourceIdentifier protectionContainerMappingResourceId = ProtectionContainerMappingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, mappingName);
ProtectionContainerMappingResource protectionContainerMapping = client.GetProtectionContainerMappingResource(protectionContainerMappingResourceId);

// invoke the operation
await protectionContainerMapping.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
