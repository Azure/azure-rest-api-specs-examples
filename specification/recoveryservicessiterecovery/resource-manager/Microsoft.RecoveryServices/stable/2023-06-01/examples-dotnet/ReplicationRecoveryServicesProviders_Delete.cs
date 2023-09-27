using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryServicesProviders_Delete.json
// this example is just showing the usage of "ReplicationRecoveryServicesProviders_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryServicesProviderResource created on azure
// for more information of creating SiteRecoveryServicesProviderResource, please refer to the document of SiteRecoveryServicesProviderResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "cloud1";
string providerName = "241641e6-ee7b-4ee4-8141-821fadda43fa";
ResourceIdentifier siteRecoveryServicesProviderResourceId = SiteRecoveryServicesProviderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, providerName);
SiteRecoveryServicesProviderResource siteRecoveryServicesProvider = client.GetSiteRecoveryServicesProviderResource(siteRecoveryServicesProviderResourceId);

// invoke the operation
await siteRecoveryServicesProvider.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
