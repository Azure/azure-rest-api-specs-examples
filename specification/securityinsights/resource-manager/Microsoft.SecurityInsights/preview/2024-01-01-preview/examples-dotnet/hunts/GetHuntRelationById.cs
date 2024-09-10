using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/hunts/GetHuntRelationById.json
// this example is just showing the usage of "HuntRelations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsHuntResource created on azure
// for more information of creating SecurityInsightsHuntResource, please refer to the document of SecurityInsightsHuntResource
string subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string huntId = "163e7b2a-a2ec-4041-aaba-d878a38f265f";
ResourceIdentifier securityInsightsHuntResourceId = SecurityInsightsHuntResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, huntId);
SecurityInsightsHuntResource securityInsightsHunt = client.GetSecurityInsightsHuntResource(securityInsightsHuntResourceId);

// get the collection of this SecurityInsightsHuntRelationResource
SecurityInsightsHuntRelationCollection collection = securityInsightsHunt.GetSecurityInsightsHuntRelations();

// invoke the operation
string huntRelationId = "2216d0e1-91e3-4902-89fd-d2df8c535096";
NullableResponse<SecurityInsightsHuntRelationResource> response = await collection.GetIfExistsAsync(huntRelationId);
SecurityInsightsHuntRelationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SecurityInsightsHuntRelationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
