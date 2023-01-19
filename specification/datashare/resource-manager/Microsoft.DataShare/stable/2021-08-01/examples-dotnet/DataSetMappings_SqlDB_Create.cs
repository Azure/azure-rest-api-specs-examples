using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/DataSetMappings_SqlDB_Create.json
// this example is just showing the usage of "DataSetMappings_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareDataSetMappingResource created on azure
// for more information of creating ShareDataSetMappingResource, please refer to the document of ShareDataSetMappingResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareSubscriptionName = "ShareSubscription1";
string dataSetMappingName = "DatasetMapping1";
ResourceIdentifier shareDataSetMappingResourceId = ShareDataSetMappingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName);
ShareDataSetMappingResource shareDataSetMapping = client.GetShareDataSetMappingResource(shareDataSetMappingResourceId);

// invoke the operation
ShareDataSetMappingData data = new SqlDBTableDataSetMapping("Database1", Guid.Parse("a08f184b-0567-4b11-ba22-a1199336d226"), "dbo", new ResourceIdentifier("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.Sql/servers/Server1"), "Table1");
ArmOperation<ShareDataSetMappingResource> lro = await shareDataSetMapping.UpdateAsync(WaitUntil.Completed, data);
ShareDataSetMappingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ShareDataSetMappingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
