using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/Instances_Delete.json
// this example is just showing the usage of "Instances_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdateInstanceResource created on azure
// for more information of creating DeviceUpdateInstanceResource, please refer to the document of DeviceUpdateInstanceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
string instanceName = "blue";
ResourceIdentifier deviceUpdateInstanceResourceId = DeviceUpdateInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, instanceName);
DeviceUpdateInstanceResource deviceUpdateInstance = client.GetDeviceUpdateInstanceResource(deviceUpdateInstanceResourceId);

// invoke the operation
await deviceUpdateInstance.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
