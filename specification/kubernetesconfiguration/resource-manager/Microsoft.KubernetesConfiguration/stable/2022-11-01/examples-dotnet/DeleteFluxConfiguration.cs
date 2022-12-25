using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration.Models;
using Azure.ResourceManager.KubernetesConfiguration;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteFluxConfiguration.json
// this example is just showing the usage of "FluxConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

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
await kubernetesFluxConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
