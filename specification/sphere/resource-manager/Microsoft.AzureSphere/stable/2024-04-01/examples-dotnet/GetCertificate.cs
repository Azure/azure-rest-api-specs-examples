using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sphere;
using Azure.ResourceManager.Sphere.Models;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCertificate.json
// this example is just showing the usage of "Certificates_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereCertificateResource created on azure
// for more information of creating SphereCertificateResource, please refer to the document of SphereCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string serialNumber = "default";
ResourceIdentifier sphereCertificateResourceId = SphereCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, serialNumber);
SphereCertificateResource sphereCertificate = client.GetSphereCertificateResource(sphereCertificateResourceId);

// invoke the operation
SphereCertificateResource result = await sphereCertificate.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SphereCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
