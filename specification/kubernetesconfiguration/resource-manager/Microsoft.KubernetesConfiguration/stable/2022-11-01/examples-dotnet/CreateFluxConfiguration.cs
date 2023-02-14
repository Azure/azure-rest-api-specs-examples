using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration;
using Azure.ResourceManager.KubernetesConfiguration.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateFluxConfiguration.json
// this example is just showing the usage of "FluxConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId1";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this KubernetesFluxConfigurationResource
string clusterRp = "Microsoft.Kubernetes";
string clusterResourceName = "connectedClusters";
string clusterName = "clusterName1";
KubernetesFluxConfigurationCollection collection = resourceGroupResource.GetKubernetesFluxConfigurations(clusterRp, clusterResourceName, clusterName);

// invoke the operation
string fluxConfigurationName = "srs-fluxconfig";
KubernetesFluxConfigurationData data = new KubernetesFluxConfigurationData()
{
    Scope = KubernetesConfigurationScope.Cluster,
    Namespace = "srs-namespace",
    SourceKind = KubernetesConfigurationSourceKind.GitRepository,
    IsReconciliationSuspended = false,
    GitRepository = new KubernetesGitRepository()
    {
        Uri = new Uri("https://github.com/Azure/arc-k8s-demo"),
        TimeoutInSeconds = 600,
        SyncIntervalInSeconds = 600,
        RepositoryRef = new KubernetesGitRepositoryRef()
        {
            Branch = "master",
        },
        HttpsCACert = "ZXhhbXBsZWNlcnRpZmljYXRl",
    },
    Kustomizations =
    {
    ["srs-kustomization1"] = new Kustomization()
    {
    Path = "./test/path",
    DependsOn =
    {
    },
    TimeoutInSeconds = 600,
    SyncIntervalInSeconds = 600,
    },
    ["srs-kustomization2"] = new Kustomization()
    {
    Path = "./other/test/path",
    DependsOn =
    {
    "srs-kustomization1"
    },
    TimeoutInSeconds = 600,
    SyncIntervalInSeconds = 600,
    RetryIntervalInSeconds = 600,
    Prune = false,
    },
    },
};
ArmOperation<KubernetesFluxConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, fluxConfigurationName, data);
KubernetesFluxConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubernetesFluxConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
