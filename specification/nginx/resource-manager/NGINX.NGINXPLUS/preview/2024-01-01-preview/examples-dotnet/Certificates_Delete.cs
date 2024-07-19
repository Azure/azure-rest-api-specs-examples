using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Nginx.Models;
using Azure.ResourceManager.Nginx;

// Generated from example definition: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Certificates_Delete.json
// this example is just showing the usage of "Certificates_Delete" operation, for the dependent resources, they will have to be created separately.

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
await nginxCertificate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
