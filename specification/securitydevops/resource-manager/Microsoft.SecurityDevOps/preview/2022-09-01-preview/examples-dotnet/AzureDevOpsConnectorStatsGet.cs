using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityDevOps;
using Azure.ResourceManager.SecurityDevOps.Models;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorStatsGet.json
// this example is just showing the usage of "AzureDevOpsConnectorStats_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureDevOpsConnectorResource created on azure
// for more information of creating AzureDevOpsConnectorResource, please refer to the document of AzureDevOpsConnectorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
string azureDevOpsConnectorName = "testconnector";
ResourceIdentifier azureDevOpsConnectorResourceId = AzureDevOpsConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureDevOpsConnectorName);
AzureDevOpsConnectorResource azureDevOpsConnector = client.GetAzureDevOpsConnectorResource(azureDevOpsConnectorResourceId);

// invoke the operation and iterate over the result
await foreach (AzureDevOpsConnectorStats item in azureDevOpsConnector.GetAzureDevOpsConnectorStatsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
