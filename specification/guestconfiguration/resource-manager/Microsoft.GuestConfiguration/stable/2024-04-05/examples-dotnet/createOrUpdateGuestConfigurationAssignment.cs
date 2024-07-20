using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/createOrUpdateGuestConfigurationAssignment.json
// this example is just showing the usage of "GuestConfigurationAssignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this GuestConfigurationVmAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/subscriptions/{0}/resourceGroups/{1}/providers/Microsoft.Compute/virtualMachines/{2}", subscriptionId, resourceGroupName, vmName));
GuestConfigurationVmAssignmentCollection collection = client.GetGuestConfigurationVmAssignments(scopeId);

// invoke the operation
string guestConfigurationAssignmentName = "NotInstalledApplicationForWindows";
GuestConfigurationAssignmentData data = new GuestConfigurationAssignmentData()
{
    Properties = new GuestConfigurationAssignmentProperties()
    {
        GuestConfiguration = new GuestConfigurationNavigation()
        {
            Name = "NotInstalledApplicationForWindows",
            Version = "1.0.0.3",
            ContentUri = new Uri("https://thisisfake/pacakge"),
            ContentHash = "123contenthash",
            ContentManagedIdentity = "test_identity",
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
ArmOperation<GuestConfigurationVmAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, guestConfigurationAssignmentName, data);
GuestConfigurationVmAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GuestConfigurationAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
