using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/NetworkSecurityPerimeterConfigurationsReconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Reconcile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SearchServiceNetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating SearchServiceNetworkSecurityPerimeterConfigurationResource, please refer to the document of SearchServiceNetworkSecurityPerimeterConfigurationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string searchServiceName = "mysearchservice";
string nspConfigName = "00000001-2222-3333-4444-111144444444.assoc1";
ResourceIdentifier searchServiceNetworkSecurityPerimeterConfigurationResourceId = SearchServiceNetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, searchServiceName, nspConfigName);
SearchServiceNetworkSecurityPerimeterConfigurationResource searchServiceNetworkSecurityPerimeterConfiguration = client.GetSearchServiceNetworkSecurityPerimeterConfigurationResource(searchServiceNetworkSecurityPerimeterConfigurationResourceId);

// invoke the operation
await searchServiceNetworkSecurityPerimeterConfiguration.ReconcileAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
