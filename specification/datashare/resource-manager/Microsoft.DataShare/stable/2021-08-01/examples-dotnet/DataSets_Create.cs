using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/DataSets_Create.json
// this example is just showing the usage of "DataSets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareDataSetResource created on azure
// for more information of creating ShareDataSetResource, please refer to the document of ShareDataSetResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareName = "Share1";
string dataSetName = "Dataset1";
ResourceIdentifier shareDataSetResourceId = ShareDataSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName, dataSetName);
ShareDataSetResource shareDataSet = client.GetShareDataSetResource(shareDataSetResourceId);

// invoke the operation
ShareDataSetData data = new BlobDataSet("C1", "file21", "SampleResourceGroup", "storage2", "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a");
ArmOperation<ShareDataSetResource> lro = await shareDataSet.UpdateAsync(WaitUntil.Completed, data);
ShareDataSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ShareDataSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
