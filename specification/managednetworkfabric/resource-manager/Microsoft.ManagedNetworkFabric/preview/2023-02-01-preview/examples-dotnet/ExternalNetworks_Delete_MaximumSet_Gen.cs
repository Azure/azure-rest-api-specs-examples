using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "ExternalNetworks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExternalNetworkResource created on azure
// for more information of creating ExternalNetworkResource, please refer to the document of ExternalNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "example-l3domain";
string externalNetworkName = "example-externalnetwork";
ResourceIdentifier externalNetworkResourceId = ExternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, externalNetworkName);
ExternalNetworkResource externalNetwork = client.GetExternalNetworkResource(externalNetworkResourceId);

// invoke the operation
await externalNetwork.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
