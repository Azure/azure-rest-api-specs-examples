using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ConfigServers_UpdatePatch.json
// this example is just showing the usage of "ConfigServers_UpdatePatch" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformConfigServerResource created on azure
// for more information of creating AppPlatformConfigServerResource, please refer to the document of AppPlatformConfigServerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
ResourceIdentifier appPlatformConfigServerResourceId = AppPlatformConfigServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
AppPlatformConfigServerResource appPlatformConfigServer = client.GetAppPlatformConfigServerResource(appPlatformConfigServerResourceId);

// invoke the operation
AppPlatformConfigServerData data = new AppPlatformConfigServerData
{
    Properties = new AppPlatformConfigServerProperties
    {
        ConfigServerGitProperty = new AppPlatformConfigServerGitProperty(new Uri("https://github.com/fake-user/fake-repository.git"))
        {
            Label = "master",
            SearchPaths = { "/" },
        },
    },
};
ArmOperation<AppPlatformConfigServerResource> lro = await appPlatformConfigServer.UpdateAsync(WaitUntil.Completed, data);
AppPlatformConfigServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformConfigServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
