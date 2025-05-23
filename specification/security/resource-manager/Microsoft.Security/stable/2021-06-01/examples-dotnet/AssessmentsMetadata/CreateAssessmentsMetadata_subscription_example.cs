using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
// this example is just showing the usage of "AssessmentsMetadata_CreateInSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "0980887d-03d6-408c-9566-532f3456804e";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SubscriptionAssessmentMetadataResource
SubscriptionAssessmentMetadataCollection collection = subscriptionResource.GetAllSubscriptionAssessmentMetadata();

// invoke the operation
string assessmentMetadataName = "ca039e75-a276-4175-aebc-bcd41e4b14b7";
SecurityAssessmentMetadataData data = new SecurityAssessmentMetadataData
{
    DisplayName = "Install endpoint protection solution on virtual machine scale sets",
    Description = "Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities.",
    RemediationDescription = "To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>",
    Categories = { SecurityAssessmentResourceCategory.Compute },
    Severity = SecurityAssessmentSeverity.Medium,
    UserImpact = SecurityAssessmentUserImpact.Low,
    ImplementationEffort = ImplementationEffort.Low,
    Threats = { SecurityThreat.DataExfiltration, SecurityThreat.DataSpillage, SecurityThreat.MaliciousInsider },
    AssessmentType = SecurityAssessmentType.CustomerManaged,
};
ArmOperation<SubscriptionAssessmentMetadataResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, assessmentMetadataName, data);
SubscriptionAssessmentMetadataResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAssessmentMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
