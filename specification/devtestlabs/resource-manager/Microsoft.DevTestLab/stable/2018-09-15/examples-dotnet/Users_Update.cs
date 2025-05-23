using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_Update.json
// this example is just showing the usage of "Users_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabUserResource created on azure
// for more information of creating DevTestLabUserResource, please refer to the document of DevTestLabUserResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{devtestlabName}";
string name = "{userName}";
ResourceIdentifier devTestLabUserResourceId = DevTestLabUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, name);
DevTestLabUserResource devTestLabUser = client.GetDevTestLabUserResource(devTestLabUserResourceId);

// invoke the operation
DevTestLabUserPatch patch = new DevTestLabUserPatch
{
    Tags =
    {
    ["tagName1"] = "tagValue1"
    },
};
DevTestLabUserResource result = await devTestLabUser.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabUserData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
