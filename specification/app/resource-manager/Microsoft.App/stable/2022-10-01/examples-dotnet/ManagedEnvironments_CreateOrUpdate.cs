using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2022-10-01/examples/ManagedEnvironments_CreateOrUpdate.json
// this example is just showing the usage of "ManagedEnvironments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerAppManagedEnvironmentResource
ContainerAppManagedEnvironmentCollection collection = resourceGroupResource.GetContainerAppManagedEnvironments();

// invoke the operation
string environmentName = "testcontainerenv";
ContainerAppManagedEnvironmentData data = new ContainerAppManagedEnvironmentData(new AzureLocation("East US"))
{
    Kind = "serverless",
    SkuName = AppContainersSkuName.Premium,
    DaprAIConnectionString = "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/",
    VnetConfiguration = new ContainerAppVnetConfiguration()
    {
        OutboundSettings = new ContainerAppManagedEnvironmentOutboundSettings()
        {
            OutBoundType = ContainerAppManagedEnvironmentOutBoundType.UserDefinedRouting,
            VirtualNetworkApplianceIP = IPAddress.Parse("192.168.1.20"),
        },
    },
    AppLogsConfiguration = new ContainerAppLogsConfiguration()
    {
        LogAnalyticsConfiguration = new ContainerAppLogAnalyticsConfiguration()
        {
            CustomerId = "string",
            SharedKey = "string",
        },
    },
    IsZoneRedundant = true,
    CustomDomainConfiguration = new ContainerAppCustomDomainConfiguration()
    {
        DnsSuffix = "www.my-name.com",
        CertificateValue = Convert.FromBase64String("Y2VydA=="),
        CertificatePassword = "private key password",
    },
    WorkloadProfiles =
    {
    new ContainerAppWorkloadProfile("GeneralPurpose",3,12),new ContainerAppWorkloadProfile("MemoryOptimized",3,6),new ContainerAppWorkloadProfile("ComputeOptimized",3,6)
    },
};
ArmOperation<ContainerAppManagedEnvironmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, environmentName, data);
ContainerAppManagedEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppManagedEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
