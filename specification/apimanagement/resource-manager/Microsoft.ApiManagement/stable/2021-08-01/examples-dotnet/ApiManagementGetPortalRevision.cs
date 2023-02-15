using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPortalRevision.json
// this example is just showing the usage of "PortalRevision_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPortalRevisionResource created on azure
// for more information of creating ApiManagementPortalRevisionResource, please refer to the document of ApiManagementPortalRevisionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string portalRevisionId = "20201112101010";
ResourceIdentifier apiManagementPortalRevisionResourceId = ApiManagementPortalRevisionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, portalRevisionId);
ApiManagementPortalRevisionResource apiManagementPortalRevision = client.GetApiManagementPortalRevisionResource(apiManagementPortalRevisionResourceId);

// invoke the operation
ApiManagementPortalRevisionResource result = await apiManagementPortalRevision.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementPortalRevisionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
