using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/deleteGuestConfigurationAssignment.json
// this example is just showing the usage of "GuestConfigurationAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVmAssignmentResource created on azure
// for more information of creating GuestConfigurationVmAssignmentResource, please refer to the document of GuestConfigurationVmAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string guestConfigurationAssignmentName = "SecureProtocol";
ResourceIdentifier guestConfigurationVmAssignmentResourceId = GuestConfigurationVmAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, guestConfigurationAssignmentName);
GuestConfigurationVmAssignmentResource guestConfigurationVmAssignment = client.GetGuestConfigurationVmAssignmentResource(guestConfigurationVmAssignmentResourceId);

// invoke the operation
await guestConfigurationVmAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
