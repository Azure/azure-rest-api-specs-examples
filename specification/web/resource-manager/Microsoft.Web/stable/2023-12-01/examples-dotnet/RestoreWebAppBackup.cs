using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/RestoreWebAppBackup.json
// this example is just showing the usage of "WebApps_Restore" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteBackupResource created on azure
// for more information of creating SiteBackupResource, please refer to the document of SiteBackupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "sitef6141";
string backupId = "123244";
ResourceIdentifier siteBackupResourceId = SiteBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, backupId);
SiteBackupResource siteBackup = client.GetSiteBackupResource(siteBackupResourceId);

// invoke the operation
RestoreRequestInfo info = new RestoreRequestInfo()
{
    StorageAccountUri = new Uri("DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>"),
    CanOverwrite = true,
    SiteName = "sitef6141",
    Databases =
    {
    new AppServiceDatabaseBackupSetting(AppServiceDatabaseType.SqlAzure)
    {
    Name = "backenddb",
    ConnectionStringName = "backend",
    ConnectionString = "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
    },new AppServiceDatabaseBackupSetting(AppServiceDatabaseType.SqlAzure)
    {
    Name = "statsdb",
    ConnectionStringName = "stats",
    ConnectionString = "DSN=data-source-name[;SERVER=value] [;PWD=value] [;UID=value] [;<Attribute>=<value>]",
    }
    },
};
await siteBackup.RestoreAsync(WaitUntil.Completed, info);

Console.WriteLine($"Succeeded");
