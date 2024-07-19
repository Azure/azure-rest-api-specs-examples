using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ServerTrustCertificatesDelete.json
// this example is just showing the usage of "ServerTrustCertificates_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceServerTrustCertificateResource created on azure
// for more information of creating ManagedInstanceServerTrustCertificateResource, please refer to the document of ManagedInstanceServerTrustCertificateResource
string subscriptionId = "38e0dc56-907f-45ba-a97c-74233baad471";
string resourceGroupName = "testrg";
string managedInstanceName = "testcl";
string certificateName = "customerCertificateName";
ResourceIdentifier managedInstanceServerTrustCertificateResourceId = ManagedInstanceServerTrustCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, certificateName);
ManagedInstanceServerTrustCertificateResource managedInstanceServerTrustCertificate = client.GetManagedInstanceServerTrustCertificateResource(managedInstanceServerTrustCertificateResourceId);

// invoke the operation
await managedInstanceServerTrustCertificate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
