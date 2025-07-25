using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.HardwareSecurityModules;

// Generated from example definition: 2025-03-31/CloudHsmClusterPrivateLinkResource_ListByCloudHsmCluster_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmClusters_ListByCloudHsmCluster" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudHsmClusterResource created on azure
// for more information of creating CloudHsmClusterResource, please refer to the document of CloudHsmClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgcloudhsm";
string cloudHsmClusterName = "chsm1";
ResourceIdentifier cloudHsmClusterResourceId = CloudHsmClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudHsmClusterName);
CloudHsmClusterResource cloudHsmCluster = client.GetCloudHsmClusterResource(cloudHsmClusterResourceId);

// invoke the operation and iterate over the result
await foreach (CloudHsmClusterPrivateLinkData item in cloudHsmCluster.GetCloudHsmClusterPrivateLinkResourcesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
