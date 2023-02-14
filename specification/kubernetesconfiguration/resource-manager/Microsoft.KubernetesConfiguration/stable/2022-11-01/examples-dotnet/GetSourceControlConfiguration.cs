using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration;
using Azure.ResourceManager.KubernetesConfiguration.Models;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/GetSourceControlConfiguration.json
// this example is just showing the usage of "SourceControlConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

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
KubernetesSourceControlConfigurationResource result = await kubernetesSourceControlConfiguration.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubernetesSourceControlConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
