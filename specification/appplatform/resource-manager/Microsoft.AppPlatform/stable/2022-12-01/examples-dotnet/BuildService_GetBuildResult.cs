using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_GetBuildResult.json
// this example is just showing the usage of "BuildService_GetBuildResult" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformBuildResource created on azure
// for more information of creating AppPlatformBuildResource, please refer to the document of AppPlatformBuildResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string buildServiceName = "default";
string buildName = "mybuild";
ResourceIdentifier appPlatformBuildResourceId = AppPlatformBuildResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName, buildName);
AppPlatformBuildResource appPlatformBuild = client.GetAppPlatformBuildResource(appPlatformBuildResourceId);

// get the collection of this AppPlatformBuildResultResource
AppPlatformBuildResultCollection collection = appPlatformBuild.GetAppPlatformBuildResults();

// invoke the operation
string buildResultName = "123";
bool result = await collection.ExistsAsync(buildResultName);

Console.WriteLine($"Succeeded: {result}");
