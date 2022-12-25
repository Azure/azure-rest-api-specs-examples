using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration.Models;
using Azure.ResourceManager.KubernetesConfiguration;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteSourceControlConfiguration.json
// this example is just showing the usage of "SourceControlConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this KubernetesSourceControlConfigurationResource created on azure
// for more information of creating KubernetesSourceControlConfigurationResource, please refer to the document of KubernetesSourceControlConfigurationResource
string subscriptionId = "subId1";
string resourceGroupName = "rg1";
string clusterRp = "Microsoft.Kubernetes";
string clusterResourceName = "connectedClusters";
string clusterName = "clusterName1";
string sourceControlConfigurationName = "SRS_GitHubConfig";
ResourceIdentifier kubernetesSourceControlConfigurationResourceId = KubernetesSourceControlConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterRp, clusterResourceName, clusterName, sourceControlConfigurationName);
KubernetesSourceControlConfigurationResource kubernetesSourceControlConfiguration = client.GetKubernetesSourceControlConfigurationResource(kubernetesSourceControlConfigurationResourceId);

// invoke the operation
await kubernetesSourceControlConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
