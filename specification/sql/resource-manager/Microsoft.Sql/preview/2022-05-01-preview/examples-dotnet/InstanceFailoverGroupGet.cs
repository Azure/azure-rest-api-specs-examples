using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/InstanceFailoverGroupGet.json
// this example is just showing the usage of "InstanceFailoverGroups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InstanceFailoverGroupResource created on azure
// for more information of creating InstanceFailoverGroupResource, please refer to the document of InstanceFailoverGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
AzureLocation locationName = new AzureLocation("Japan East");
string failoverGroupName = "failover-group-test";
ResourceIdentifier instanceFailoverGroupResourceId = InstanceFailoverGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, locationName, failoverGroupName);
InstanceFailoverGroupResource instanceFailoverGroup = client.GetInstanceFailoverGroupResource(instanceFailoverGroupResourceId);

// invoke the operation
InstanceFailoverGroupResource result = await instanceFailoverGroup.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InstanceFailoverGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
