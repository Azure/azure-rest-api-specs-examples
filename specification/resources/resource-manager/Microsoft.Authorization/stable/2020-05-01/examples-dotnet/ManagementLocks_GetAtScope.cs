using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_GetAtScope.json
// this example is just showing the usage of "ManagementLocks_GetByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementLockResource created on azure
// for more information of creating ManagementLockResource, please refer to the document of ManagementLockResource
string scope = "subscriptions/subscriptionId";
string lockName = "testlock";
ResourceIdentifier managementLockResourceId = ManagementLockResource.CreateResourceIdentifier(scope, lockName);
ManagementLockResource managementLock = client.GetManagementLockResource(managementLockResourceId);

// invoke the operation
ManagementLockResource result = await managementLock.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagementLockData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
