using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevTestLabs;
using Azure.ResourceManager.DevTestLabs.Models;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Execute.json
// this example is just showing the usage of "ServiceFabricSchedules_Execute" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabServiceFabricScheduleResource created on azure
// for more information of creating DevTestLabServiceFabricScheduleResource, please refer to the document of DevTestLabServiceFabricScheduleResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "@me";
string serviceFabricName = "{serviceFrabicName}";
string name = "{scheduleName}";
ResourceIdentifier devTestLabServiceFabricScheduleResourceId = DevTestLabServiceFabricScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName, serviceFabricName, name);
DevTestLabServiceFabricScheduleResource devTestLabServiceFabricSchedule = client.GetDevTestLabServiceFabricScheduleResource(devTestLabServiceFabricScheduleResourceId);

// invoke the operation
await devTestLabServiceFabricSchedule.ExecuteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
