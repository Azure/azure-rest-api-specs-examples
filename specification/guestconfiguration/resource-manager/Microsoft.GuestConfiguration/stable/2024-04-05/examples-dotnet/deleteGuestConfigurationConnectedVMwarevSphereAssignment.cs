using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/deleteGuestConfigurationConnectedVMwarevSphereAssignment.json
// this example is just showing the usage of "GuestConfigurationConnectedVMwarevSphereAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVMwarevSphereAssignmentResource created on azure
// for more information of creating GuestConfigurationVMwarevSphereAssignmentResource, please refer to the document of GuestConfigurationVMwarevSphereAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string guestConfigurationAssignmentName = "SecureProtocol";
ResourceIdentifier guestConfigurationVMwarevSphereAssignmentResourceId = GuestConfigurationVMwarevSphereAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, guestConfigurationAssignmentName);
GuestConfigurationVMwarevSphereAssignmentResource guestConfigurationVMwarevSphereAssignment = client.GetGuestConfigurationVMwarevSphereAssignmentResource(guestConfigurationVMwarevSphereAssignmentResourceId);

// invoke the operation
await guestConfigurationVMwarevSphereAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
