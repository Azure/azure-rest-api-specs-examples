using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountGetObjectReplicationPolicy.json
// this example is just showing the usage of "ObjectReplicationPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ObjectReplicationPolicyResource created on azure
// for more information of creating ObjectReplicationPolicyResource, please refer to the document of ObjectReplicationPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string accountName = "sto2527";
string objectReplicationPolicyId = "{objectReplicationPolicy-Id}";
ResourceIdentifier objectReplicationPolicyResourceId = ObjectReplicationPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, objectReplicationPolicyId);
ObjectReplicationPolicyResource objectReplicationPolicy = client.GetObjectReplicationPolicyResource(objectReplicationPolicyResourceId);

// invoke the operation
ObjectReplicationPolicyResource result = await objectReplicationPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ObjectReplicationPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
