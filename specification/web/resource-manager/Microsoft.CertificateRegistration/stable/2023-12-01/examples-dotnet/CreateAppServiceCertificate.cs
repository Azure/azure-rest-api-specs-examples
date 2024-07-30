using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/CreateAppServiceCertificate.json
// this example is just showing the usage of "AppServiceCertificateOrders_CreateOrUpdateCertificate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceCertificateOrderResource created on azure
// for more information of creating AppServiceCertificateOrderResource, please refer to the document of AppServiceCertificateOrderResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string certificateOrderName = "SampleCertificateOrderName";
ResourceIdentifier appServiceCertificateOrderResourceId = AppServiceCertificateOrderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, certificateOrderName);
AppServiceCertificateOrderResource appServiceCertificateOrder = client.GetAppServiceCertificateOrderResource(appServiceCertificateOrderResourceId);

// get the collection of this AppServiceCertificateResource
AppServiceCertificateCollection collection = appServiceCertificateOrder.GetAppServiceCertificates();

// invoke the operation
string name = "SampleCertName1";
AppServiceCertificateData data = new AppServiceCertificateData(new AzureLocation("Global"))
{
    KeyVaultId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
    KeyVaultSecretName = "SampleSecretName1",
};
ArmOperation<AppServiceCertificateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
AppServiceCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppServiceCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
