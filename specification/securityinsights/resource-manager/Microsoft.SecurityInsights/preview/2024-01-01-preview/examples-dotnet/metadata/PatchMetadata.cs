using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/metadata/PatchMetadata.json
// this example is just showing the usage of "Metadata_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsMetadataResource created on azure
// for more information of creating SecurityInsightsMetadataResource, please refer to the document of SecurityInsightsMetadataResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string metadataName = "metadataName";
ResourceIdentifier securityInsightsMetadataResourceId = SecurityInsightsMetadataResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, metadataName);
SecurityInsightsMetadataResource securityInsightsMetadata = client.GetSecurityInsightsMetadataResource(securityInsightsMetadataResourceId);

// invoke the operation
SecurityInsightsMetadataPatch patch = new SecurityInsightsMetadataPatch()
{
    Author = new SecurityInsightsMetadataAuthor()
    {
        Name = "User Name",
        Email = "email@microsoft.com",
    },
};
SecurityInsightsMetadataResource result = await securityInsightsMetadata.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
