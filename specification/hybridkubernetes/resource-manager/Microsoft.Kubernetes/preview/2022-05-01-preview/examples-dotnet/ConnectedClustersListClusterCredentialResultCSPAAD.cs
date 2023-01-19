using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kubernetes;
using Azure.ResourceManager.Kubernetes.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridkubernetes/resource-manager/Microsoft.Kubernetes/preview/2022-05-01-preview/examples/ConnectedClustersListClusterCredentialResultCSPAAD.json
// this example is just showing the usage of "ConnectedCluster_ListClusterUserCredential" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectedClusterResource created on azure
// for more information of creating ConnectedClusterResource, please refer to the document of ConnectedClusterResource
string subscriptionId = "1bfbb5d0-917e-4346-9026-1d3b344417f5";
string resourceGroupName = "k8sc-rg";
string clusterName = "testCluster";
ResourceIdentifier connectedClusterResourceId = ConnectedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ConnectedClusterResource connectedCluster = client.GetConnectedClusterResource(connectedClusterResourceId);

// invoke the operation
ListClusterUserCredentialProperties properties = new ListClusterUserCredentialProperties(AuthenticationMethod.AAD, true);
CredentialResults result = await connectedCluster.GetClusterUserCredentialAsync(properties);

Console.WriteLine($"Succeeded: {result}");
