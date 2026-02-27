using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_CreateOrUpdateAtScope.json
// this example is just showing the usage of "ManagementLocks_CreateOrUpdateByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this ManagementLockResource
string scope = "subscriptions/subscriptionId";
ManagementLockCollection collection = client.GetGenericResource(new ResourceIdentifier(scope)).GetManagementLocks();

// invoke the operation
string lockName = "testlock";
ManagementLockData data = new ManagementLockData(ManagementLockLevel.ReadOnly);
ArmOperation<ManagementLockResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, lockName, data);
ManagementLockResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagementLockData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
