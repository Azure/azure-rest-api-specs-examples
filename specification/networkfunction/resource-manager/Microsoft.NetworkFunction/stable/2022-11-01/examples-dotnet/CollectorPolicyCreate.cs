using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkFunction.Models;
using Azure.ResourceManager.NetworkFunction;

// Generated from example definition: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyCreate.json
// this example is just showing the usage of "CollectorPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AzureTrafficCollectorResource created on azure
// for more information of creating AzureTrafficCollectorResource, please refer to the document of AzureTrafficCollectorResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string azureTrafficCollectorName = "atc";
ResourceIdentifier azureTrafficCollectorResourceId = AzureTrafficCollectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureTrafficCollectorName);
AzureTrafficCollectorResource azureTrafficCollector = client.GetAzureTrafficCollectorResource(azureTrafficCollectorResourceId);

// get the collection of this CollectorPolicyResource
CollectorPolicyCollection collection = azureTrafficCollector.GetCollectorPolicies();

// invoke the operation
string collectorPolicyName = "cp1";
CollectorPolicyData data = new CollectorPolicyData(new AzureLocation("West US"))
{
    IngestionPolicy = new IngestionPolicyPropertiesFormat
    {
        IngestionType = IngestionType.Ipfix,
        IngestionSources = {new IngestionSourcesPropertiesFormat
        {
        SourceType = IngestionSourceType.Resource,
        ResourceId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteCircuits/circuitName",
        }},
    },
    EmissionPolicies = {new EmissionPoliciesPropertiesFormat
    {
    EmissionType = EmissionType.Ipfix,
    EmissionDestinations = {new EmissionPolicyDestination
    {
    DestinationType = EmissionDestinationType.AzureMonitor,
    }},
    }},
};
ArmOperation<CollectorPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, collectorPolicyName, data);
CollectorPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CollectorPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
