using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ContainerApps_Patch.json
// this example is just showing the usage of "ContainerApps_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppResource created on azure
// for more information of creating ContainerAppResource, please refer to the document of ContainerAppResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string containerAppName = "testcontainerapp0";
ResourceIdentifier containerAppResourceId = ContainerAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName);
ContainerAppResource containerApp = client.GetContainerAppResource(containerAppResourceId);

// invoke the operation
ContainerAppData data = new ContainerAppData(new AzureLocation("East US"))
{
    Configuration = new ContainerAppConfiguration
    {
        Ingress = new ContainerAppIngressConfiguration
        {
            External = true,
            TargetPort = 3000,
            Traffic = {new ContainerAppRevisionTrafficWeight
            {
            RevisionName = "testcontainerapp0-ab1234",
            Weight = 100,
            Label = "production",
            }},
            CustomDomains = {new ContainerAppCustomDomain("www.my-name.com")
            {
            BindingType = ContainerAppCustomDomainBindingType.SniEnabled,
            CertificateId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
            }, new ContainerAppCustomDomain("www.my-other-name.com")
            {
            BindingType = ContainerAppCustomDomainBindingType.SniEnabled,
            CertificateId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
            }},
            IPSecurityRestrictions = {new ContainerAppIPSecurityRestrictionRule("Allow work IP A subnet", "192.168.1.1/32", ContainerAppIPRuleAction.Allow)
            {
            Description = "Allowing all IP's within the subnet below to access containerapp",
            }, new ContainerAppIPSecurityRestrictionRule("Allow work IP B subnet", "192.168.1.1/8", ContainerAppIPRuleAction.Allow)
            {
            Description = "Allowing all IP's within the subnet below to access containerapp",
            }},
            StickySessionsAffinity = Affinity.Sticky,
        },
        Dapr = new ContainerAppDaprConfiguration
        {
            IsEnabled = true,
            AppProtocol = ContainerAppProtocol.Http,
            AppPort = 3000,
            HttpReadBufferSize = 30,
            HttpMaxRequestSize = 10,
            LogLevel = ContainerAppDaprLogLevel.Debug,
            IsApiLoggingEnabled = true,
        },
        EnableMetrics = true,
        MaxInactiveRevisions = 10,
        ServiceType = "redis",
    },
    Template = new ContainerAppTemplate
    {
        InitContainers = {new ContainerAppInitContainer
        {
        Image = "repo/testcontainerapp0:v4",
        Name = "testinitcontainerApp0",
        Resources = new AppContainerResources
        {
        Cpu = 0.5,
        Memory = "1Gi",
        },
        }},
        Containers = {new ContainerAppContainer
        {
        Probes = {new ContainerAppProbe
        {
        HttpGet = new ContainerAppHttpRequestInfo(8080)
        {
        HttpHeaders = {new ContainerAppHttpHeaderInfo("Custom-Header", "Awesome")},
        Path = "/health",
        },
        InitialDelaySeconds = 3,
        PeriodSeconds = 3,
        ProbeType = ContainerAppProbeType.Liveness,
        }},
        Image = "repo/testcontainerapp0:v1",
        Name = "testcontainerapp0",
        }},
        Scale = new ContainerAppScale
        {
            MinReplicas = 1,
            MaxReplicas = 5,
            CooldownPeriod = 350,
            PollingInterval = 35,
            Rules = {new ContainerAppScaleRule
            {
            Name = "httpscalingrule",
            Custom = new ContainerAppCustomScaleRule
            {
            CustomScaleRuleType = "http",
            Metadata =
            {
            ["concurrentRequests"] = "50"
            },
            },
            }},
        },
        ServiceBinds = {new ContainerAppServiceBind
        {
        ServiceId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service"),
        Name = "service",
        }},
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<ContainerAppResource> lro = await containerApp.UpdateAsync(WaitUntil.Completed, data);
ContainerAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
