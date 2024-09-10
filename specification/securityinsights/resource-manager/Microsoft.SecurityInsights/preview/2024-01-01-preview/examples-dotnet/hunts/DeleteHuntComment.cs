using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/hunts/DeleteHuntComment.json
// this example is just showing the usage of "HuntComments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsHuntCommentResource created on azure
// for more information of creating SecurityInsightsHuntCommentResource, please refer to the document of SecurityInsightsHuntCommentResource
string subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string huntId = "163e7b2a-a2ec-4041-aaba-d878a38f265f";
string huntCommentId = "2216d0e1-91e3-4902-89fd-d2df8c123456";
ResourceIdentifier securityInsightsHuntCommentResourceId = SecurityInsightsHuntCommentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, huntId, huntCommentId);
SecurityInsightsHuntCommentResource securityInsightsHuntComment = client.GetSecurityInsightsHuntCommentResource(securityInsightsHuntCommentResourceId);

// invoke the operation
await securityInsightsHuntComment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
