using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Nginx;

// Generated from example definition: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-06-01-preview/examples/Certificates_Get.json
// this example is just showing the usage of "Certificates_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NginxCertificateResource created on azure
// for more information of creating NginxCertificateResource, please refer to the document of NginxCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string deploymentName = "myDeployment";
string certificateName = "default";
ResourceIdentifier nginxCertificateResourceId = NginxCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deploymentName, certificateName);
NginxCertificateResource nginxCertificate = client.GetNginxCertificateResource(nginxCertificateResourceId);

// invoke the operation
NginxCertificateResource result = await nginxCertificate.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NginxCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
