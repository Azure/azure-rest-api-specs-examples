using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/EncryptionProtectorCreateOrUpdateKeyVault.json
// this example is just showing the usage of "EncryptionProtectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EncryptionProtectorResource created on azure
// for more information of creating EncryptionProtectorResource, please refer to the document of EncryptionProtectorResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string serverName = "sqlcrudtest-4645";
EncryptionProtectorName encryptionProtectorName = EncryptionProtectorName.Current;
ResourceIdentifier encryptionProtectorResourceId = EncryptionProtectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, encryptionProtectorName);
EncryptionProtectorResource encryptionProtector = client.GetEncryptionProtectorResource(encryptionProtectorResourceId);

// invoke the operation
EncryptionProtectorData data = new EncryptionProtectorData
{
    ServerKeyName = "someVault_someKey_01234567890123456789012345678901",
    ServerKeyType = SqlServerKeyType.AzureKeyVault,
    IsAutoRotationEnabled = false,
};
ArmOperation<EncryptionProtectorResource> lro = await encryptionProtector.UpdateAsync(WaitUntil.Completed, data);
EncryptionProtectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EncryptionProtectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
