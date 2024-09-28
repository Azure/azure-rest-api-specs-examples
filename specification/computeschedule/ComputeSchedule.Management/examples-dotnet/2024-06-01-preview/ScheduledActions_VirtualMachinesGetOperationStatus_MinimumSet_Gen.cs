using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2024-08-15-preview/ScheduledActions_VirtualMachinesGetOperationStatus_MinimumSet_Gen.json
// this example is just showing the usage of "ScheduledActions_GetVirtualMachineOperationStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "DE84A209-5715-43E7-BC76-3E208A9A82C5";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string locationparameter = "csay";
GetOperationStatusContent content = new GetOperationStatusContent(new string[]
{
"01080d2f-1dca-4610-afb4-dd25eec1f3df"
}, "01080d2f-1dca-4610-afb4-dd25eec1f3c1");
GetOperationStatusResult result = await subscriptionResource.GetVirtualMachineOperationStatusAsync(locationparameter, content);

Console.WriteLine($"Succeeded: {result}");
