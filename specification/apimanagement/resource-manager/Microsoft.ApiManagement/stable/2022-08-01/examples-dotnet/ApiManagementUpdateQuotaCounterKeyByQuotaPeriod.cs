using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateQuotaCounterKeyByQuotaPeriod.json
// this example is just showing the usage of "QuotaByPeriodKeys_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementServiceResource created on azure
// for more information of creating ApiManagementServiceResource, please refer to the document of ApiManagementServiceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
ResourceIdentifier apiManagementServiceResourceId = ApiManagementServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
ApiManagementServiceResource apiManagementService = client.GetApiManagementServiceResource(apiManagementServiceResourceId);

// invoke the operation
string quotaCounterKey = "ba";
string quotaPeriodKey = "0_P3Y6M4DT12H30M5S";
QuotaCounterValueUpdateContent content = new QuotaCounterValueUpdateContent()
{
    CallsCount = 0,
    KbTransferred = 0,
};
QuotaCounterContract result = await apiManagementService.UpdateQuotaByPeriodKeyAsync(quotaCounterKey, quotaPeriodKey, content);

Console.WriteLine($"Succeeded: {result}");
