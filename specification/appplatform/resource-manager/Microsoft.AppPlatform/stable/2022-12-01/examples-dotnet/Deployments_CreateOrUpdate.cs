using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_CreateOrUpdate.json
// this example is just showing the usage of "Deployments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformAppResource created on azure
// for more information of creating AppPlatformAppResource, please refer to the document of AppPlatformAppResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string appName = "myapp";
ResourceIdentifier appPlatformAppResourceId = AppPlatformAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, appName);
AppPlatformAppResource appPlatformApp = client.GetAppPlatformAppResource(appPlatformAppResourceId);

// get the collection of this AppPlatformDeploymentResource
AppPlatformDeploymentCollection collection = appPlatformApp.GetAppPlatformDeployments();

// invoke the operation
string deploymentName = "mydeployment";
AppPlatformDeploymentData data = new AppPlatformDeploymentData()
{
    Properties = new AppPlatformDeploymentProperties()
    {
        Source = new SourceUploadedUserSourceInfo()
        {
            ArtifactSelector = "sub-module-1",
            RelativePath = "resources/a172cedcae47474b615c54d510a5d84a8dea3032e958587430b413538be3f333-2019082605-e3095339-1723-44b7-8b5e-31b1003978bc",
            Version = "1.0",
        },
        DeploymentSettings = new AppPlatformDeploymentSettings()
        {
            ResourceRequests = new AppPlatformDeploymentResourceRequirements()
            {
                Cpu = "1000m",
                Memory = "3Gi",
            },
            EnvironmentVariables =
            {
            ["env"] = "test",
            },
            AddonConfigs =
            {
            ["ApplicationConfigurationService"] = new Dictionary<string, BinaryData>()
            {
            ["patterns"] = BinaryData.FromObjectAsJson(new object[] { "mypattern" }),
            },
            },
            LivenessProbe = new AppInstanceProbe(false)
            {
                ProbeAction = new AppInstanceHttpGetAction()
                {
                    Path = "/health",
                    Scheme = AppInstanceHttpSchemeType.Http,
                },
                InitialDelayInSeconds = 30,
                PeriodInSeconds = 10,
                FailureThreshold = 3,
            },
            ReadinessProbe = new AppInstanceProbe(false)
            {
                ProbeAction = new AppInstanceHttpGetAction()
                {
                    Path = "/health",
                    Scheme = AppInstanceHttpSchemeType.Http,
                },
                InitialDelayInSeconds = 30,
                PeriodInSeconds = 10,
                FailureThreshold = 3,
            },
            StartupProbe = null,
            TerminationGracePeriodInSeconds = 30,
        },
    },
    Sku = new AppPlatformSku()
    {
        Name = "S0",
        Tier = "Standard",
        Capacity = 1,
    },
};
ArmOperation<AppPlatformDeploymentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deploymentName, data);
AppPlatformDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
