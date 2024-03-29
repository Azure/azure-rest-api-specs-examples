using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration;
using Azure.ResourceManager.KubernetesConfiguration.Models;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/PatchFluxConfiguration.json
// this example is just showing the usage of "FluxConfigurations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubernetesFluxConfigurationResource created on azure
// for more information of creating KubernetesFluxConfigurationResource, please refer to the document of KubernetesFluxConfigurationResource
string subscriptionId = "subId1";
string resourceGroupName = "rg1";
string clusterRp = "Microsoft.Kubernetes";
string clusterResourceName = "connectedClusters";
string clusterName = "clusterName1";
string fluxConfigurationName = "srs-fluxconfig";
ResourceIdentifier kubernetesFluxConfigurationResourceId = KubernetesFluxConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterRp, clusterResourceName, clusterName, fluxConfigurationName);
KubernetesFluxConfigurationResource kubernetesFluxConfiguration = client.GetKubernetesFluxConfigurationResource(kubernetesFluxConfigurationResourceId);

// invoke the operation
KubernetesFluxConfigurationPatch patch = new KubernetesFluxConfigurationPatch()
{
    Suspend = true,
    GitRepository = new KubernetesGitRepositoryUpdateContent()
    {
        Uri = new Uri("https://github.com/jonathan-innis/flux2-kustomize-helm-example.git"),
    },
    Kustomizations =
    {
    ["srs-kustomization1"] = null,
    ["srs-kustomization2"] = new KustomizationUpdateContent()
    {
    Path = "./test/alt-path",
    DependsOn =
    {
    },
    SyncIntervalInSeconds = 300,
    },
    ["srs-kustomization3"] = new KustomizationUpdateContent()
    {
    Path = "./test/another-path",
    SyncIntervalInSeconds = 300,
    },
    },
};
ArmOperation<KubernetesFluxConfigurationResource> lro = await kubernetesFluxConfiguration.UpdateAsync(WaitUntil.Completed, patch);
KubernetesFluxConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubernetesFluxConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
