using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildpackBinding_Get.json
// this example is just showing the usage of "BuildpackBinding_Get" operation, for the dependent resources, they will have to be created separately.

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
string builderName = "default";
ResourceIdentifier appPlatformBuilderResourceId = AppPlatformBuilderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName, builderName);
AppPlatformBuilderResource appPlatformBuilder = client.GetAppPlatformBuilderResource(appPlatformBuilderResourceId);

// get the collection of this AppPlatformBuildpackBindingResource
AppPlatformBuildpackBindingCollection collection = appPlatformBuilder.GetAppPlatformBuildpackBindings();

// invoke the operation
string buildpackBindingName = "myBuildpackBinding";
NullableResponse<AppPlatformBuildpackBindingResource> response = await collection.GetIfExistsAsync(buildpackBindingName);
AppPlatformBuildpackBindingResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppPlatformBuildpackBindingData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
