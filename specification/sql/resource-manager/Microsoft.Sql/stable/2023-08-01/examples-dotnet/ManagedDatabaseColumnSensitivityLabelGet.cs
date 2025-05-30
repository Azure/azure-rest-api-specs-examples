using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ManagedDatabaseColumnSensitivityLabelGet.json
// this example is just showing the usage of "ManagedDatabaseSensitivityLabels_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseColumnResource created on azure
// for more information of creating ManagedDatabaseColumnResource, please refer to the document of ManagedDatabaseColumnResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string managedInstanceName = "myManagedInstanceName";
string databaseName = "myDatabase";
string schemaName = "dbo";
string tableName = "myTable";
string columnName = "myColumn";
ResourceIdentifier managedDatabaseColumnResourceId = ManagedDatabaseColumnResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName);
ManagedDatabaseColumnResource managedDatabaseColumn = client.GetManagedDatabaseColumnResource(managedDatabaseColumnResourceId);

// get the collection of this ManagedDatabaseSensitivityLabelResource
ManagedDatabaseSensitivityLabelCollection collection = managedDatabaseColumn.GetManagedDatabaseSensitivityLabels();

// invoke the operation
SensitivityLabelSource sensitivityLabelSource = SensitivityLabelSource.Current;
NullableResponse<ManagedDatabaseSensitivityLabelResource> response = await collection.GetIfExistsAsync(sensitivityLabelSource);
ManagedDatabaseSensitivityLabelResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SensitivityLabelData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
