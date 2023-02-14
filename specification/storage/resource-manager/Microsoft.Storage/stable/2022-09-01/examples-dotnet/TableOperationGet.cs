using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Storage;
using Azure.ResourceManager.Storage.Models;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/TableOperationGet.json
// this example is just showing the usage of "Table_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(tableName);

Console.WriteLine($"Succeeded: {result}");
