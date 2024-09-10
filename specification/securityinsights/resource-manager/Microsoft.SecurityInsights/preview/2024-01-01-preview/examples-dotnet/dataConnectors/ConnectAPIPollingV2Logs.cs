using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectors/ConnectAPIPollingV2Logs.json
// this example is just showing the usage of "DataConnectors_Connect" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsDataConnectorResource created on azure
// for more information of creating SecurityInsightsDataConnectorResource, please refer to the document of SecurityInsightsDataConnectorResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string dataConnectorId = "316ec55e-7138-4d63-ab18-90c8a60fd1c8";
ResourceIdentifier securityInsightsDataConnectorResourceId = SecurityInsightsDataConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, dataConnectorId);
SecurityInsightsDataConnectorResource securityInsightsDataConnector = client.GetSecurityInsightsDataConnectorResource(securityInsightsDataConnectorResourceId);

// invoke the operation
DataConnectorConnectContent content = new DataConnectorConnectContent()
{
    Kind = ConnectAuthKind.APIKey,
    ApiKey = "123456789",
    DataCollectionEndpoint = "https://test.eastus.ingest.monitor.azure.com",
    DataCollectionRuleImmutableId = "dcr-34adsj9o7d6f9de204478b9cgb43b631",
    OutputStream = "Custom-MyTableRawData",
    RequestConfigUserInputValues =
    {
    BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    ["displayText"] = "Organization Name",
    ["placeHolderName"] = "{{placeHolder1}}",
    ["placeHolderValue"] = "somePlaceHolderValue",
    ["requestObjectKey"] = "apiEndpoint"})
    },
};
await securityInsightsDataConnector.ConnectAsync(content);

Console.WriteLine($"Succeeded");
