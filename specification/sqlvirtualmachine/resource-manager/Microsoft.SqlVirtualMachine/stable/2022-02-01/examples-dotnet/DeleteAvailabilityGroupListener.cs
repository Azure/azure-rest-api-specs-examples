using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SqlVirtualMachine.Models;
using Azure.ResourceManager.SqlVirtualMachine;

// Generated from example definition: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/stable/2022-02-01/examples/DeleteAvailabilityGroupListener.json
// this example is just showing the usage of "AvailabilityGroupListeners_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvailabilityGroupListenerResource created on azure
// for more information of creating AvailabilityGroupListenerResource, please refer to the document of AvailabilityGroupListenerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVmGroupName = "testvmgroup";
string availabilityGroupListenerName = "agl-test";
ResourceIdentifier availabilityGroupListenerResourceId = AvailabilityGroupListenerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVmGroupName, availabilityGroupListenerName);
AvailabilityGroupListenerResource availabilityGroupListener = client.GetAvailabilityGroupListenerResource(availabilityGroupListenerResourceId);

// invoke the operation
await availabilityGroupListener.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
