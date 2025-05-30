using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CustomerInsights.Models;
using Azure.ResourceManager.CustomerInsights;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsDelete.json
// this example is just showing the usage of "RoleAssignments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleAssignmentResourceFormatResource created on azure
// for more information of creating RoleAssignmentResourceFormatResource, please refer to the document of RoleAssignmentResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string assignmentName = "assignmentName8976";
ResourceIdentifier roleAssignmentResourceFormatResourceId = RoleAssignmentResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, assignmentName);
RoleAssignmentResourceFormatResource roleAssignmentResourceFormat = client.GetRoleAssignmentResourceFormatResource(roleAssignmentResourceFormatResourceId);

// invoke the operation
await roleAssignmentResourceFormat.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
