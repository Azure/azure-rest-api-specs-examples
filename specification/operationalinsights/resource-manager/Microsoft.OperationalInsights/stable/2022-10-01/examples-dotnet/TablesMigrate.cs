using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.OperationalInsights;
using Azure.ResourceManager.OperationalInsights.Models;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/TablesMigrate.json
// this example is just showing the usage of "Tables_Migrate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsTableResource created on azure
// for more information of creating OperationalInsightsTableResource, please refer to the document of OperationalInsightsTableResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "oiautorest6685";
string workspaceName = "oiautorest6685";
string tableName = "table1_CL";
ResourceIdentifier operationalInsightsTableResourceId = OperationalInsightsTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, tableName);
OperationalInsightsTableResource operationalInsightsTable = client.GetOperationalInsightsTableResource(operationalInsightsTableResourceId);

// invoke the operation
await operationalInsightsTable.MigrateAsync();

Console.WriteLine($"Succeeded");
