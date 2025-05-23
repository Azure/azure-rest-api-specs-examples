using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListWorkspaceCertificates.json
// this example is just showing the usage of "WorkspaceCertificate_ListByWorkspace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceContractResource created on azure
// for more information of creating WorkspaceContractResource, please refer to the document of WorkspaceContractResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
ResourceIdentifier workspaceContractResourceId = WorkspaceContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId);
WorkspaceContractResource workspaceContract = client.GetWorkspaceContractResource(workspaceContractResourceId);

// get the collection of this ServiceWorkspaceCertificateResource
ServiceWorkspaceCertificateCollection collection = workspaceContract.GetServiceWorkspaceCertificates();

// invoke the operation and iterate over the result
await foreach (ServiceWorkspaceCertificateResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiManagementCertificateData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
