using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/UpdateAppServiceCertificate.json
// this example is just showing the usage of "AppServiceCertificateOrders_UpdateCertificate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceCertificateResource created on azure
// for more information of creating AppServiceCertificateResource, please refer to the document of AppServiceCertificateResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string certificateOrderName = "SampleCertificateOrderName";
string name = "SampleCertName1";
ResourceIdentifier appServiceCertificateResourceId = AppServiceCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, certificateOrderName, name);
AppServiceCertificateResource appServiceCertificate = client.GetAppServiceCertificateResource(appServiceCertificateResourceId);

// invoke the operation
AppServiceCertificatePatch patch = new AppServiceCertificatePatch()
{
    KeyVaultId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
    KeyVaultSecretName = "SampleSecretName1",
};
AppServiceCertificateResource result = await appServiceCertificate.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppServiceCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
