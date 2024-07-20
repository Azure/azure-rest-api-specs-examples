using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/listAllGuestConfigurationAssignmentReports.json
// this example is just showing the usage of "GuestConfigurationAssignmentReports_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVmAssignmentResource created on azure
// for more information of creating GuestConfigurationVmAssignmentResource, please refer to the document of GuestConfigurationVmAssignmentResource
string subscriptionId = "mySubscriptionid";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string guestConfigurationAssignmentName = "AuditSecureProtocol";
ResourceIdentifier guestConfigurationVmAssignmentResourceId = GuestConfigurationVmAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, guestConfigurationAssignmentName);
GuestConfigurationVmAssignmentResource guestConfigurationVmAssignment = client.GetGuestConfigurationVmAssignmentResource(guestConfigurationVmAssignmentResourceId);

// invoke the operation and iterate over the result
await foreach (GuestConfigurationAssignmentReport item in guestConfigurationVmAssignment.GetReportsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
