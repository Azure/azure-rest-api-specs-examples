using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Certificates_CreateOrUpdate.json
// this example is just showing the usage of "Certificates_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformCertificateResource created on azure
// for more information of creating AppPlatformCertificateResource, please refer to the document of AppPlatformCertificateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string certificateName = "mycertificate";
ResourceIdentifier appPlatformCertificateResourceId = AppPlatformCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, certificateName);
AppPlatformCertificateResource appPlatformCertificate = client.GetAppPlatformCertificateResource(appPlatformCertificateResourceId);

// invoke the operation
AppPlatformCertificateData data = new AppPlatformCertificateData()
{
    Properties = new AppPlatformKeyVaultCertificateProperties(new Uri("https://myvault.vault.azure.net"), "mycert")
    {
        CertVersion = "08a219d06d874795a96db47e06fbb01e",
    },
};
ArmOperation<AppPlatformCertificateResource> lro = await appPlatformCertificate.UpdateAsync(WaitUntil.Completed, data);
AppPlatformCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
