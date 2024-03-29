using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ProvisionedClustersGetUpgradeProfile.json
// this example is just showing the usage of "ProvisionedClusters_GetUpgradeProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProvisionedClusterUpgradeProfileResource created on azure
// for more information of creating ProvisionedClusterUpgradeProfileResource, please refer to the document of ProvisionedClusterUpgradeProfileResource
string subscriptionId = "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
string resourceGroupName = "test-arcappliance-resgrp";
string resourceName = "test-hybridakscluster";
ResourceIdentifier provisionedClusterUpgradeProfileResourceId = ProvisionedClusterUpgradeProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ProvisionedClusterUpgradeProfileResource provisionedClusterUpgradeProfile = client.GetProvisionedClusterUpgradeProfileResource(provisionedClusterUpgradeProfileResourceId);

// invoke the operation
ProvisionedClusterUpgradeProfileResource result = await provisionedClusterUpgradeProfile.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProvisionedClusterUpgradeProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
