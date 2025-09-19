using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/TransitHub_Delete.json
// this example is just showing the usage of "TransitHubResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveTransitHubResource created on azure
// for more information of creating VirtualEnclaveTransitHubResource, please refer to the document of VirtualEnclaveTransitHubResource
string subscriptionId = "CA1CB369-DD26-4DB2-9D43-9AFEF0F22093";
string resourceGroupName = "rgopenapi";
string communityName = "TestMyCommunity";
string transitHubName = "TestThName";
ResourceIdentifier virtualEnclaveTransitHubResourceId = VirtualEnclaveTransitHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communityName, transitHubName);
VirtualEnclaveTransitHubResource virtualEnclaveTransitHub = client.GetVirtualEnclaveTransitHubResource(virtualEnclaveTransitHubResourceId);

// invoke the operation
await virtualEnclaveTransitHub.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
