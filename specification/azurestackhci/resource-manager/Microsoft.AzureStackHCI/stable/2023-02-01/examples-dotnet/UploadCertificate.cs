using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/UploadCertificate.json
// this example is just showing the usage of "Clusters_UploadCertificate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterResource created on azure
// for more information of creating HciClusterResource, please refer to the document of HciClusterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
ResourceIdentifier hciClusterResourceId = HciClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterResource hciCluster = client.GetHciClusterResource(hciClusterResourceId);

// invoke the operation
HciClusterCertificateContent content = new HciClusterCertificateContent()
{
    Certificates =
    {
    "base64cert","base64cert"
    },
};
await hciCluster.UploadCertificateAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
