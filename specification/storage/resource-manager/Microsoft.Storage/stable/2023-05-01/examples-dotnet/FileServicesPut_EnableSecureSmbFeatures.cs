using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
// this example is just showing the usage of "FileServices_SetServiceProperties" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FileServiceResource created on azure
// for more information of creating FileServiceResource, please refer to the document of FileServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res4410";
string accountName = "sto8607";
ResourceIdentifier fileServiceResourceId = FileServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
FileServiceResource fileService = client.GetFileServiceResource(fileServiceResourceId);

// invoke the operation
FileServiceData data = new FileServiceData()
{
    ProtocolSmbSetting = new SmbSetting()
    {
        Versions = "SMB2.1;SMB3.0;SMB3.1.1",
        AuthenticationMethods = "NTLMv2;Kerberos",
        KerberosTicketEncryption = "RC4-HMAC;AES-256",
        ChannelEncryption = "AES-128-CCM;AES-128-GCM;AES-256-GCM",
    },
};
ArmOperation<FileServiceResource> lro = await fileService.CreateOrUpdateAsync(WaitUntil.Completed, data);
FileServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FileServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
