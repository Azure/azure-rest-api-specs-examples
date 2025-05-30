using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/JavaComponents_CreateOrUpdate.json
// this example is just showing the usage of "JavaComponents_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentResource created on azure
// for more information of creating ContainerAppManagedEnvironmentResource, please refer to the document of ContainerAppManagedEnvironmentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "myenvironment";
ResourceIdentifier containerAppManagedEnvironmentResourceId = ContainerAppManagedEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName);
ContainerAppManagedEnvironmentResource containerAppManagedEnvironment = client.GetContainerAppManagedEnvironmentResource(containerAppManagedEnvironmentResourceId);

// get the collection of this JavaComponentResource
JavaComponentCollection collection = containerAppManagedEnvironment.GetJavaComponents();

// invoke the operation
string name = "myjavacomponent";
JavaComponentData data = new JavaComponentData
{
    Properties = new SpringBootAdminComponent
    {
        Configurations = {new JavaComponentConfigurationProperty
        {
        PropertyName = "spring.boot.admin.ui.enable-toasts",
        Value = "true",
        }, new JavaComponentConfigurationProperty
        {
        PropertyName = "spring.boot.admin.monitor.status-interval",
        Value = "10000ms",
        }},
        Scale = new JavaComponentPropertiesScale
        {
            MinReplicas = 1,
            MaxReplicas = 1,
        },
    },
};
ArmOperation<JavaComponentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
JavaComponentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
JavaComponentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
