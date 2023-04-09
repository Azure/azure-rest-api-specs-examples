using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;
using Azure.ResourceManager.HybridContainerService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/DeleteProvisionedCluster.json
// this example is just showing the usage of "ProvisionedClusters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProvisionedClusterResource created on azure
// for more information of creating ProvisionedClusterResource, please refer to the document of ProvisionedClusterResource
string subscriptionId = "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
string resourceGroupName = "test-arcappliance-resgrp";
string resourceName = "test-hybridakscluster";
ResourceIdentifier provisionedClusterResourceId = ProvisionedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ProvisionedClusterResource provisionedCluster = client.GetProvisionedClusterResource(provisionedClusterResourceId);

// invoke the operation
await provisionedCluster.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
