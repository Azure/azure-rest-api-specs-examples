using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetWebAppBackup.json
// this example is just showing the usage of "WebApps_GetBackupStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteBackupResource created on azure
// for more information of creating SiteBackupResource, please refer to the document of SiteBackupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "sitef6141";
string backupId = "12345";
ResourceIdentifier siteBackupResourceId = SiteBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, backupId);
SiteBackupResource siteBackup = client.GetSiteBackupResource(siteBackupResourceId);

// invoke the operation
SiteBackupResource result = await siteBackup.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebAppBackupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
