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

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/UpdateAzureStorageAccounts.json
// this example is just showing the usage of "WebApps_UpdateAzureStorageAccounts" operation, for the dependent resources, they will have to be created separately.

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
AzureStoragePropertyDictionary azureStorageAccounts = new AzureStoragePropertyDictionary()
{
    Properties =
    {
    ["account1"] = new AppServiceStorageAccessInfo()
    {
    StorageType = AppServiceStorageType.AzureFiles,
    AccountName = "testsa",
    ShareName = "web",
    AccessKey = "26515^%@#*",
    MountPath = "/mounts/a/files",
    },
    },
};
AzureStoragePropertyDictionary result = await webSite.UpdateAzureStorageAccountsAsync(azureStorageAccounts);

Console.WriteLine($"Succeeded: {result}");
