using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityDevOps.Models;
using Azure.ResourceManager.SecurityDevOps;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorUpdate.json
// this example is just showing the usage of "AzureDevOpsConnector_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
AzureDevOpsConnectorData data = new AzureDevOpsConnectorData(new AzureLocation("West US"))
{
    Tags =
    {
    ["client"] = "dev-client",
    ["env"] = "dev"
    },
};
ArmOperation<AzureDevOpsConnectorResource> lro = await azureDevOpsConnector.UpdateAsync(WaitUntil.Completed, data);
AzureDevOpsConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AzureDevOpsConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
