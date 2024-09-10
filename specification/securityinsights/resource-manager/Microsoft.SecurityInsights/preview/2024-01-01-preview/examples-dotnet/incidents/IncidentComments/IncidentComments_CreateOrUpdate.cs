using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/incidents/IncidentComments/IncidentComments_CreateOrUpdate.json
// this example is just showing the usage of "IncidentComments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsIncidentCommentResource created on azure
// for more information of creating SecurityInsightsIncidentCommentResource, please refer to the document of SecurityInsightsIncidentCommentResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string incidentId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
string incidentCommentId = "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014";
ResourceIdentifier securityInsightsIncidentCommentResourceId = SecurityInsightsIncidentCommentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, incidentId, incidentCommentId);
SecurityInsightsIncidentCommentResource securityInsightsIncidentComment = client.GetSecurityInsightsIncidentCommentResource(securityInsightsIncidentCommentResourceId);

// invoke the operation
SecurityInsightsIncidentCommentData data = new SecurityInsightsIncidentCommentData()
{
    Message = "Some message",
};
ArmOperation<SecurityInsightsIncidentCommentResource> lro = await securityInsightsIncidentComment.UpdateAsync(WaitUntil.Completed, data);
SecurityInsightsIncidentCommentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsIncidentCommentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
