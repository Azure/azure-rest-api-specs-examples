using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/entities/expand/PostExpandEntity.json
// this example is just showing the usage of "Entities_Expand" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
EntityExpandContent content = new EntityExpandContent()
{
    EndOn = DateTimeOffset.Parse("2019-05-26T00:00:00.000Z"),
    ExpansionId = Guid.Parse("a77992f3-25e9-4d01-99a4-5ff606cc410a"),
    StartOn = DateTimeOffset.Parse("2019-04-25T00:00:00.000Z"),
};
EntityExpandResult result = await securityInsightsEntity.ExpandAsync(content);

Console.WriteLine($"Succeeded: {result}");
