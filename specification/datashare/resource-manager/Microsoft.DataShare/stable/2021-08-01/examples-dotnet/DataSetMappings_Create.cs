using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/DataSetMappings_Create.json
// this example is just showing the usage of "DataSetMappings_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareSubscriptionResource created on azure
// for more information of creating ShareSubscriptionResource, please refer to the document of ShareSubscriptionResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareSubscriptionName = "ShareSubscription1";
ResourceIdentifier shareSubscriptionResourceId = ShareSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareSubscriptionName);
ShareSubscriptionResource shareSubscription = client.GetShareSubscriptionResource(shareSubscriptionResourceId);

// get the collection of this ShareDataSetMappingResource
ShareDataSetMappingCollection collection = shareSubscription.GetShareDataSetMappings();

// invoke the operation
string dataSetMappingName = "DatasetMapping1";
ShareDataSetMappingData data = new BlobDataSetMapping("C1", Guid.Parse("a08f184b-0567-4b11-ba22-a1199336d226"), "file21", "SampleResourceGroup", "storage2", "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a");
ArmOperation<ShareDataSetMappingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataSetMappingName, data);
ShareDataSetMappingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ShareDataSetMappingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
