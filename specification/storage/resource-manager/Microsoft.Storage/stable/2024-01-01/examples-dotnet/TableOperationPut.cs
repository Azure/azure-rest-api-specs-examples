using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/TableOperationPut.json
// this example is just showing the usage of "Table_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TableServiceResource created on azure
// for more information of creating TableServiceResource, please refer to the document of TableServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
ResourceIdentifier tableServiceResourceId = TableServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
TableServiceResource tableService = client.GetTableServiceResource(tableServiceResourceId);

// get the collection of this TableResource
TableCollection collection = tableService.GetTables();

// invoke the operation
string tableName = "table6185";
TableData data = new TableData();
ArmOperation<TableResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, tableName, data);
TableResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
