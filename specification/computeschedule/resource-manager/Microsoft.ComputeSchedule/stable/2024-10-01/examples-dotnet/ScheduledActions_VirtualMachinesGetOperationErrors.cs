using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2024-10-01/ScheduledActions_VirtualMachinesGetOperationErrors.json
// this example is just showing the usage of "ScheduledActions_GetVirtualMachineOperationErrors" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "D8E30CC0-2763-4FCC-84A8-3C5659281032";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string locationparameter = "eastus2euap";
GetOperationErrorsContent content = new GetOperationErrorsContent(new string[] { "23480d2f-1dca-4610-afb4-dd25eec1f34r" });
GetOperationErrorsResult result = await subscriptionResource.GetVirtualMachineOperationErrorsAsync(locationparameter, content);

Console.WriteLine($"Succeeded: {result}");
