using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: 2025-03-01-preview/ServicePutOperation_example_max.json
// this example is just showing the usage of "ServiceResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedApplicationResource created on azure
// for more information of creating ServiceFabricManagedApplicationResource, please refer to the document of ServiceFabricManagedApplicationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string applicationName = "myApp";
ResourceIdentifier serviceFabricManagedApplicationResourceId = ServiceFabricManagedApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationName);
ServiceFabricManagedApplicationResource serviceFabricManagedApplication = client.GetServiceFabricManagedApplicationResource(serviceFabricManagedApplicationResourceId);

// get the collection of this ServiceFabricManagedServiceResource
ServiceFabricManagedServiceCollection collection = serviceFabricManagedApplication.GetServiceFabricManagedServices();

// invoke the operation
string serviceName = "myService";
ServiceFabricManagedServiceData data = new ServiceFabricManagedServiceData(new AzureLocation("eastus"))
{
    Properties = new StatelessServiceProperties("myServiceType", new SingletonPartitionScheme(), 5)
    {
        MinInstanceCount = 3,
        MinInstancePercentage = 30,
        ServicePackageActivationMode = ManagedServicePackageActivationMode.SharedProcess,
        ServiceDnsName = "myservicednsname.myApp",
        PlacementConstraints = "NodeType==frontend",
        CorrelationScheme = { new ManagedServiceCorrelation(ManagedServiceCorrelationScheme.AlignedAffinity, "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService1") },
        ServiceLoadMetrics = {new ManagedServiceLoadMetric("metric1")
        {
        Weight = ManagedServiceLoadMetricWeight.Low,
        DefaultLoad = 3,
        }},
        ServicePlacementPolicies = { new ServicePlacementNonPartiallyPlaceServicePolicy() },
        DefaultMoveCost = ServiceFabricManagedServiceMoveCost.Medium,
        ScalingPolicies = { new ManagedServiceScalingPolicy(new PartitionInstanceCountScalingMechanism(3, 9, 2), new AveragePartitionLoadScalingTrigger("metricName", 2, 8, "00:01:00")) },
    },
    Tags =
    {
    ["a"] = "b"
    },
};
ArmOperation<ServiceFabricManagedServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceName, data);
ServiceFabricManagedServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
