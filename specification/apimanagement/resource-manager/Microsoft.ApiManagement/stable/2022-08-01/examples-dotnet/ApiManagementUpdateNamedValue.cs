using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateNamedValue.json
// this example is just showing the usage of "NamedValue_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementNamedValueResource created on azure
// for more information of creating ApiManagementNamedValueResource, please refer to the document of ApiManagementNamedValueResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string namedValueId = "testprop2";
ResourceIdentifier apiManagementNamedValueResourceId = ApiManagementNamedValueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, namedValueId);
ApiManagementNamedValueResource apiManagementNamedValue = client.GetApiManagementNamedValueResource(apiManagementNamedValueResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
ApiManagementNamedValuePatch patch = new ApiManagementNamedValuePatch()
{
    Tags =
    {
    "foo","bar2"
    },
    IsSecret = false,
    DisplayName = "prop3name",
    Value = "propValue",
};
ArmOperation<ApiManagementNamedValueResource> lro = await apiManagementNamedValue.UpdateAsync(WaitUntil.Completed, ifMatch, patch);
ApiManagementNamedValueResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementNamedValueData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
