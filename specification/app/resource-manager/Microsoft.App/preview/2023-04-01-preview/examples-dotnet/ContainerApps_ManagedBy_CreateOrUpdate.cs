using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2023-04-01-preview/examples/ContainerApps_ManagedBy_CreateOrUpdate.json
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
string containerAppName = "testcontainerAppManagedBy";
ContainerAppData data = new ContainerAppData(new AzureLocation("East US"))
{
    ManagedBy = "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp",
    EnvironmentId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
    Configuration = new ContainerAppConfiguration()
    {
        Ingress = new ContainerAppIngressConfiguration()
        {
            External = true,
            TargetPort = 3000,
            ExposedPort = 4000,
            Transport = ContainerAppIngressTransportMethod.Tcp,
            Traffic =
            {
            new ContainerAppRevisionTrafficWeight()
            {
            RevisionName = "testcontainerAppManagedBy-ab1234",
            Weight = 100,
            }
            },
        },
    },
    Template = new ContainerAppTemplate()
    {
        Containers =
        {
        new ContainerAppContainer()
        {
        Probes =
        {
        new ContainerAppProbe()
        {
        InitialDelaySeconds = 3,
        PeriodSeconds = 3,
        TcpSocket = new ContainerAppTcpSocketRequestInfo(8080),
        ProbeType = ContainerAppProbeType.Liveness,
        }
        },
        Image = "repo/testcontainerAppManagedBy:v1",
        Name = "testcontainerAppManagedBy",
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
            Name = "tcpscalingrule",
            Tcp = new ContainerAppTcpScaleRule()
            {
            Metadata =
            {
            ["concurrentConnections"] = "50",
            },
            },
            }
            },
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
