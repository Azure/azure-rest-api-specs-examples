using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/DataSets_SynapseWorkspaceSqlPoolTable_Create.json
// this example is just showing the usage of "DataSets_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareDataSetResource created on azure
// for more information of creating ShareDataSetResource, please refer to the document of ShareDataSetResource
string subscriptionId = "0f3dcfc3-18f8-4099-b381-8353e19d43a7";
string resourceGroupName = "SampleResourceGroup";
string accountName = "sourceAccount";
string shareName = "share1";
string dataSetName = "dataset1";
ResourceIdentifier shareDataSetResourceId = ShareDataSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareName, dataSetName);
ShareDataSetResource shareDataSet = client.GetShareDataSetResource(shareDataSetResourceId);

// invoke the operation
ShareDataSetData data = new SynapseWorkspaceSqlPoolTableDataSet(new ResourceIdentifier("/subscriptions/0f3dcfc3-18f8-4099-b381-8353e19d43a7/resourceGroups/SampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/sqlPools/ExampleSqlPool/schemas/dbo/tables/table1"));
ArmOperation<ShareDataSetResource> lro = await shareDataSet.UpdateAsync(WaitUntil.Completed, data);
ShareDataSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ShareDataSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
