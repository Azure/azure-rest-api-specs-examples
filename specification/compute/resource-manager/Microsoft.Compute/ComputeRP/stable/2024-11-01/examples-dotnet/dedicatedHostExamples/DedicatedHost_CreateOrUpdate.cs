using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/dedicatedHostExamples/DedicatedHost_CreateOrUpdate.json
// this example is just showing the usage of "DedicatedHosts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHostGroupResource created on azure
// for more information of creating DedicatedHostGroupResource, please refer to the document of DedicatedHostGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string hostGroupName = "myDedicatedHostGroup";
ResourceIdentifier dedicatedHostGroupResourceId = DedicatedHostGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName);
DedicatedHostGroupResource dedicatedHostGroup = client.GetDedicatedHostGroupResource(dedicatedHostGroupResourceId);

// get the collection of this DedicatedHostResource
DedicatedHostCollection collection = dedicatedHostGroup.GetDedicatedHosts();

// invoke the operation
string hostName = "myDedicatedHost";
DedicatedHostData data = new DedicatedHostData(new AzureLocation("westus"), new ComputeSku
{
    Name = "DSv3-Type1",
})
{
    PlatformFaultDomain = 1,
    Tags =
    {
    ["department"] = "HR"
    },
};
ArmOperation<DedicatedHostResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, hostName, data);
DedicatedHostResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedHostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
