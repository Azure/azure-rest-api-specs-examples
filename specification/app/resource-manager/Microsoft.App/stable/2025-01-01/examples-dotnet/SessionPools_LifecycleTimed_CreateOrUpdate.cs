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

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/SessionPools_LifecycleTimed_CreateOrUpdate.json
// this example is just showing the usage of "ContainerAppsSessionPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SessionPoolResource
SessionPoolCollection collection = resourceGroupResource.GetSessionPools();

// invoke the operation
string sessionPoolName = "testsessionpool";
SessionPoolData data = new SessionPoolData(new AzureLocation("East US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    EnvironmentId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
    ContainerType = ContainerType.CustomContainer,
    PoolManagementType = PoolManagementType.Dynamic,
    ScaleConfiguration = new SessionPoolScaleConfiguration
    {
        MaxConcurrentSessions = 500,
        ReadySessionInstances = 100,
    },
    DynamicPoolLifecycleConfiguration = new SessionPoolLifecycleConfiguration
    {
        LifecycleType = SessionPoolLifecycleType.OnContainerExit,
        MaxAlivePeriodInSeconds = 86400,
    },
    CustomContainerTemplate = new CustomContainerTemplate
    {
        RegistryCredentials = new SessionRegistryCredentials
        {
            Server = "test.azurecr.io",
            Identity = "/subscriptions/7a497526-bb8d-4816-9795-db1418a1f977/resourcegroups/test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testSP",
        },
        Containers = {new SessionContainer
        {
        Image = "repo/testcontainer:v4",
        Name = "testinitcontainer",
        Command = {"/bin/sh"},
        Args = {"-c", "while true; do echo hello; sleep 10;done"},
        Resources = new SessionContainerResources
        {
        Cpu = 0.25,
        Memory = "0.5Gi",
        },
        }},
        IngressTargetPort = 80,
    },
    SessionNetworkStatus = SessionNetworkStatus.EgressEnabled,
    ManagedIdentitySettings = {new SessionPoolManagedIdentitySetting("system")
    {
    Lifecycle = ContainerAppIdentitySettingsLifeCycle.Main,
    }},
};
ArmOperation<SessionPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sessionPoolName, data);
SessionPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SessionPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
