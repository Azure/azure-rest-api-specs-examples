using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-04-15/examples/AFDOriginGroups_Create.json
// this example is just showing the usage of "FrontDoorOriginGroups_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProfileResource created on azure
// for more information of creating ProfileResource, please refer to the document of ProfileResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
ResourceIdentifier profileResourceId = ProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName);
ProfileResource profile = client.GetProfileResource(profileResourceId);

// get the collection of this FrontDoorOriginGroupResource
FrontDoorOriginGroupCollection collection = profile.GetFrontDoorOriginGroups();

// invoke the operation
string originGroupName = "origingroup1";
FrontDoorOriginGroupData data = new FrontDoorOriginGroupData
{
    LoadBalancingSettings = new LoadBalancingSettings
    {
        SampleSize = 3,
        SuccessfulSamplesRequired = 3,
        AdditionalLatencyInMilliseconds = 1000,
    },
    HealthProbeSettings = new HealthProbeSettings
    {
        ProbePath = "/path2",
        ProbeRequestType = HealthProbeRequestType.NotSet,
        ProbeProtocol = HealthProbeProtocol.NotSet,
        ProbeIntervalInSeconds = 10,
    },
    TrafficRestorationTimeInMinutes = 5,
};
ArmOperation<FrontDoorOriginGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, originGroupName, data);
FrontDoorOriginGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorOriginGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
