using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AlertsManagement;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/AlertsMetaData_MonitorService.json
// this example is just showing the usage of "Alerts_MetaData" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
RetrievedInformationIdentifier identifier = RetrievedInformationIdentifier.MonitorServiceList;
ServiceAlertMetadata result = await tenantResource.GetServiceAlertMetadataAsync(identifier);

Console.WriteLine($"Succeeded: {result}");
