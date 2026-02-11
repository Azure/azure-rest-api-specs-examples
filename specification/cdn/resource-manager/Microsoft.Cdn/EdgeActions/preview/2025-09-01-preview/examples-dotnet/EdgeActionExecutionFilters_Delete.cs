using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EdgeActions.Models;
using Azure.ResourceManager.EdgeActions;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/EdgeActions/preview/2025-09-01-preview/examples/EdgeActionExecutionFilters_Delete.json
// this example is just showing the usage of "EdgeActionExecutionFilters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeActionExecutionFilterResource created on azure
// for more information of creating EdgeActionExecutionFilterResource, please refer to the document of EdgeActionExecutionFilterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testrg";
string edgeActionName = "edgeAction1";
string executionFilter = "executionFilter1";
ResourceIdentifier edgeActionExecutionFilterResourceId = EdgeActionExecutionFilterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, edgeActionName, executionFilter);
EdgeActionExecutionFilterResource edgeActionExecutionFilter = client.GetEdgeActionExecutionFilterResource(edgeActionExecutionFilterResourceId);

// invoke the operation
await edgeActionExecutionFilter.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
