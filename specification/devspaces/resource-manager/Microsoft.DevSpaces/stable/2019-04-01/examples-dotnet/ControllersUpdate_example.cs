using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevSpaces.Models;
using Azure.ResourceManager.DevSpaces;

// Generated from example definition: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersUpdate_example.json
// this example is just showing the usage of "Controllers_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ControllerResource created on azure
// for more information of creating ControllerResource, please refer to the document of ControllerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string name = "myControllerResource";
ResourceIdentifier controllerResourceId = ControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
ControllerResource controller = client.GetControllerResource(controllerResourceId);

// invoke the operation
ControllerPatch patch = new ControllerPatch
{
    Tags =
    {
    ["key"] = "value"
    },
    TargetContainerHostCredentialsBase64 = "QmFzZTY0IEVuY29kZWQgVmFsdWUK",
};
ControllerResource result = await controller.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ControllerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
