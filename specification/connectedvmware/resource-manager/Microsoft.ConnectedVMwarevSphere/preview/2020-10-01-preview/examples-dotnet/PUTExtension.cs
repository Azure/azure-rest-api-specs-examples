using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2020-10-01-preview/examples/PUTExtension.json
// this example is just showing the usage of "MachineExtensions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string name = "myMachine";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// get the collection of this MachineExtensionResource
MachineExtensionCollection collection = virtualMachine.GetMachineExtensions();

// invoke the operation
string extensionName = "CustomScriptExtension";
MachineExtensionData data = new MachineExtensionData(new AzureLocation("eastus2euap"))
{
    Publisher = "Microsoft.Compute",
    MachineExtensionType = "CustomScriptExtension",
    TypeHandlerVersion = "1.10",
    Settings = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["commandToExecute"] = "powershell.exe -c \"Get-Process | Where-Object { $_.CPU -gt 10000 }\""
    }),
};
ArmOperation<MachineExtensionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, extensionName, data);
MachineExtensionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
