using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.LabServices;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/putLabPlan.json
// this example is just showing the usage of "LabPlans_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this LabPlanResource
LabPlanCollection collection = resourceGroupResource.GetLabPlans();

// invoke the operation
string labPlanName = "testlabplan";
LabPlanData data = new LabPlanData(new AzureLocation("westus"))
{
    DefaultConnectionProfile = new LabConnectionProfile
    {
        WebSshAccess = LabVirtualMachineConnectionType.None,
        WebRdpAccess = LabVirtualMachineConnectionType.None,
        ClientSshAccess = LabVirtualMachineConnectionType.Public,
        ClientRdpAccess = LabVirtualMachineConnectionType.Public,
    },
    DefaultAutoShutdownProfile = new LabAutoShutdownProfile
    {
        ShutdownOnDisconnect = LabServicesEnableState.Enabled,
        ShutdownWhenNotConnected = LabServicesEnableState.Enabled,
        ShutdownOnIdle = LabVirtualMachineShutdownOnIdleMode.UserAbsence,
        DisconnectDelay = XmlConvert.ToTimeSpan("PT5M"),
        NoConnectDelay = XmlConvert.ToTimeSpan("PT5M"),
        IdleDelay = XmlConvert.ToTimeSpan("PT5M"),
    },
    DefaultNetworkSubnetId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
    SharedGalleryId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig"),
    SupportInfo = new LabPlanSupportInfo
    {
        Uri = new Uri("help.contoso.com"),
        Email = "help@contoso.com",
        Phone = "+1-202-555-0123",
        Instructions = "Contact support for help.",
    },
};
ArmOperation<LabPlanResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, labPlanName, data);
LabPlanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LabPlanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
