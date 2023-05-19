using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/HostPool_Create.json
// this example is just showing the usage of "HostPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HostPoolResource
HostPoolCollection collection = resourceGroupResource.GetHostPools();

// invoke the operation
string hostPoolName = "hostPool1";
HostPoolData data = new HostPoolData(new AzureLocation("centralus"), HostPoolType.Pooled, HostPoolLoadBalancerType.BreadthFirst, PreferredAppGroupType.Desktop)
{
    FriendlyName = "friendly",
    Description = "des1",
    PersonalDesktopAssignmentType = PersonalDesktopAssignmentType.Automatic,
    CustomRdpProperty = null,
    MaxSessionLimit = 999999,
    RegistrationInfo = new HostPoolRegistrationInfo()
    {
        ExpireOn = DateTimeOffset.Parse("2020-10-01T14:01:54.9571247Z"),
        RegistrationTokenOperation = HostPoolRegistrationTokenOperation.Update,
    },
    VmTemplate = "{json:json}",
    SsoAdfsAuthority = "https://adfs",
    SsoClientId = "client",
    SsoClientSecretKeyVaultPath = "https://keyvault/secret",
    SsoSecretType = HostPoolSsoSecretType.SharedKey,
    StartVmOnConnect = false,
    AgentUpdate = new SessionHostAgentUpdateProperties()
    {
        UpdateType = SessionHostComponentUpdateType.Scheduled,
        DoesUseSessionHostLocalTime = false,
        MaintenanceWindowTimeZone = "Alaskan Standard Time",
        MaintenanceWindows =
        {
        new SessionHostMaintenanceWindowProperties()
        {
        Hour = 7,
        DayOfWeek = DesktopVirtualizationDayOfWeek.Friday,
        },new SessionHostMaintenanceWindowProperties()
        {
        Hour = 8,
        DayOfWeek = DesktopVirtualizationDayOfWeek.Saturday,
        }
        },
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ArmOperation<HostPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, hostPoolName, data);
HostPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HostPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
