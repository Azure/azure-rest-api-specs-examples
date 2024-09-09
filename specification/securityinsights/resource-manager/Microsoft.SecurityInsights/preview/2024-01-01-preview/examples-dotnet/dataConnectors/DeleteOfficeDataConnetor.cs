using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/dataConnectors/DeleteOfficeDataConnetor.json
// this example is just showing the usage of "DataConnectors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsDataConnectorResource created on azure
// for more information of creating SecurityInsightsDataConnectorResource, please refer to the document of SecurityInsightsDataConnectorResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string dataConnectorId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
ResourceIdentifier securityInsightsDataConnectorResourceId = SecurityInsightsDataConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, dataConnectorId);
SecurityInsightsDataConnectorResource securityInsightsDataConnector = client.GetSecurityInsightsDataConnectorResource(securityInsightsDataConnectorResourceId);

// invoke the operation
await securityInsightsDataConnector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
