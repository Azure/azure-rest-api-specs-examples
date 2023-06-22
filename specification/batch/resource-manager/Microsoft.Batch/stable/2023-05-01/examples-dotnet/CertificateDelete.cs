using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/CertificateDelete.json
// this example is just showing the usage of "Certificate_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountCertificateResource created on azure
// for more information of creating BatchAccountCertificateResource, please refer to the document of BatchAccountCertificateResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string certificateName = "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e";
ResourceIdentifier batchAccountCertificateResourceId = BatchAccountCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, certificateName);
BatchAccountCertificateResource batchAccountCertificate = client.GetBatchAccountCertificateResource(batchAccountCertificateResourceId);

// invoke the operation
await batchAccountCertificate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
