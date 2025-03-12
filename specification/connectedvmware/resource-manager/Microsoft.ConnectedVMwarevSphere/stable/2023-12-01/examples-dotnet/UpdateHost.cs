using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-12-01/examples/UpdateHost.json
// this example is just showing the usage of "Hosts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VMwareHostResource created on azure
// for more information of creating VMwareHostResource, please refer to the document of VMwareHostResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string hostName = "HRHost";
ResourceIdentifier vMwareHostResourceId = VMwareHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostName);
VMwareHostResource vMwareHost = client.GetVMwareHostResource(vMwareHostResourceId);

// invoke the operation
VMwareResourcePatchContent content = new VMwareResourcePatchContent
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
VMwareHostResource result = await vMwareHost.UpdateAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VMwareHostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
