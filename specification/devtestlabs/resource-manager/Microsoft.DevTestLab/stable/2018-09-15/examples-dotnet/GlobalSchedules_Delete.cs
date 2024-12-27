using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Delete.json
// this example is just showing the usage of "GlobalSchedules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabGlobalScheduleResource created on azure
// for more information of creating DevTestLabGlobalScheduleResource, please refer to the document of DevTestLabGlobalScheduleResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string name = "labvmautostart";
ResourceIdentifier devTestLabGlobalScheduleResourceId = DevTestLabGlobalScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DevTestLabGlobalScheduleResource devTestLabGlobalSchedule = client.GetDevTestLabGlobalScheduleResource(devTestLabGlobalScheduleResourceId);

// invoke the operation
await devTestLabGlobalSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
