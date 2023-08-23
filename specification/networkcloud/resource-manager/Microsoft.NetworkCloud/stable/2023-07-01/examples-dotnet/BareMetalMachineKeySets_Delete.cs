using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/BareMetalMachineKeySets_Delete.json
// this example is just showing the usage of "BareMetalMachineKeySets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudBareMetalMachineKeySetResource created on azure
// for more information of creating NetworkCloudBareMetalMachineKeySetResource, please refer to the document of NetworkCloudBareMetalMachineKeySetResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
string bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
ResourceIdentifier networkCloudBareMetalMachineKeySetResourceId = NetworkCloudBareMetalMachineKeySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, bareMetalMachineKeySetName);
NetworkCloudBareMetalMachineKeySetResource networkCloudBareMetalMachineKeySet = client.GetNetworkCloudBareMetalMachineKeySetResource(networkCloudBareMetalMachineKeySetResourceId);

// invoke the operation
await networkCloudBareMetalMachineKeySet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
