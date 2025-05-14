using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.StandbyPool.Models;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: 2025-03-01/StandbyContainerGroupPools_Delete.json
// this example is just showing the usage of "StandbyContainerGroupPoolResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StandbyContainerGroupPoolResource created on azure
// for more information of creating StandbyContainerGroupPoolResource, please refer to the document of StandbyContainerGroupPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000009";
string resourceGroupName = "rgstandbypool";
string standbyContainerGroupPoolName = "pool";
ResourceIdentifier standbyContainerGroupPoolResourceId = StandbyContainerGroupPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, standbyContainerGroupPoolName);
StandbyContainerGroupPoolResource standbyContainerGroupPool = client.GetStandbyContainerGroupPoolResource(standbyContainerGroupPoolResourceId);

// invoke the operation
await standbyContainerGroupPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
