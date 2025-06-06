using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiCenter;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_Get.json
// this example is just showing the usage of "ApiDefinitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterApiVersionResource created on azure
// for more information of creating ApiCenterApiVersionResource, please refer to the document of ApiCenterApiVersionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
string workspaceName = "default";
string apiName = "echo-api";
string versionName = "2023-01-01";
ResourceIdentifier apiCenterApiVersionResourceId = ApiCenterApiVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceName, apiName, versionName);
ApiCenterApiVersionResource apiCenterApiVersion = client.GetApiCenterApiVersionResource(apiCenterApiVersionResourceId);

// get the collection of this ApiCenterApiDefinitionResource
ApiCenterApiDefinitionCollection collection = apiCenterApiVersion.GetApiCenterApiDefinitions();

// invoke the operation
string definitionName = "openapi";
NullableResponse<ApiCenterApiDefinitionResource> response = await collection.GetIfExistsAsync(definitionName);
ApiCenterApiDefinitionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiCenterApiDefinitionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
