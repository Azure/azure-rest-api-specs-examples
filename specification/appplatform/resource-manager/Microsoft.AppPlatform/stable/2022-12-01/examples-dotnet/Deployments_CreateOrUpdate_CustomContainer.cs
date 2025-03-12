using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_CreateOrUpdate_CustomContainer.json
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
AppPlatformDeploymentData data = new AppPlatformDeploymentData
{
    Properties = new AppPlatformDeploymentProperties
    {
        Source = new AppPlatformCustomContainerUserSourceInfo
        {
            CustomContainer = new AppPlatformCustomContainer
            {
                Server = "myacr.azurecr.io",
                ContainerImage = "myContainerImage:v1",
                Command = { "/bin/sh" },
                Args = { "-c", "while true; do echo hello; sleep 10;done" },
                ImageRegistryCredential = new AppPlatformImageRegistryCredential
                {
                    Username = "myUsername",
                    Password = "myPassword",
                },
                LanguageFramework = "springboot",
            },
        },
        DeploymentSettings = new AppPlatformDeploymentSettings
        {
            ResourceRequests = new AppPlatformDeploymentResourceRequirements
            {
                Cpu = "1000m",
                Memory = "3Gi",
            },
            EnvironmentVariables =
            {
            ["env"] = "test"
            },
            LivenessProbe = new AppInstanceProbe(false)
            {
                ProbeAction = new AppInstanceHttpGetAction
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
                ProbeAction = new AppInstanceHttpGetAction
                {
                    Path = "/health",
                    Scheme = AppInstanceHttpSchemeType.Http,
                },
                InitialDelayInSeconds = 30,
                PeriodInSeconds = 10,
                FailureThreshold = 3,
            },
            StartupProbe = default,
            TerminationGracePeriodInSeconds = 30,
        },
    },
};
ArmOperation<AppPlatformDeploymentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deploymentName, data);
AppPlatformDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
