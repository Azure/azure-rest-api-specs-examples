using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_Delete.json
// this example is just showing the usage of "NotificationChannels_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabNotificationChannelResource created on azure
// for more information of creating DevTestLabNotificationChannelResource, please refer to the document of DevTestLabNotificationChannelResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string name = "{notificationChannelName}";
ResourceIdentifier devTestLabNotificationChannelResourceId = DevTestLabNotificationChannelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabNotificationChannelResource devTestLabNotificationChannel = client.GetDevTestLabNotificationChannelResource(devTestLabNotificationChannelResourceId);

// invoke the operation
await devTestLabNotificationChannel.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
