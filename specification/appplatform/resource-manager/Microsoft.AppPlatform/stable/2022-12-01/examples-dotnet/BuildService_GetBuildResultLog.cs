using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_GetBuildResultLog.json
// this example is just showing the usage of "BuildService_GetBuildResultLog" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformBuildResultResource created on azure
// for more information of creating AppPlatformBuildResultResource, please refer to the document of AppPlatformBuildResultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string buildServiceName = "default";
string buildName = "mybuild";
string buildResultName = "123";
ResourceIdentifier appPlatformBuildResultResourceId = AppPlatformBuildResultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName, buildName, buildResultName);
AppPlatformBuildResultResource appPlatformBuildResult = client.GetAppPlatformBuildResultResource(appPlatformBuildResultResourceId);

// invoke the operation
AppPlatformBuildResultLog result = await appPlatformBuildResult.GetBuildResultLogAsync();

Console.WriteLine($"Succeeded: {result}");
