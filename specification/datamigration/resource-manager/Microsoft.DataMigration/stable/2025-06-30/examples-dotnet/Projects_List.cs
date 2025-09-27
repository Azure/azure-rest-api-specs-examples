using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/Projects_List.json
// this example is just showing the usage of "Projects_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataMigrationServiceResource created on azure
// for more information of creating DataMigrationServiceResource, please refer to the document of DataMigrationServiceResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
string serviceName = "DmsSdkService";
ResourceIdentifier dataMigrationServiceResourceId = DataMigrationServiceResource.CreateResourceIdentifier(subscriptionId, groupName, serviceName);
DataMigrationServiceResource dataMigrationService = client.GetDataMigrationServiceResource(dataMigrationServiceResourceId);

// get the collection of this DataMigrationProjectResource
DataMigrationProjectCollection collection = dataMigrationService.GetDataMigrationProjects();

// invoke the operation and iterate over the result
await foreach (DataMigrationProjectResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataMigrationProjectData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
