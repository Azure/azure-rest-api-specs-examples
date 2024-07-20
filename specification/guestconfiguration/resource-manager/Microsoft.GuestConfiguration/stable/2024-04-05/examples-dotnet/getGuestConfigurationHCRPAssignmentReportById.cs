using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.GuestConfiguration.Models;
using Azure.ResourceManager.GuestConfiguration;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2024-04-05/examples/getGuestConfigurationHCRPAssignmentReportById.json
// this example is just showing the usage of "GuestConfigurationHCRPAssignmentReports_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GuestConfigurationHcrpAssignmentResource created on azure
// for more information of creating GuestConfigurationHcrpAssignmentResource, please refer to the document of GuestConfigurationHcrpAssignmentResource
string subscriptionId = "mySubscriptionid";
string resourceGroupName = "myResourceGroupName";
string machineName = "myMachineName";
string guestConfigurationAssignmentName = "AuditSecureProtocol";
ResourceIdentifier guestConfigurationHcrpAssignmentResourceId = GuestConfigurationHcrpAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, guestConfigurationAssignmentName);
GuestConfigurationHcrpAssignmentResource guestConfigurationHcrpAssignment = client.GetGuestConfigurationHcrpAssignmentResource(guestConfigurationHcrpAssignmentResourceId);

// invoke the operation
string reportId = "7367cbb8-ae99-47d0-a33b-a283564d2cb1";
GuestConfigurationAssignmentReport result = await guestConfigurationHcrpAssignment.GetReportAsync(reportId);

Console.WriteLine($"Succeeded: {result}");
