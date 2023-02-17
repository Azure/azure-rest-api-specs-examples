using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildServiceBuilder_CreateOrUpdate.json
// this example is just showing the usage of "BuildServiceBuilder_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformBuilderResource created on azure
// for more information of creating AppPlatformBuilderResource, please refer to the document of AppPlatformBuilderResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string buildServiceName = "default";
string builderName = "mybuilder";
ResourceIdentifier appPlatformBuilderResourceId = AppPlatformBuilderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName, builderName);
AppPlatformBuilderResource appPlatformBuilder = client.GetAppPlatformBuilderResource(appPlatformBuilderResourceId);

// invoke the operation
AppPlatformBuilderData data = new AppPlatformBuilderData()
{
    Properties = new AppPlatformBuilderProperties()
    {
        Stack = new AppPlatformClusterStackProperties()
        {
            Id = "io.buildpacks.stacks.bionic",
            Version = "base",
        },
        BuildpackGroups =
        {
        new BuildpacksGroupProperties()
        {
        Name = "mix",
        Buildpacks =
        {
        new WritableSubResource()
        {
        Id = new ResourceIdentifier("tanzu-buildpacks/java-azure"),
        }
        },
        }
        },
    },
};
ArmOperation<AppPlatformBuilderResource> lro = await appPlatformBuilder.UpdateAsync(WaitUntil.Completed, data);
AppPlatformBuilderResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformBuilderData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
