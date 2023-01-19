using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotHub;
using Azure.ResourceManager.IotHub.Models;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_certificatesdelete.json
// this example is just showing the usage of "Certificates_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotHubCertificateDescriptionResource created on azure
// for more information of creating IotHubCertificateDescriptionResource, please refer to the document of IotHubCertificateDescriptionResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string resourceName = "myhub";
string certificateName = "cert";
ResourceIdentifier iotHubCertificateDescriptionResourceId = IotHubCertificateDescriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, certificateName);
IotHubCertificateDescriptionResource iotHubCertificateDescription = client.GetIotHubCertificateDescriptionResource(iotHubCertificateDescriptionResourceId);

// invoke the operation
string ifMatch = "AAAAAAAADGk=";
await iotHubCertificateDescription.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
