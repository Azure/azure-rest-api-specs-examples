using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementGetCertificate.json
// this example is just showing the usage of "Certificate_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementCertificateResource created on azure
// for more information of creating ApiManagementCertificateResource, please refer to the document of ApiManagementCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string certificateId = "templateCert1";
ResourceIdentifier apiManagementCertificateResourceId = ApiManagementCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, certificateId);
ApiManagementCertificateResource apiManagementCertificate = client.GetApiManagementCertificateResource(apiManagementCertificateResourceId);

// invoke the operation
ApiManagementCertificateResource result = await apiManagementCertificate.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiManagementCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
