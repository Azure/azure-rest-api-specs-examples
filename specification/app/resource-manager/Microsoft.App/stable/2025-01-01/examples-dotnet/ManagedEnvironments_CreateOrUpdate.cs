using Azure;
using Azure.ResourceManager;
using System;
using System.Text;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironments_CreateOrUpdate.json
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
    Identity = new ManagedServiceIdentity("SystemAssigned, UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity")] = new UserAssignedIdentity()
        },
    },
    DaprAIConnectionString = "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/",
    VnetConfiguration = new ContainerAppVnetConfiguration
    {
        InfrastructureSubnetId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/RGName/providers/Microsoft.Network/virtualNetworks/VNetName/subnets/subnetName1"),
    },
    AppLogsConfiguration = new ContainerAppLogsConfiguration
    {
        LogAnalyticsConfiguration = new ContainerAppLogAnalyticsConfiguration
        {
            CustomerId = "string",
            SharedKey = "string",
        },
    },
    IsZoneRedundant = true,
    CustomDomainConfiguration = new ContainerAppCustomDomainConfiguration
    {
        DnsSuffix = "www.my-name.com",
        CertificateValue = Encoding.UTF8.GetBytes("Y2VydA=="),
        CertificatePassword = "1234",
    },
    WorkloadProfiles = {new ContainerAppWorkloadProfile("My-GP-01", "GeneralPurpose")
    {
    MinimumNodeCount = 3,
    MaximumNodeCount = 12,
    }, new ContainerAppWorkloadProfile("My-MO-01", "MemoryOptimized")
    {
    MinimumNodeCount = 3,
    MaximumNodeCount = 6,
    }, new ContainerAppWorkloadProfile("My-CO-01", "ComputeOptimized")
    {
    MinimumNodeCount = 3,
    MaximumNodeCount = 6,
    }, new ContainerAppWorkloadProfile("My-consumption-01", "Consumption")},
    IsMtlsEnabled = true,
    IsEnabled = true,
};
ArmOperation<ContainerAppManagedEnvironmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, environmentName, data);
ContainerAppManagedEnvironmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppManagedEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
