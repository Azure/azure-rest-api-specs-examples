using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_CreateorUpdate.json
// this example is just showing the usage of "Jobs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ContainerAppJobResource
ContainerAppJobCollection collection = resourceGroupResource.GetContainerAppJobs();

// invoke the operation
string jobName = "testcontainerappsjob0";
ContainerAppJobData data = new ContainerAppJobData(new AzureLocation("East US"))
{
    EnvironmentId = "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    Configuration = new ContainerAppJobConfiguration(ContainerAppJobTriggerType.Manual, 10)
    {
        ReplicaRetryLimit = 10,
        ManualTriggerConfig = new JobConfigurationManualTriggerConfig()
        {
            ReplicaCompletionCount = 1,
            Parallelism = 4,
        },
    },
    Template = new ContainerAppJobTemplate()
    {
        InitContainers =
        {
        new ContainerAppInitContainer()
        {
        Image = "repo/testcontainerappsjob0:v4",
        Name = "testinitcontainerAppsJob0",
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
        Cpu = 0.5,
        Memory = "1Gi",
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
        InitialDelaySeconds = 5,
        PeriodSeconds = 3,
        ProbeType = ContainerAppProbeType.Liveness,
        }
        },
        Image = "repo/testcontainerappsjob0:v1",
        Name = "testcontainerappsjob0",
        }
        },
    },
};
ArmOperation<ContainerAppJobResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, jobName, data);
ContainerAppJobResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppJobData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
