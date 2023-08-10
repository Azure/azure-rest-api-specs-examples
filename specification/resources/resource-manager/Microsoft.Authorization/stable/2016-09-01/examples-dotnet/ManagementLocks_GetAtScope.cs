using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2016-09-01/examples/ManagementLocks_GetAtScope.json
// this example is just showing the usage of "ManagementLocks_GetByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ManagementLockResource
string scope = "subscriptions/subscriptionId";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
ManagementLockCollection collection = client.GetGenericResource(scopeId).GetManagementLocks();

// invoke the operation
string lockName = "testlock";
bool result = await collection.ExistsAsync(lockName);

Console.WriteLine($"Succeeded: {result}");
