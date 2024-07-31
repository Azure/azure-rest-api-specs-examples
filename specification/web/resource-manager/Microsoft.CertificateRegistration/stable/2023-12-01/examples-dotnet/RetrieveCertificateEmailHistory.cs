using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/RetrieveCertificateEmailHistory.json
// this example is just showing the usage of "AppServiceCertificateOrders_RetrieveCertificateEmailHistory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceCertificateOrderResource created on azure
// for more information of creating AppServiceCertificateOrderResource, please refer to the document of AppServiceCertificateOrderResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "SampleCertOrder";
ResourceIdentifier appServiceCertificateOrderResourceId = AppServiceCertificateOrderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
AppServiceCertificateOrderResource appServiceCertificateOrder = client.GetAppServiceCertificateOrderResource(appServiceCertificateOrderResourceId);

// invoke the operation and iterate over the result
await foreach (AppServiceCertificateEmail item in appServiceCertificateOrder.RetrieveCertificateEmailHistoryAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
