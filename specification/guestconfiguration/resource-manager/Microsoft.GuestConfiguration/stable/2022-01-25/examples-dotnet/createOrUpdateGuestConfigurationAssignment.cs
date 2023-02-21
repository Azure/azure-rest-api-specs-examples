using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.GuestConfiguration;
using Azure.ResourceManager.GuestConfiguration.Models;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/createOrUpdateGuestConfigurationAssignment.json
// this example is just showing the usage of "GuestConfigurationAssignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVmAssignmentResource created on azure
// for more information of creating GuestConfigurationVmAssignmentResource, please refer to the document of GuestConfigurationVmAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string guestConfigurationAssignmentName = "NotInstalledApplicationForWindows";
ResourceIdentifier guestConfigurationVmAssignmentResourceId = GuestConfigurationVmAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, guestConfigurationAssignmentName);
GuestConfigurationVmAssignmentResource guestConfigurationVmAssignment = client.GetGuestConfigurationVmAssignmentResource(guestConfigurationVmAssignmentResourceId);

// invoke the operation
GuestConfigurationAssignmentData data = new GuestConfigurationAssignmentData()
{
    Properties = new GuestConfigurationAssignmentProperties()
    {
        GuestConfiguration = new GuestConfigurationNavigation()
        {
            Name = "NotInstalledApplicationForWindows",
            Version = "1.*",
            ContentUri = new Uri("https://thisisfake/pacakge"),
            ContentHash = "123contenthash",
            AssignmentType = GuestConfigurationAssignmentType.ApplyAndAutoCorrect,
            ConfigurationParameters =
            {
            new GuestConfigurationParameter()
            {
            Name = "[InstalledApplication]NotInstalledApplicationResource1;Name",
            Value = "NotePad,sql",
            }
            },
        },
        Context = "Azure policy",
    },
    Name = "NotInstalledApplicationForWindows",
    Location = new AzureLocation("westcentralus"),
};
ArmOperation<GuestConfigurationVmAssignmentResource> lro = await guestConfigurationVmAssignment.UpdateAsync(WaitUntil.Completed, data);
GuestConfigurationVmAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GuestConfigurationAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
