using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ManagedInstanceKeyDelete.json
// this example is just showing the usage of "ManagedInstanceKeys_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceKeyResource created on azure
// for more information of creating ManagedInstanceKeyResource, please refer to the document of ManagedInstanceKeyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string managedInstanceName = "sqlcrudtest-4645";
string keyName = "someVault_someKey_01234567890123456789012345678901";
ResourceIdentifier managedInstanceKeyResourceId = ManagedInstanceKeyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, keyName);
ManagedInstanceKeyResource managedInstanceKey = client.GetManagedInstanceKeyResource(managedInstanceKeyResourceId);

// invoke the operation
await managedInstanceKey.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
