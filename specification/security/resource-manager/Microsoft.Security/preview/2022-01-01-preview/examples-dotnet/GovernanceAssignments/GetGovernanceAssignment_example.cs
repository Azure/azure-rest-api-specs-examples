using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceAssignments/GetGovernanceAssignment_example.json
// this example is just showing the usage of "GovernanceAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAssessmentResource created on azure
// for more information of creating SecurityAssessmentResource, please refer to the document of SecurityAssessmentResource
string scope = "subscriptions/c32e05d9-7207-4e22-bdf4-4f7d9c72e5fd/resourceGroups/compute_servers/providers/Microsoft.Compute/virtualMachines/win2012";
string assessmentName = "6b9421dd-5555-2251-9b3d-2be58e2f82cd";
ResourceIdentifier securityAssessmentResourceId = SecurityAssessmentResource.CreateResourceIdentifier(scope, assessmentName);
SecurityAssessmentResource securityAssessment = client.GetSecurityAssessmentResource(securityAssessmentResourceId);

// get the collection of this GovernanceAssignmentResource
GovernanceAssignmentCollection collection = securityAssessment.GetGovernanceAssignments();

// invoke the operation
string assignmentKey = "6634ff9f-127b-4bf2-8e6e-b1737f5e789c";
bool result = await collection.ExistsAsync(assignmentKey);

Console.WriteLine($"Succeeded: {result}");
