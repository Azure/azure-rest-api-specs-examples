using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildpackBinding_Delete.json
// this example is just showing the usage of "BuildpackBinding_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformBuildpackBindingResource created on azure
// for more information of creating AppPlatformBuildpackBindingResource, please refer to the document of AppPlatformBuildpackBindingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string buildServiceName = "default";
string builderName = "default";
string buildpackBindingName = "myBuildpackBinding";
ResourceIdentifier appPlatformBuildpackBindingResourceId = AppPlatformBuildpackBindingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName);
AppPlatformBuildpackBindingResource appPlatformBuildpackBinding = client.GetAppPlatformBuildpackBindingResource(appPlatformBuildpackBindingResourceId);

// invoke the operation
await appPlatformBuildpackBinding.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
