using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_Update.json
// this example is just showing the usage of "Environments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabEnvironmentResource created on azure
// for more information of creating DevTestLabEnvironmentResource, please refer to the document of DevTestLabEnvironmentResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string userName = "@me";
string name = "{environmentName}";
ResourceIdentifier devTestLabEnvironmentResourceId = DevTestLabEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName, name);
DevTestLabEnvironmentResource devTestLabEnvironment = client.GetDevTestLabEnvironmentResource(devTestLabEnvironmentResourceId);

// invoke the operation
DevTestLabEnvironmentPatch patch = new DevTestLabEnvironmentPatch
{
    Tags =
    {
    ["tagName1"] = "tagValue1"
    },
};
DevTestLabEnvironmentResource result = await devTestLabEnvironment.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
