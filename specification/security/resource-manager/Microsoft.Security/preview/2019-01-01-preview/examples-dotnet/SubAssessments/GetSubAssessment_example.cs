using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/GetSubAssessment_example.json
// this example is just showing the usage of "SubAssessments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAssessmentResource created on azure
// for more information of creating SecurityAssessmentResource, please refer to the document of SecurityAssessmentResource
string scope = "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/DEMORG/providers/Microsoft.Compute/virtualMachines/vm2";
string assessmentName = "1195afff-c881-495e-9bc5-1486211ae03f";
ResourceIdentifier securityAssessmentResourceId = SecurityAssessmentResource.CreateResourceIdentifier(scope, assessmentName);
SecurityAssessmentResource securityAssessment = client.GetSecurityAssessmentResource(securityAssessmentResourceId);

// get the collection of this SecuritySubAssessmentResource
SecuritySubAssessmentCollection collection = securityAssessment.GetSecuritySubAssessments();

// invoke the operation
string subAssessmentName = "95f7da9c-a2a4-1322-0758-fcd24ef09b85";
bool result = await collection.ExistsAsync(subAssessmentName);

Console.WriteLine($"Succeeded: {result}");
