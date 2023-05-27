using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_Delete.json
// this example is just showing the usage of "DefaultCniNetworks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DefaultCniNetworkResource created on azure
// for more information of creating DefaultCniNetworkResource, please refer to the document of DefaultCniNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string defaultCniNetworkName = "defaultCniNetworkName";
ResourceIdentifier defaultCniNetworkResourceId = DefaultCniNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, defaultCniNetworkName);
DefaultCniNetworkResource defaultCniNetwork = client.GetDefaultCniNetworkResource(defaultCniNetworkResourceId);

// invoke the operation
await defaultCniNetwork.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
