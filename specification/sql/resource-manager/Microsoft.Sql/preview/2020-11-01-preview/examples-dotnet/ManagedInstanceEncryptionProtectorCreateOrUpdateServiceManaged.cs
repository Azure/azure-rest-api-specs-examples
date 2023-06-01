using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceEncryptionProtectorCreateOrUpdateServiceManaged.json
// this example is just showing the usage of "ManagedInstanceEncryptionProtectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string managedInstanceName = "sqlcrudtest-4645";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// get the collection of this ManagedInstanceEncryptionProtectorResource
ManagedInstanceEncryptionProtectorCollection collection = managedInstance.GetManagedInstanceEncryptionProtectors();

// invoke the operation
EncryptionProtectorName encryptionProtectorName = EncryptionProtectorName.Current;
ManagedInstanceEncryptionProtectorData data = new ManagedInstanceEncryptionProtectorData()
{
    ServerKeyName = "ServiceManaged",
    ServerKeyType = SqlServerKeyType.ServiceManaged,
};
ArmOperation<ManagedInstanceEncryptionProtectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, encryptionProtectorName, data);
ManagedInstanceEncryptionProtectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceEncryptionProtectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
