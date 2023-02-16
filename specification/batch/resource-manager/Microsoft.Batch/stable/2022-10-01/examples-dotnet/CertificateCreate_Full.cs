using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/CertificateCreate_Full.json
// this example is just showing the usage of "Certificate_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// get the collection of this BatchAccountCertificateResource
BatchAccountCertificateCollection collection = batchAccount.GetBatchAccountCertificates();

// invoke the operation
string certificateName = "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e";
BatchAccountCertificateCreateOrUpdateContent content = new BatchAccountCertificateCreateOrUpdateContent()
{
    ThumbprintAlgorithm = "sha1",
    Thumbprint = BinaryData.FromString("0a0e4f50d51beadeac1d35afc5116098e7902e6e"),
    Format = BatchAccountCertificateFormat.Pfx,
    Data = BinaryData.FromString("MIIJsgIBAzCCCW4GCSqGSIb3DQE..."),
    Password = "<ExamplePassword>",
};
ArmOperation<BatchAccountCertificateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, certificateName, content);
BatchAccountCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
