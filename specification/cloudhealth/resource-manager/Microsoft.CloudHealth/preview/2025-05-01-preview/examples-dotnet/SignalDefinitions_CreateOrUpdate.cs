using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/SignalDefinitions_CreateOrUpdate.json
// this example is just showing the usage of "SignalDefinition_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelSignalDefinitionResource created on azure
// for more information of creating HealthModelSignalDefinitionResource, please refer to the document of HealthModelSignalDefinitionResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
string healthModelName = "myHealthModel";
string signalDefinitionName = "sig1";
ResourceIdentifier healthModelSignalDefinitionResourceId = HealthModelSignalDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName, signalDefinitionName);
HealthModelSignalDefinitionResource healthModelSignalDefinition = client.GetHealthModelSignalDefinitionResource(healthModelSignalDefinitionResourceId);

// invoke the operation
HealthModelSignalDefinitionData data = new HealthModelSignalDefinitionData
{
    Properties = new ResourceMetricSignalDefinitionProperties(new EntitySignalEvaluationRule
    {
        DegradedRule = new EntitySignalThresholdRule(EntitySignalOperator.LowerThan, "65"),
        UnhealthyRule = new EntitySignalThresholdRule(EntitySignalOperator.LowerThan, "60"),
    }, "microsoft.compute/virtualMachines", "cpuusage", "PT1M", MetricAggregationType.None)
    {
        Dimension = "nodename",
        DimensionFilter = "node1",
        DisplayName = "cpu usage",
        RefreshInterval = EntitySignalRefreshInterval.PT1M,
        Labels =
        {
        ["key4788"] = "ixfvzsfnpvkkbrce"
        },
        DataUnit = "byte",
    },
};
ArmOperation<HealthModelSignalDefinitionResource> lro = await healthModelSignalDefinition.UpdateAsync(WaitUntil.Completed, data);
HealthModelSignalDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthModelSignalDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
