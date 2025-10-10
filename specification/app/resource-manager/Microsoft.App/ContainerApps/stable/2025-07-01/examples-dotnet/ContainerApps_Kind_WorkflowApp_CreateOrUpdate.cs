using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ContainerApps_Kind_WorkflowApp_CreateOrUpdate.json
// this example is just showing the usage of "ContainerApps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerAppResource
ContainerAppCollection collection = resourceGroupResource.GetContainerApps();

// invoke the operation
string containerAppName = "testcontainerAppKind";
ContainerAppData data = new ContainerAppData(new AzureLocation("East Us"))
{
    ManagedBy = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind",
    Kind = ContainerAppKind.Workflowapp,
    ManagedEnvironmentId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3"),
    Configuration = new ContainerAppConfiguration
    {
        ActiveRevisionsMode = ContainerAppActiveRevisionsMode.Single,
        Ingress = new ContainerAppIngressConfiguration
        {
            External = true,
            TargetPort = 443,
            AllowInsecure = false,
        },
    },
    Template = new ContainerAppTemplate
    {
        Containers = {new ContainerAppContainer
        {
        Image = "default/logicapps-base:latest",
        Name = "logicapps-container",
        Resources = new AppContainerResources
        {
        Cpu = 1,
        Memory = "2.0Gi",
        },
        }},
        Scale = new ContainerAppScale
        {
            MinReplicas = 1,
            MaxReplicas = 30,
            CooldownPeriod = 350,
            PollingInterval = 35,
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
