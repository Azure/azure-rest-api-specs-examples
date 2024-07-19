using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager.Models;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PUT-WithCustomHeaders.json
// this example is just showing the usage of "Profiles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager2583";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this TrafficManagerProfileResource
TrafficManagerProfileCollection collection = resourceGroupResource.GetTrafficManagerProfiles();

// invoke the operation
string profileName = "azuresdkfornetautoresttrafficmanager6192";
TrafficManagerProfileData data = new TrafficManagerProfileData()
{
    ProfileStatus = TrafficManagerProfileStatus.Enabled,
    TrafficRoutingMethod = TrafficRoutingMethod.Performance,
    DnsConfig = new TrafficManagerDnsConfig()
    {
        RelativeName = "azuresdkfornetautoresttrafficmanager6192",
        Ttl = 35,
    },
    MonitorConfig = new TrafficManagerMonitorConfig()
    {
        Protocol = TrafficManagerMonitorProtocol.Http,
        Port = 80,
        Path = "/testpath.aspx",
        IntervalInSeconds = 10,
        TimeoutInSeconds = 5,
        ToleratedNumberOfFailures = 2,
        CustomHeaders =
        {
        new TrafficManagerMonitorConfigCustomHeaderInfo()
        {
        Name = "header-1",
        Value = "value-1",
        },new TrafficManagerMonitorConfigCustomHeaderInfo()
        {
        Name = "header-2",
        Value = "value-2",
        }
        },
        ExpectedStatusCodeRanges =
        {
        new ExpectedStatusCodeRangeInfo()
        {
        Min = 200,
        Max = 205,
        },new ExpectedStatusCodeRangeInfo()
        {
        Min = 400,
        Max = 410,
        }
        },
    },
    Endpoints =
    {
    new TrafficManagerEndpointData()
    {
    Target = "foobar.contoso.com",
    EndpointStatus = TrafficManagerEndpointStatus.Enabled,
    EndpointLocation = "North Europe",
    CustomHeaders =
    {
    new TrafficManagerEndpointCustomHeaderInfo()
    {
    Name = "header-2",
    Value = "value-2-overridden",
    }
    },
    Name = "My external endpoint",
    ResourceType = new ResourceType("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
    }
    },
    TrafficViewEnrollmentStatus = TrafficViewEnrollmentStatus.Disabled,
    Location = new AzureLocation("global"),
};
ArmOperation<TrafficManagerProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, profileName, data);
TrafficManagerProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficManagerProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
