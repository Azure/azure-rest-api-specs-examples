using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/Projects_Delete.json
// this example is just showing the usage of "Projects_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataMigrationProjectResource created on azure
// for more information of creating DataMigrationProjectResource, please refer to the document of DataMigrationProjectResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
string serviceName = "DmsSdkService";
string projectName = "DmsSdkProject";
ResourceIdentifier dataMigrationProjectResourceId = DataMigrationProjectResource.CreateResourceIdentifier(subscriptionId, groupName, serviceName, projectName);
DataMigrationProjectResource dataMigrationProject = client.GetDataMigrationProjectResource(dataMigrationProjectResourceId);

// invoke the operation
await dataMigrationProject.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
