using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana.Models;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2024-11-01-preview/IntegrationFabrics_Create.json
// this example is just showing the usage of "IntegrationFabric_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedGrafanaResource created on azure
// for more information of creating ManagedGrafanaResource, please refer to the document of ManagedGrafanaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string workspaceName = "myWorkspace";
ResourceIdentifier managedGrafanaResourceId = ManagedGrafanaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
ManagedGrafanaResource managedGrafana = client.GetManagedGrafanaResource(managedGrafanaResourceId);

// get the collection of this GrafanaIntegrationFabricResource
GrafanaIntegrationFabricCollection collection = managedGrafana.GetGrafanaIntegrationFabrics();

// invoke the operation
string integrationFabricName = "sampleIntegration";
GrafanaIntegrationFabricData data = new GrafanaIntegrationFabricData(new AzureLocation("West US"))
{
    Properties = new GrafanaIntegrationFabricProperties
    {
        TargetResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myAks"),
        DataSourceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Monitor/accounts/myAmw"),
        Scenarios = { "scenario1", "scenario2" },
    },
};
ArmOperation<GrafanaIntegrationFabricResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, integrationFabricName, data);
GrafanaIntegrationFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GrafanaIntegrationFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
