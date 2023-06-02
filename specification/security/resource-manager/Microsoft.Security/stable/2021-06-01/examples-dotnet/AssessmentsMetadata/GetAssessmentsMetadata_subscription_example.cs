using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_subscription_example.json
// this example is just showing the usage of "AssessmentsMetadata_GetInSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionAssessmentMetadataResource created on azure
// for more information of creating SubscriptionAssessmentMetadataResource, please refer to the document of SubscriptionAssessmentMetadataResource
string subscriptionId = "0980887d-03d6-408c-9566-532f3456804e";
string assessmentMetadataName = "21300918-b2e3-0346-785f-c77ff57d243b";
ResourceIdentifier subscriptionAssessmentMetadataResourceId = SubscriptionAssessmentMetadataResource.CreateResourceIdentifier(subscriptionId, assessmentMetadataName);
SubscriptionAssessmentMetadataResource subscriptionAssessmentMetadata = client.GetSubscriptionAssessmentMetadataResource(subscriptionAssessmentMetadataResourceId);

// invoke the operation
SubscriptionAssessmentMetadataResource result = await subscriptionAssessmentMetadata.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityAssessmentMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
