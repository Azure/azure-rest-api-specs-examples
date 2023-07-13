using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sphere;
using Azure.ResourceManager.Sphere.Models;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostRetrieveCatalogCertChain.json
// this example is just showing the usage of "Certificates_RetrieveCertChain" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereCertificateResource created on azure
// for more information of creating SphereCertificateResource, please refer to the document of SphereCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string serialNumber = "active";
ResourceIdentifier sphereCertificateResourceId = SphereCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, serialNumber);
SphereCertificateResource sphereCertificate = client.GetSphereCertificateResource(sphereCertificateResourceId);

// invoke the operation
SphereCertificateChainResult result = await sphereCertificate.RetrieveCertChainAsync();

Console.WriteLine($"Succeeded: {result}");
