using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/threatintelligence/DeleteThreatIntelligence.json
// this example is just showing the usage of "ThreatIntelligenceIndicators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsThreatIntelligenceIndicatorResource created on azure
// for more information of creating SecurityInsightsThreatIntelligenceIndicatorResource, please refer to the document of SecurityInsightsThreatIntelligenceIndicatorResource
string subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string name = "d9cd6f0b-96b9-3984-17cd-a779d1e15a93";
ResourceIdentifier securityInsightsThreatIntelligenceIndicatorResourceId = SecurityInsightsThreatIntelligenceIndicatorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, name);
SecurityInsightsThreatIntelligenceIndicatorResource securityInsightsThreatIntelligenceIndicator = client.GetSecurityInsightsThreatIntelligenceIndicatorResource(securityInsightsThreatIntelligenceIndicatorResourceId);

// invoke the operation
await securityInsightsThreatIntelligenceIndicator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
