using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/DeleteInstancePool.json
// this example is just showing the usage of "InstancePools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InstancePoolResource created on azure
// for more information of creating InstancePoolResource, please refer to the document of InstancePoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "group1";
string instancePoolName = "testIP";
ResourceIdentifier instancePoolResourceId = InstancePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instancePoolName);
InstancePoolResource instancePool = client.GetInstancePoolResource(instancePoolResourceId);

// invoke the operation
await instancePool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
