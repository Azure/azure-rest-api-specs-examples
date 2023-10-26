using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsListByDatabase.json
// this example is just showing the usage of "Scripts_ListByDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDatabaseResource created on azure
// for more information of creating KustoDatabaseResource, please refer to the document of KustoDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "Kustodatabase8";
ResourceIdentifier kustoDatabaseResourceId = KustoDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
KustoDatabaseResource kustoDatabase = client.GetKustoDatabaseResource(kustoDatabaseResourceId);

// get the collection of this KustoScriptResource
KustoScriptCollection collection = kustoDatabase.GetKustoScripts();

// invoke the operation and iterate over the result
await foreach (KustoScriptResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    KustoScriptData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
