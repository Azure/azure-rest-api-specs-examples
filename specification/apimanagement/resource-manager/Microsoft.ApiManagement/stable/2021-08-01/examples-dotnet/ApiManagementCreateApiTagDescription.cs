using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiTagDescription.json
// this example is just showing the usage of "ApiTagDescription_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiTagDescriptionResource created on azure
// for more information of creating ApiTagDescriptionResource, please refer to the document of ApiTagDescriptionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5931a75ae4bbd512a88c680b";
string tagDescriptionId = "tagId1";
ResourceIdentifier apiTagDescriptionResourceId = ApiTagDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, tagDescriptionId);
ApiTagDescriptionResource apiTagDescription = client.GetApiTagDescriptionResource(apiTagDescriptionResourceId);

// invoke the operation
ApiTagDescriptionCreateOrUpdateContent content = new ApiTagDescriptionCreateOrUpdateContent()
{
    Description = "Some description that will be displayed for operation's tag if the tag is assigned to operation of the API",
    ExternalDocsUri = new Uri("http://some.url/additionaldoc"),
    ExternalDocsDescription = "Description of the external docs resource",
};
ArmOperation<ApiTagDescriptionResource> lro = await apiTagDescription.UpdateAsync(WaitUntil.Completed, content);
ApiTagDescriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiTagDescriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
