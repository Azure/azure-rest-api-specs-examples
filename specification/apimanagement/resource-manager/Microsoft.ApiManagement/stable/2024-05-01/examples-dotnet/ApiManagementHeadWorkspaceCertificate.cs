using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementHeadWorkspaceCertificate.json
// this example is just showing the usage of "WorkspaceCertificate_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceCertificateResource created on azure
// for more information of creating ServiceWorkspaceCertificateResource, please refer to the document of ServiceWorkspaceCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string certificateId = "templateCert1";
ResourceIdentifier serviceWorkspaceCertificateResourceId = ServiceWorkspaceCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, certificateId);
ServiceWorkspaceCertificateResource serviceWorkspaceCertificate = client.GetServiceWorkspaceCertificateResource(serviceWorkspaceCertificateResourceId);

// invoke the operation
bool result = await serviceWorkspaceCertificate.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
