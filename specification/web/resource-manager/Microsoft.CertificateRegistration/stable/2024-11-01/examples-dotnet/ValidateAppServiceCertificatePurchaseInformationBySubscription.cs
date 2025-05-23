using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2024-11-01/examples/ValidateAppServiceCertificatePurchaseInformationBySubscription.json
// this example is just showing the usage of "AppServiceCertificateOrders_ValidatePurchaseInformation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AppServiceCertificateOrderData data = new AppServiceCertificateOrderData(new AzureLocation("Global"))
{
    Certificates =
    {
    ["SampleCertName1"] = new AppServiceCertificateProperties
    {
    KeyVaultId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
    KeyVaultSecretName = "SampleSecretName1",
    },
    ["SampleCertName2"] = new AppServiceCertificateProperties
    {
    KeyVaultId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/microsoft.keyvault/vaults/SamplevaultName"),
    KeyVaultSecretName = "SampleSecretName2",
    }
    },
    DistinguishedName = "CN=SampleCustomDomain.com",
    ValidityInYears = 2,
    KeySize = 2048,
    ProductType = CertificateProductType.StandardDomainValidatedSsl,
    IsAutoRenew = true,
};
await subscriptionResource.ValidateAppServiceCertificateOrderPurchaseInformationAsync(data);

Console.WriteLine("Succeeded");
