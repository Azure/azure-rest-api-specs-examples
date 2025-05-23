using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Orbital;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileCreate.json
// this example is just showing the usage of "ContactProfiles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
string resourceGroupName = "contoso-Rgp";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this OrbitalContactProfileResource
OrbitalContactProfileCollection collection = resourceGroupResource.GetOrbitalContactProfiles();

// invoke the operation
string contactProfileName = "CONTOSO-CP";
OrbitalContactProfileData data = new OrbitalContactProfileData(new AzureLocation("eastus2"))
{
    MinimumViableContactDuration = XmlConvert.ToTimeSpan("PT1M"),
    MinimumElevationDegrees = 5,
    AutoTrackingConfiguration = AutoTrackingConfiguration.Disabled,
    EventHubUri = new Uri("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.EventHub/namespaces/contosoHub/eventhubs/contosoHub"),
    NetworkSubnetId = new ResourceIdentifier("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Network/virtualNetworks/contoso-vnet/subnets/orbital-delegated-subnet"),
    Links = {new OrbitalContactProfileLink("contoso-uplink", OrbitalLinkPolarization.Lhcp, OrbitalLinkDirection.Uplink, new OrbitalContactProfileLinkChannel[]
    {
    new OrbitalContactProfileLinkChannel("contoso-uplink-channel", 2250, 2, new OrbitalContactEndpoint(IPAddress.Parse("10.1.0.4"), "ContosoTest_Uplink", "50000", OrbitalContactProtocol.Tcp))
    })
    {
    GainOverTemperature = 0,
    EirpdBW = 45,
    }, new OrbitalContactProfileLink("contoso-downlink", OrbitalLinkPolarization.Rhcp, OrbitalLinkDirection.Downlink, new OrbitalContactProfileLinkChannel[]
    {
    new OrbitalContactProfileLinkChannel("contoso-downlink-channel", 8160, 15, new OrbitalContactEndpoint(IPAddress.Parse("10.1.0.5"), "ContosoTest_Downlink", "50001", OrbitalContactProtocol.Udp))
    })
    {
    GainOverTemperature = 25,
    EirpdBW = 0,
    }},
};
ArmOperation<OrbitalContactProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, contactProfileName, data);
OrbitalContactProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OrbitalContactProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
