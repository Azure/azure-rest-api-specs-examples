using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/getGuestConfigurationConnectedVMwarevSphereAssignment.json
// this example is just showing the usage of "GuestConfigurationConnectedVMwarevSphereAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this GuestConfigurationVMwarevSphereAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string vmName = "myVMName";
string scope = $"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualmachines/{vmName}";
GuestConfigurationVMwarevSphereAssignmentCollection collection = client.GetGuestConfigurationVMwarevSphereAssignments(new ResourceIdentifier(scope));

// invoke the operation
string guestConfigurationAssignmentName = "SecureProtocol";
NullableResponse<GuestConfigurationVMwarevSphereAssignmentResource> response = await collection.GetIfExistsAsync(guestConfigurationAssignmentName);
GuestConfigurationVMwarevSphereAssignmentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    GuestConfigurationAssignmentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
