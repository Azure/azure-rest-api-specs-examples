using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/entities/timeline/PostTimelineEntity.json
// this example is just showing the usage of "EntitiesGetTimeline_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsEntityResource created on azure
// for more information of creating SecurityInsightsEntityResource, please refer to the document of SecurityInsightsEntityResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string entityId = "e1d3d618-e11f-478b-98e3-bb381539a8e1";
ResourceIdentifier securityInsightsEntityResourceId = SecurityInsightsEntityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, entityId);
SecurityInsightsEntityResource securityInsightsEntity = client.GetSecurityInsightsEntityResource(securityInsightsEntityResourceId);

// invoke the operation and iterate over the result
EntityTimelineContent content = new EntityTimelineContent(DateTimeOffset.Parse("2021-09-01T00:00:00.000Z"), DateTimeOffset.Parse("2021-10-01T00:00:00.000Z"))
{
    NumberOfBucket = 4,
};
await foreach (EntityTimelineItem item in securityInsightsEntity.GetEntitiesGetTimelinesAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
