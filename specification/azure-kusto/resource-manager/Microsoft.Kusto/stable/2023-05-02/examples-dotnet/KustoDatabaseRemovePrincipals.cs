using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabaseRemovePrincipals.json
// this example is just showing the usage of "Databases_RemovePrincipals" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDatabaseResource created on azure
// for more information of creating KustoDatabaseResource, please refer to the document of KustoDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "KustoDatabase8";
ResourceIdentifier kustoDatabaseResourceId = KustoDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
KustoDatabaseResource kustoDatabase = client.GetKustoDatabaseResource(kustoDatabaseResourceId);

// invoke the operation and iterate over the result
DatabasePrincipalList databasePrincipalsToRemove = new DatabasePrincipalList()
{
    Value =
    {
    new KustoDatabasePrincipal(KustoDatabasePrincipalRole.Admin,"Some User",KustoDatabasePrincipalType.User)
    {
    Fqn = "aaduser=some_guid",
    Email = "user@microsoft.com",
    AppId = "",
    },new KustoDatabasePrincipal(KustoDatabasePrincipalRole.Viewer,"Kusto",KustoDatabasePrincipalType.Group)
    {
    Fqn = "aadgroup=some_guid",
    Email = "kusto@microsoft.com",
    AppId = "",
    },new KustoDatabasePrincipal(KustoDatabasePrincipalRole.Admin,"SomeApp",KustoDatabasePrincipalType.App)
    {
    Fqn = "aadapp=some_guid_app_id",
    Email = "",
    AppId = "some_guid_app_id",
    }
    },
};
await foreach (KustoDatabasePrincipal item in kustoDatabase.RemovePrincipalsAsync(databasePrincipalsToRemove))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
