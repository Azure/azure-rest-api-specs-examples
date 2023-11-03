using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionCreate.json
// this example is just showing the usage of "NetworkFunctionDefinitionVersions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionDefinitionGroupResource created on azure
// for more information of creating NetworkFunctionDefinitionGroupResource, please refer to the document of NetworkFunctionDefinitionGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string publisherName = "TestPublisher";
string networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
ResourceIdentifier networkFunctionDefinitionGroupResourceId = NetworkFunctionDefinitionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, networkFunctionDefinitionGroupName);
NetworkFunctionDefinitionGroupResource networkFunctionDefinitionGroup = client.GetNetworkFunctionDefinitionGroupResource(networkFunctionDefinitionGroupResourceId);

// get the collection of this NetworkFunctionDefinitionVersionResource
NetworkFunctionDefinitionVersionCollection collection = networkFunctionDefinitionGroup.GetNetworkFunctionDefinitionVersions();

// invoke the operation
string networkFunctionDefinitionVersionName = "1.0.0";
NetworkFunctionDefinitionVersionData data = new NetworkFunctionDefinitionVersionData(new AzureLocation("eastus"))
{
    Properties = new ContainerizedNetworkFunctionDefinitionVersion()
    {
        NetworkFunctionTemplate = new AzureArcKubernetesNetworkFunctionTemplate()
        {
            NetworkFunctionApplications =
            {
            new AzureArcKubernetesHelmApplication()
            {
            ArtifactProfile = new AzureArcKubernetesArtifactProfile()
            {
            HelmArtifactProfile = new HelmArtifactProfile()
            {
            HelmPackageName = "fed-rbac",
            HelmPackageVersionRange = "~2.1.3",
            RegistryValuesPaths =
            {
            "global.registry.docker.repoPath"
            },
            ImagePullSecretsValuesPaths =
            {
            "global.imagePullSecrets"
            },
            },
            ArtifactStoreId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/testArtifactStore"),
            },
            DeployParametersMappingRuleProfile = new AzureArcKubernetesDeployMappingRuleProfile()
            {
            HelmMappingRuleProfile = new HelmMappingRuleProfile()
            {
            ReleaseNamespace = "{deployParameters.namesapce}",
            ReleaseName = "{deployParameters.releaseName}",
            HelmPackageVersion = "2.1.3",
            Values = "",
            Options = new HelmMappingRuleProfileConfig()
            {
            InstallOptions = new HelmInstallConfig()
            {
            Atomic = "true",
            Wait = "waitValue",
            Timeout = "30",
            },
            UpgradeOptions = new HelmUpgradeConfig()
            {
            Atomic = "true",
            Wait = "waitValue",
            Timeout = "30",
            },
            },
            },
            ApplicationEnablement = ApplicationEnablement.Enabled,
            },
            Name = "fedrbac",
            DependsOnProfile = new DependsOnProfile()
            {
            InstallDependsOn =
            {
            },
            UninstallDependsOn =
            {
            },
            UpdateDependsOn =
            {
            },
            },
            }
            },
        },
        DeployParameters = "{\"type\":\"object\",\"properties\":{\"releaseName\":{\"type\":\"string\"},\"namespace\":{\"type\":\"string\"}}}",
    },
};
ArmOperation<NetworkFunctionDefinitionVersionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkFunctionDefinitionVersionName, data);
NetworkFunctionDefinitionVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFunctionDefinitionVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
