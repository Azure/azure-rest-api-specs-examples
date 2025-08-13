using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Files_Get.json
// this example is just showing the usage of "Files_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataMigrationProjectFileResource created on azure
// for more information of creating DataMigrationProjectFileResource, please refer to the document of DataMigrationProjectFileResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
string serviceName = "DmsSdkService";
string projectName = "DmsSdkProject";
string fileName = "x114d023d8";
ResourceIdentifier dataMigrationProjectFileResourceId = DataMigrationProjectFileResource.CreateResourceIdentifier(subscriptionId, groupName, serviceName, projectName, fileName);
DataMigrationProjectFileResource dataMigrationProjectFile = client.GetDataMigrationProjectFileResource(dataMigrationProjectFileResourceId);

// invoke the operation
DataMigrationProjectFileResource result = await dataMigrationProjectFile.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataMigrationProjectFileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
