using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerTrustCertificatesCreate.json
// this example is just showing the usage of "ServerTrustCertificates_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceServerTrustCertificateResource created on azure
// for more information of creating ManagedInstanceServerTrustCertificateResource, please refer to the document of ManagedInstanceServerTrustCertificateResource
string subscriptionId = "0574222d-5c7f-489c-a172-b3013eafab53";
string resourceGroupName = "testrg";
string managedInstanceName = "testcl";
string certificateName = "customerCertificateName";
ResourceIdentifier managedInstanceServerTrustCertificateResourceId = ManagedInstanceServerTrustCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, certificateName);
ManagedInstanceServerTrustCertificateResource managedInstanceServerTrustCertificate = client.GetManagedInstanceServerTrustCertificateResource(managedInstanceServerTrustCertificateResourceId);

// invoke the operation
ServerTrustCertificateData data = new ServerTrustCertificateData()
{
    PublicBlob = "308203AE30820296A0030201020210",
};
ArmOperation<ManagedInstanceServerTrustCertificateResource> lro = await managedInstanceServerTrustCertificate.UpdateAsync(WaitUntil.Completed, data);
ManagedInstanceServerTrustCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServerTrustCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
