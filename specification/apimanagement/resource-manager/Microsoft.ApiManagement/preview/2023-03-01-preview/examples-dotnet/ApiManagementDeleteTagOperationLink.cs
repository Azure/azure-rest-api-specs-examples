using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementDeleteTagOperationLink.json
// this example is just showing the usage of "TagOperationLink_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceTagOperationLinkResource created on azure
// for more information of creating ServiceTagOperationLinkResource, please refer to the document of ServiceTagOperationLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string tagId = "tag1";
string operationLinkId = "link1";
ResourceIdentifier serviceTagOperationLinkResourceId = ServiceTagOperationLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, tagId, operationLinkId);
ServiceTagOperationLinkResource serviceTagOperationLink = client.GetServiceTagOperationLinkResource(serviceTagOperationLinkResourceId);

// invoke the operation
await serviceTagOperationLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
