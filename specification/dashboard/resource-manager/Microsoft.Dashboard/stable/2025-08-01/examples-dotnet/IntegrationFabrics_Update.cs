using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2025-08-01/IntegrationFabrics_Update.json
// this example is just showing the usage of "IntegrationFabric_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GrafanaIntegrationFabricResource created on azure
// for more information of creating GrafanaIntegrationFabricResource, please refer to the document of GrafanaIntegrationFabricResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string workspaceName = "myWorkspace";
string integrationFabricName = "sampleIntegration";
ResourceIdentifier grafanaIntegrationFabricResourceId = GrafanaIntegrationFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, integrationFabricName);
GrafanaIntegrationFabricResource grafanaIntegrationFabric = client.GetGrafanaIntegrationFabricResource(grafanaIntegrationFabricResourceId);

// invoke the operation
GrafanaIntegrationFabricPatch patch = new GrafanaIntegrationFabricPatch
{
    Tags =
    {
    ["Environment"] = "Dev 2"
    },
    IntegrationFabricPropertiesUpdateParametersScenarios = { "scenario1" },
};
ArmOperation<GrafanaIntegrationFabricResource> lro = await grafanaIntegrationFabric.UpdateAsync(WaitUntil.Completed, patch);
GrafanaIntegrationFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GrafanaIntegrationFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
