using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/ListSubAssessments_example.json
// this example is just showing the usage of "SubAssessments_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityAssessmentResource created on azure
// for more information of creating SecurityAssessmentResource, please refer to the document of SecurityAssessmentResource
string scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string assessmentName = "82e20e14-edc5-4373-bfc4-f13121257c37";
ResourceIdentifier securityAssessmentResourceId = SecurityAssessmentResource.CreateResourceIdentifier(scope, assessmentName);
SecurityAssessmentResource securityAssessment = client.GetSecurityAssessmentResource(securityAssessmentResourceId);

// get the collection of this SecuritySubAssessmentResource
SecuritySubAssessmentCollection collection = securityAssessment.GetSecuritySubAssessments();

// invoke the operation and iterate over the result
await foreach (SecuritySubAssessmentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SecuritySubAssessmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
