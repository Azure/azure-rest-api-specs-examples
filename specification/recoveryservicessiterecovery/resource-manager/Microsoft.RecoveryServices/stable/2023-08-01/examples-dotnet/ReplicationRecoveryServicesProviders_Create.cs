using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationRecoveryServicesProviders_Create.json
// this example is just showing the usage of "ReplicationRecoveryServicesProviders_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryFabricResource created on azure
// for more information of creating SiteRecoveryFabricResource, please refer to the document of SiteRecoveryFabricResource
string subscriptionId = "cb53d0c3-bd59-4721-89bc-06916a9147ef";
string resourceGroupName = "resourcegroup1";
string resourceName = "migrationvault";
string fabricName = "vmwarefabric1";
ResourceIdentifier siteRecoveryFabricResourceId = SiteRecoveryFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName);
SiteRecoveryFabricResource siteRecoveryFabric = client.GetSiteRecoveryFabricResource(siteRecoveryFabricResourceId);

// get the collection of this SiteRecoveryServicesProviderResource
SiteRecoveryServicesProviderCollection collection = siteRecoveryFabric.GetSiteRecoveryServicesProviders();

// invoke the operation
string providerName = "vmwareprovider1";
SiteRecoveryServicesProviderCreateOrUpdateContent content = new SiteRecoveryServicesProviderCreateOrUpdateContent(new SiteRecoveryAddRecoveryServicesProviderProperties("vmwareprovider1", new IdentityProviderContent(Guid.Parse("72f988bf-86f1-41af-91ab-2d7cd011db47"), "f66fce08-c0c6-47a1-beeb-0ede5ea94f90", "141360b8-5686-4240-a027-5e24e6affeba", "https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874", "https://login.microsoftonline.com"), new IdentityProviderContent(Guid.Parse("72f988bf-86f1-41af-91ab-2d7cd011db47"), "f66fce08-c0c6-47a1-beeb-0ede5ea94f90", "141360b8-5686-4240-a027-5e24e6affeba", "https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874", "https://login.microsoftonline.com")));
ArmOperation<SiteRecoveryServicesProviderResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, providerName, content);
SiteRecoveryServicesProviderResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryServicesProviderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
