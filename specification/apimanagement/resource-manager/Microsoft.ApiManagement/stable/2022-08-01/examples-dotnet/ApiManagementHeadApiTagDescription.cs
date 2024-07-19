using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiTagDescription.json
// this example is just showing the usage of "ApiTagDescription_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiTagDescriptionResource created on azure
// for more information of creating ApiTagDescriptionResource, please refer to the document of ApiTagDescriptionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "59d6bb8f1f7fab13dc67ec9b";
string tagDescriptionId = "59306a29e4bbd510dc24e5f9";
ResourceIdentifier apiTagDescriptionResourceId = ApiTagDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, tagDescriptionId);
ApiTagDescriptionResource apiTagDescription = client.GetApiTagDescriptionResource(apiTagDescriptionResourceId);

// invoke the operation
bool result = await apiTagDescription.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
