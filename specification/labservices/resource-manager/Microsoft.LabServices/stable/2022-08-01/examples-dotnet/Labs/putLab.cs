using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LabServices;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/putLab.json
// this example is just showing the usage of "Labs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this LabResource
LabCollection collection = resourceGroupResource.GetLabs();

// invoke the operation
string labName = "testlab";
LabData data = new LabData(new AzureLocation("westus"))
{
    AutoShutdownProfile = new LabAutoShutdownProfile()
    {
        ShutdownOnDisconnect = LabServicesEnableState.Enabled,
        ShutdownWhenNotConnected = LabServicesEnableState.Enabled,
        ShutdownOnIdle = LabVirtualMachineShutdownOnIdleMode.UserAbsence,
        DisconnectDelay = XmlConvert.ToTimeSpan("PT5M"),
        NoConnectDelay = XmlConvert.ToTimeSpan("PT5M"),
        IdleDelay = XmlConvert.ToTimeSpan("PT5M"),
    },
    ConnectionProfile = new LabConnectionProfile()
    {
        WebSshAccess = LabVirtualMachineConnectionType.None,
        WebRdpAccess = LabVirtualMachineConnectionType.None,
        ClientSshAccess = LabVirtualMachineConnectionType.Public,
        ClientRdpAccess = LabVirtualMachineConnectionType.Public,
    },
    VirtualMachineProfile = new LabVirtualMachineProfile(LabVirtualMachineCreateOption.TemplateVm, new LabVirtualMachineImageReference()
    {
        Offer = "WindowsServer",
        Publisher = "Microsoft",
        Sku = "2019-Datacenter",
        Version = "2019.0.20190410",
    }, new LabServicesSku("Medium"), XmlConvert.ToTimeSpan("PT10H"), new LabVirtualMachineCredential("test-user"))
    {
        AdditionalCapabilitiesInstallGpuDrivers = LabServicesEnableState.Disabled,
        UseSharedPassword = LabServicesEnableState.Disabled,
    },
    SecurityProfile = new LabSecurityProfile()
    {
        OpenAccess = LabServicesEnableState.Disabled,
    },
    LabPlanId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan"),
    Title = "Test Lab",
    Description = "This is a test lab.",
    NetworkProfile = new LabNetworkProfile()
    {
        SubnetId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
    },
};
ArmOperation<LabResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, labName, data);
LabResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LabData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
