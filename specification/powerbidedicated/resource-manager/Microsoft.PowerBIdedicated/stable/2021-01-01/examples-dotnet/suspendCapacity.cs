using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PowerBIDedicated.Models;
using Azure.ResourceManager.PowerBIDedicated;

// Generated from example definition: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/suspendCapacity.json
// this example is just showing the usage of "Capacities_Suspend" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedCapacityResource created on azure
// for more information of creating DedicatedCapacityResource, please refer to the document of DedicatedCapacityResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
string dedicatedCapacityName = "azsdktest";
ResourceIdentifier dedicatedCapacityResourceId = DedicatedCapacityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dedicatedCapacityName);
DedicatedCapacityResource dedicatedCapacity = client.GetDedicatedCapacityResource(dedicatedCapacityResourceId);

// invoke the operation
await dedicatedCapacity.SuspendAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
