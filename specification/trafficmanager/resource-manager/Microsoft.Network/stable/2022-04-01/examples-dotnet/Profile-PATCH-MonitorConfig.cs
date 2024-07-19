using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.TrafficManager.Models;
using Azure.ResourceManager.TrafficManager;

// Generated from example definition: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-PATCH-MonitorConfig.json
// this example is just showing the usage of "Profiles_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficManagerProfileResource created on azure
// for more information of creating TrafficManagerProfileResource, please refer to the document of TrafficManagerProfileResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "azuresdkfornetautoresttrafficmanager2583";
string profileName = "azuresdkfornetautoresttrafficmanager6192";
ResourceIdentifier trafficManagerProfileResourceId = TrafficManagerProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
TrafficManagerProfileResource trafficManagerProfile = client.GetTrafficManagerProfileResource(trafficManagerProfileResourceId);

// invoke the operation
TrafficManagerProfileData data = new TrafficManagerProfileData()
{
    MonitorConfig = new TrafficManagerMonitorConfig()
    {
        Protocol = TrafficManagerMonitorProtocol.Http,
        Port = 80,
        Path = "/testpath.aspx",
        IntervalInSeconds = 30,
        TimeoutInSeconds = 6,
        ToleratedNumberOfFailures = 4,
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
    },
};
TrafficManagerProfileResource result = await trafficManagerProfile.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrafficManagerProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
