using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.GuestConfiguration;
using Azure.ResourceManager.GuestConfiguration.Models;

// Generated from example definition: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllGuestConfigurationHCRPAssignmentReports.json
// this example is just showing the usage of "GuestConfigurationHCRPAssignmentReports_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (GuestConfigurationAssignmentReport item in guestConfigurationHcrpAssignment.GetReportsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
