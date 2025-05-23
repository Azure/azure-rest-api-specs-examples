using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildService_CreateOrUpdateBuild.json
// this example is just showing the usage of "BuildService_CreateOrUpdateBuild" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AppPlatformBuildResource
AppPlatformBuildCollection collection = appPlatformBuildService.GetAppPlatformBuilds();

// invoke the operation
string buildName = "mybuild";
AppPlatformBuildData data = new AppPlatformBuildData
{
    Properties = new AppPlatformBuildProperties
    {
        RelativePath = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-20210601-3ed9f4a2-986b-4bbd-b833-a42dccb2f777",
        Builder = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default",
        AgentPool = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/agentPools/default",
        Env =
        {
        ["environmentVariable"] = "test"
        },
        ResourceRequests = new AppPlatformBuildResourceRequirements
        {
            Cpu = "1",
            Memory = "2Gi",
        },
    },
};
ArmOperation<AppPlatformBuildResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, buildName, data);
AppPlatformBuildResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformBuildData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
