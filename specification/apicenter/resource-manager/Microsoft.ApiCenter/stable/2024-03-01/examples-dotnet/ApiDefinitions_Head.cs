using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiCenter.Models;
using Azure.ResourceManager.ApiCenter;

// Generated from example definition: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_Head.json
// this example is just showing the usage of "ApiDefinitions_Head" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiCenterApiDefinitionResource created on azure
// for more information of creating ApiCenterApiDefinitionResource, please refer to the document of ApiCenterApiDefinitionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso-resources";
string serviceName = "contoso";
string workspaceName = "default";
string apiName = "echo-api";
string versionName = "2023-01-01";
string definitionName = "openapi";
ResourceIdentifier apiCenterApiDefinitionResourceId = ApiCenterApiDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceName, apiName, versionName, definitionName);
ApiCenterApiDefinitionResource apiCenterApiDefinition = client.GetApiCenterApiDefinitionResource(apiCenterApiDefinitionResourceId);

// invoke the operation
bool result = await apiCenterApiDefinition.HeadAsync();

Console.WriteLine($"Succeeded: {result}");
