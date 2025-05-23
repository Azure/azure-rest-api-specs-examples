using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkAnalytics.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetworkAnalytics;

// Generated from example definition: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Create_MaximumSet_Gen.json
// this example is just showing the usage of "DataProducts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "aoiresourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataProductResource
DataProductCollection collection = resourceGroupResource.GetDataProducts();

// invoke the operation
string dataProductName = "dataproduct01";
DataProductData data = new DataProductData(new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    Publisher = "Microsoft",
    Product = "MCC",
    MajorVersion = "1.0.0",
    Owners = { "abc@micros.com" },
    Redundancy = DataProductControlState.Disabled,
    PurviewAccount = "testpurview",
    PurviewCollection = "134567890",
    PrivateLinksEnabled = DataProductControlState.Disabled,
    PublicNetworkAccess = DataProductControlState.Enabled,
    CustomerManagedKeyEncryptionEnabled = DataProductControlState.Enabled,
    CustomerEncryptionKey = new EncryptionKeyDetails(new Uri("https://KeyVault.vault.azure.net"), "keyName", "keyVersion"),
    Networkacls = new DataProductNetworkAcls(new NetworkAnalyticsVirtualNetworkRule[]
{
new NetworkAnalyticsVirtualNetworkRule("/subscriptions/subscriptionId/resourcegroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/virtualNetworkName/subnets/subnetName")
{
Action = "Allow",
State = "",
}
}, new NetworkAnalyticsIPRules[]
{
new NetworkAnalyticsIPRules("Allow")
{
Value = "1.1.1.1",
}
}, new string[] { "1.1.1.1" }, NetworkAclDefaultAction.Allow),
    ManagedResourceGroupConfiguration = new NetworkAnalyticsManagedResourceGroupConfiguration("managedResourceGroupName", new AzureLocation("eastus")),
    CurrentMinorVersion = "1.0.1",
    Tags =
    {
    ["userSpecifiedKeyName"] = "userSpecifiedKeyValue"
    },
};
ArmOperation<DataProductResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataProductName, data);
DataProductResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataProductData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
