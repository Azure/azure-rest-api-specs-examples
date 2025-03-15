using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
// this example is just showing the usage of "AssessmentsMetadata_DeleteInSubscription" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionAssessmentMetadataResource created on azure
// for more information of creating SubscriptionAssessmentMetadataResource, please refer to the document of SubscriptionAssessmentMetadataResource
string subscriptionId = "0980887d-03d6-408c-9566-532f3456804e";
string assessmentMetadataName = "ca039e75-a276-4175-aebc-bcd41e4b14b7";
ResourceIdentifier subscriptionAssessmentMetadataResourceId = SubscriptionAssessmentMetadataResource.CreateResourceIdentifier(subscriptionId, assessmentMetadataName);
SubscriptionAssessmentMetadataResource subscriptionAssessmentMetadata = client.GetSubscriptionAssessmentMetadataResource(subscriptionAssessmentMetadataResourceId);

// invoke the operation
await subscriptionAssessmentMetadata.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
