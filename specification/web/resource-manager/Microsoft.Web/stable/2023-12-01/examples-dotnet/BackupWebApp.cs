using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/BackupWebApp.json
// this example is just showing the usage of "WebApps_Backup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "sitef6141";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// invoke the operation
WebAppBackupInfo info = new WebAppBackupInfo()
{
    BackupName = "abcdwe",
    IsEnabled = true,
    StorageAccountUri = new Uri("DefaultEndpointsProtocol=https;AccountName=storagesample;AccountKey=<account-key>"),
    BackupSchedule = new WebAppBackupSchedule(7, BackupFrequencyUnit.Day, true, 30)
    {
        StartOn = DateTimeOffset.Parse("2022-09-02T17:33:11.641Z"),
    },
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
WebAppBackupData result = await webSite.BackupAsync(info);

// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {result.Id}");
