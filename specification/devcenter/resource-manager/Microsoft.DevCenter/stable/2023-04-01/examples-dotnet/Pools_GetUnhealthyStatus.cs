using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Pools_GetUnhealthyStatus.json
// this example is just showing the usage of "Pools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterProjectResource created on azure
// for more information of creating DevCenterProjectResource, please refer to the document of DevCenterProjectResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string projectName = "DevProject";
ResourceIdentifier devCenterProjectResourceId = DevCenterProjectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName);
DevCenterProjectResource devCenterProject = client.GetDevCenterProjectResource(devCenterProjectResourceId);

// get the collection of this DevCenterPoolResource
DevCenterPoolCollection collection = devCenterProject.GetDevCenterPools();

// invoke the operation
string poolName = "DevPool";
NullableResponse<DevCenterPoolResource> response = await collection.GetIfExistsAsync(poolName);
DevCenterPoolResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DevCenterPoolData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
