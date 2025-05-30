using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ManagedInstancePrivateLinkResourcesGet.json
// this example is just showing the usage of "ManagedInstancePrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstancePrivateLinkResource created on azure
// for more information of creating ManagedInstancePrivateLinkResource, please refer to the document of ManagedInstancePrivateLinkResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string managedInstanceName = "test-cl";
string groupName = "plr";
ResourceIdentifier managedInstancePrivateLinkResourceId = ManagedInstancePrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, groupName);
ManagedInstancePrivateLinkResource managedInstancePrivateLink = client.GetManagedInstancePrivateLinkResource(managedInstancePrivateLinkResourceId);

// invoke the operation
ManagedInstancePrivateLinkResource result = await managedInstancePrivateLink.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstancePrivateLinkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
