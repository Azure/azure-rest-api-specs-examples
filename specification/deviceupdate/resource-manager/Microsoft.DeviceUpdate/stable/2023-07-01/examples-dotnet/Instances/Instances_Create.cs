using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceUpdate.Models;
using Azure.ResourceManager.DeviceUpdate;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/Instances_Create.json
// this example is just showing the usage of "Instances_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdateAccountResource created on azure
// for more information of creating DeviceUpdateAccountResource, please refer to the document of DeviceUpdateAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
ResourceIdentifier deviceUpdateAccountResourceId = DeviceUpdateAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
DeviceUpdateAccountResource deviceUpdateAccount = client.GetDeviceUpdateAccountResource(deviceUpdateAccountResourceId);

// get the collection of this DeviceUpdateInstanceResource
DeviceUpdateInstanceCollection collection = deviceUpdateAccount.GetDeviceUpdateInstances();

// invoke the operation
string instanceName = "blue";
DeviceUpdateInstanceData data = new DeviceUpdateInstanceData(new AzureLocation("westus2"))
{
    IotHubs = { new DeviceUpdateIotHubSettings(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Devices/IotHubs/blue-contoso-hub")) },
    EnableDiagnostics = false,
    DiagnosticStorageProperties = new DiagnosticStorageProperties(DiagnosticStorageAuthenticationType.KeyBased, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/adu-resource-group/providers/Microsoft.Storage/storageAccounts/testAccount"))
    {
        ConnectionString = "string",
    },
};
ArmOperation<DeviceUpdateInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, instanceName, data);
DeviceUpdateInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceUpdateInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
