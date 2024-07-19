using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/getVMSSGuestConfigurationAssignmentReportById.json
// this example is just showing the usage of "GuestConfigurationAssignmentReportsVMSS_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVmssAssignmentResource created on azure
// for more information of creating GuestConfigurationVmssAssignmentResource, please refer to the document of GuestConfigurationVmssAssignmentResource
string subscriptionId = "mySubscriptionid";
string resourceGroupName = "myResourceGroupName";
string vmssName = "myvmss";
string name = "AuditSecureProtocol";
ResourceIdentifier guestConfigurationVmssAssignmentResourceId = GuestConfigurationVmssAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmssName, name);
GuestConfigurationVmssAssignmentResource guestConfigurationVmssAssignment = client.GetGuestConfigurationVmssAssignmentResource(guestConfigurationVmssAssignmentResourceId);

// invoke the operation
string id = "7367cbb8-ae99-47d0-a33b-a283564d2cb1";
GuestConfigurationAssignmentReport result = await guestConfigurationVmssAssignment.GetReportAsync(id);

Console.WriteLine($"Succeeded: {result}");
