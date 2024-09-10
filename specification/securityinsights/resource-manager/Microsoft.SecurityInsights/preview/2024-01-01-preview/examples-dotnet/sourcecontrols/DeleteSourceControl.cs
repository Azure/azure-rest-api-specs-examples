using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/sourcecontrols/DeleteSourceControl.json
// this example is just showing the usage of "SourceControls_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsSourceControlResource created on azure
// for more information of creating SecurityInsightsSourceControlResource, please refer to the document of SecurityInsightsSourceControlResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string sourceControlId = "789e0c1f-4a3d-43ad-809c-e713b677b04a";
ResourceIdentifier securityInsightsSourceControlResourceId = SecurityInsightsSourceControlResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sourceControlId);
SecurityInsightsSourceControlResource securityInsightsSourceControl = client.GetSecurityInsightsSourceControlResource(securityInsightsSourceControlResourceId);

// invoke the operation
RepositoryAccessProperties repositoryAccess = new RepositoryAccessProperties(RepositoryAccessKind.OAuth)
{
    Code = "939fd7c6caf754f4f41f",
    State = "state",
    ClientId = "54b3c2c0-1f48-4a1c-af9f-6399c3240b73",
};
SourceControlOperationWarning result = await securityInsightsSourceControl.DeleteAsync(repositoryAccess);

Console.WriteLine($"Succeeded: {result}");
