using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_Patch.json
// this example is just showing the usage of "BareMetalMachines_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BareMetalMachineResource created on azure
// for more information of creating BareMetalMachineResource, please refer to the document of BareMetalMachineResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string bareMetalMachineName = "bareMetalMachineName";
ResourceIdentifier bareMetalMachineResourceId = BareMetalMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, bareMetalMachineName);
BareMetalMachineResource bareMetalMachine = client.GetBareMetalMachineResource(bareMetalMachineResourceId);

// invoke the operation
BareMetalMachinePatch patch = new BareMetalMachinePatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
    MachineDetails = "machinedetails",
};
ArmOperation<BareMetalMachineResource> lro = await bareMetalMachine.UpdateAsync(WaitUntil.Completed, patch);
BareMetalMachineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BareMetalMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
