using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.GuestConfiguration;
using Azure.ResourceManager.GuestConfiguration.Models;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationAssignmentReportById.json
// this example is just showing the usage of "GuestConfigurationAssignmentReports_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationVmAssignmentResource created on azure
// for more information of creating GuestConfigurationVmAssignmentResource, please refer to the document of GuestConfigurationVmAssignmentResource
string subscriptionId = "mySubscriptionid";
string resourceGroupName = "myResourceGroupName";
string vmName = "myvm";
string guestConfigurationAssignmentName = "AuditSecureProtocol";
ResourceIdentifier guestConfigurationVmAssignmentResourceId = GuestConfigurationVmAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, guestConfigurationAssignmentName);
GuestConfigurationVmAssignmentResource guestConfigurationVmAssignment = client.GetGuestConfigurationVmAssignmentResource(guestConfigurationVmAssignmentResourceId);

// invoke the operation
string reportId = "7367cbb8-ae99-47d0-a33b-a283564d2cb1";
GuestConfigurationAssignmentReport result = await guestConfigurationVmAssignment.GetReportAsync(reportId);

Console.WriteLine($"Succeeded: {result}");
