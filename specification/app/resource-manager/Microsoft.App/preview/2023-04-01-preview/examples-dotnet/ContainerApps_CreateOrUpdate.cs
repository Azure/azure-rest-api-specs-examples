using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2023-04-01-preview/examples/ContainerApps_CreateOrUpdate.json
// this example is just showing the usage of "ContainerApps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerAppResource
ContainerAppCollection collection = resourceGroupResource.GetContainerApps();

// invoke the operation
string containerAppName = "testcontainerApp0";
ContainerAppData data = new ContainerAppData(new AzureLocation("East US"))
{
    EnvironmentId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
    WorkloadProfileName = "My-GP-01",
    Configuration = new ContainerAppConfiguration()
    {
        Ingress = new ContainerAppIngressConfiguration()
        {
            External = true,
            TargetPort = 3000,
            Traffic =
            {
            new ContainerAppRevisionTrafficWeight()
            {
            RevisionName = "testcontainerApp0-ab1234",
            Weight = 100,
            Label = "production",
            }
            },
            CustomDomains =
            {
            new ContainerAppCustomDomain("www.my-name.com")
            {
            BindingType = ContainerAppCustomDomainBindingType.SniEnabled,
            CertificateId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com"),
            },new ContainerAppCustomDomain("www.my-other-name.com")
            {
            BindingType = ContainerAppCustomDomainBindingType.SniEnabled,
            CertificateId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com"),
            }
            },
            IPSecurityRestrictions =
            {
            new ContainerAppIPSecurityRestrictionRule("Allow work IP A subnet","192.168.1.1/32",ContainerAppIPRuleAction.Allow)
            {
            Description = "Allowing all IP's within the subnet below to access containerapp",
            },new ContainerAppIPSecurityRestrictionRule("Allow work IP B subnet","192.168.1.1/8",ContainerAppIPRuleAction.Allow)
            {
            Description = "Allowing all IP's within the subnet below to access containerapp",
            }
            },
            StickySessionsAffinity = Affinity.Sticky,
            ClientCertificateMode = ContainerAppIngressClientCertificateMode.Accept,
            CorsPolicy = new ContainerAppCorsPolicy(new string[]
{
"https://a.test.com","https://b.test.com"
})
            {
                AllowedMethods =
                {
                "GET","POST"
                },
                AllowedHeaders =
                {
                "HEADER1","HEADER2"
                },
                ExposeHeaders =
                {
                "HEADER3","HEADER4"
                },
                MaxAge = 1234,
                AllowCredentials = true,
            },
        },
        Dapr = new ContainerAppDaprConfiguration()
        {
            IsEnabled = true,
            AppProtocol = ContainerAppProtocol.Http,
            AppPort = 3000,
            HttpReadBufferSize = 30,
            HttpMaxRequestSize = 10,
            LogLevel = ContainerAppDaprLogLevel.Debug,
            IsApiLoggingEnabled = true,
        },
        MaxInactiveRevisions = 10,
        ServiceType = "redis",
    },
    Template = new ContainerAppTemplate()
    {
        InitContainers =
        {
        new ContainerAppInitContainer()
        {
        Image = "repo/testcontainerApp0:v4",
        Name = "testinitcontainerApp0",
        Command =
        {
        "/bin/sh"
        },
        Args =
        {
        "-c","while true; do echo hello; sleep 10;done"
        },
        Resources = new AppContainerResources()
        {
        Cpu = 0.2,
        Memory = "100Mi",
        },
        }
        },
        Containers =
        {
        new ContainerAppContainer()
        {
        Probes =
        {
        new ContainerAppProbe()
        {
        HttpGet = new ContainerAppHttpRequestInfo(8080)
        {
        HttpHeaders =
        {
        new ContainerAppHttpHeaderInfo("Custom-Header","Awesome")
        },
        Path = "/health",
        },
        InitialDelaySeconds = 3,
        PeriodSeconds = 3,
        ProbeType = ContainerAppProbeType.Liveness,
        }
        },
        Image = "repo/testcontainerApp0:v1",
        Name = "testcontainerApp0",
        }
        },
        Scale = new ContainerAppScale()
        {
            MinReplicas = 1,
            MaxReplicas = 5,
            Rules =
            {
            new ContainerAppScaleRule()
            {
            Name = "httpscalingrule",
            Custom = new ContainerAppCustomScaleRule()
            {
            CustomScaleRuleType = "http",
            Metadata =
            {
            ["concurrentRequests"] = "50",
            },
            },
            }
            },
        },
        ServiceBinds =
        {
        new ServiceBind()
        {
        ServiceId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/redisService"),
        Name = "redisService",
        }
        },
    },
};
ArmOperation<ContainerAppResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerAppName, data);
ContainerAppResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
