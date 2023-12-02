using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2021-02-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
// this example is just showing the usage of "CertificateOrdersDiagnostics_GetAppServiceCertificateOrderDetectorResponse" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceCertificateOrderResource created on azure
// for more information of creating AppServiceCertificateOrderResource, please refer to the document of AppServiceCertificateOrderResource
string subscriptionId = "5700fc96-77b4-4f8d-afce-c353d8c443bd";
string resourceGroupName = "Sample-WestUSResourceGroup";
string certificateOrderName = "SampleCertificateOrderName";
ResourceIdentifier appServiceCertificateOrderResourceId = AppServiceCertificateOrderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, certificateOrderName);
AppServiceCertificateOrderResource appServiceCertificateOrder = client.GetAppServiceCertificateOrderResource(appServiceCertificateOrderResourceId);

// get the collection of this CertificateOrderDetectorResource
CertificateOrderDetectorCollection collection = appServiceCertificateOrder.GetCertificateOrderDetectors();

// invoke the operation
string detectorName = "AutoRenewStatus";
NullableResponse<CertificateOrderDetectorResource> response = await collection.GetIfExistsAsync(detectorName);
CertificateOrderDetectorResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppServiceDetectorData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
