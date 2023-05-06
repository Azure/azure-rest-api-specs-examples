using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Qumulo;
using Azure.ResourceManager.Qumulo.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "FileSystems_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QumuloFileSystemResource created on azure
// for more information of creating QumuloFileSystemResource, please refer to the document of QumuloFileSystemResource
string subscriptionId = "ulseeqylxb";
string resourceGroupName = "rgQumulo";
string fileSystemName = "nauwwbfoqehgbhdsmkewoboyxeqg";
ResourceIdentifier qumuloFileSystemResourceId = QumuloFileSystemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fileSystemName);
QumuloFileSystemResource qumuloFileSystemResource = client.GetQumuloFileSystemResource(qumuloFileSystemResourceId);

// invoke the operation
await qumuloFileSystemResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
