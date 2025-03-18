using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PowerBIDedicated.Models;
using Azure.ResourceManager.PowerBIDedicated;

// Generated from example definition: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/deleteAutoScaleVCore.json
// this example is just showing the usage of "AutoScaleVCores_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutoScaleVCoreResource created on azure
// for more information of creating AutoScaleVCoreResource, please refer to the document of AutoScaleVCoreResource
string subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
string resourceGroupName = "TestRG";
string vcoreName = "testvcore";
ResourceIdentifier autoScaleVCoreResourceId = AutoScaleVCoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vcoreName);
AutoScaleVCoreResource autoScaleVCore = client.GetAutoScaleVCoreResource(autoScaleVCoreResourceId);

// invoke the operation
await autoScaleVCore.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
