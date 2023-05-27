using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Delete.json
// this example is just showing the usage of "BareMetalMachineKeySets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BareMetalMachineKeySetResource created on azure
// for more information of creating BareMetalMachineKeySetResource, please refer to the document of BareMetalMachineKeySetResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
string bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
ResourceIdentifier bareMetalMachineKeySetResourceId = BareMetalMachineKeySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, bareMetalMachineKeySetName);
BareMetalMachineKeySetResource bareMetalMachineKeySet = client.GetBareMetalMachineKeySetResource(bareMetalMachineKeySetResourceId);

// invoke the operation
await bareMetalMachineKeySet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
