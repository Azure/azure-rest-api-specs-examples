using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HardwareSecurityModules;

// Generated from example definition: 2025-03-31/CloudHsmCluster_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmCluster_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgcloudhsm";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CloudHsmClusterResource
CloudHsmClusterCollection collection = resourceGroupResource.GetCloudHsmClusters();

// invoke the operation
string cloudHsmClusterName = "chsm1";
CloudHsmClusterData data = new CloudHsmClusterData(new AzureLocation("eastus2"))
{
    Properties = new CloudHsmClusterProperties
    {
        PublicNetworkAccess = CloudHsmClusterPublicNetworkAccess.Disabled,
    },
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity-1")] = new UserAssignedIdentity()
        },
    },
    Sku = new CloudHsmClusterSku(CloudHsmClusterSkuFamily.B, CloudHsmClusterSkuName.StandardB1),
    Tags =
    {
    ["Dept"] = "hsm",
    ["Environment"] = "dogfood"
    },
};
ArmOperation<CloudHsmClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, cloudHsmClusterName, data);
CloudHsmClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudHsmClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
