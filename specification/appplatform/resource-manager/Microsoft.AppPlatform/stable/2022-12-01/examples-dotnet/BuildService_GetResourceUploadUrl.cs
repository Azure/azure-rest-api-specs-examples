using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_GetResourceUploadUrl.json
// this example is just showing the usage of "BuildService_GetResourceUploadUri" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformBuildServiceResource created on azure
// for more information of creating AppPlatformBuildServiceResource, please refer to the document of AppPlatformBuildServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string buildServiceName = "default";
ResourceIdentifier appPlatformBuildServiceResourceId = AppPlatformBuildServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, buildServiceName);
AppPlatformBuildServiceResource appPlatformBuildService = client.GetAppPlatformBuildServiceResource(appPlatformBuildServiceResourceId);

// invoke the operation
AppPlatformResourceUploadResult result = await appPlatformBuildService.GetResourceUploadUriAsync();

Console.WriteLine($"Succeeded: {result}");
