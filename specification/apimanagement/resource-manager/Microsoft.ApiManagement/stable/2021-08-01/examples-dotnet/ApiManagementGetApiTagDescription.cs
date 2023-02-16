using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiTagDescription.json
// this example is just showing the usage of "ApiTagDescription_Get" operation, for the dependent resources, they will have to be created separately.

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
ApiTagDescriptionResource result = await apiTagDescription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiTagDescriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
