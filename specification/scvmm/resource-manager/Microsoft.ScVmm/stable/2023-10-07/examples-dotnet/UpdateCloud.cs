using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/UpdateCloud.json
// this example is just showing the usage of "Clouds_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmCloudResource created on azure
// for more information of creating ScVmmCloudResource, please refer to the document of ScVmmCloudResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string cloudResourceName = "HRCloud";
ResourceIdentifier scVmmCloudResourceId = ScVmmCloudResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudResourceName);
ScVmmCloudResource scVmmCloud = client.GetScVmmCloudResource(scVmmCloudResourceId);

// invoke the operation
ScVmmResourcePatch patch = new ScVmmResourcePatch
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<ScVmmCloudResource> lro = await scVmmCloud.UpdateAsync(WaitUntil.Completed, patch);
ScVmmCloudResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScVmmCloudData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
