using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkAnalytics.Models;
using Azure.ResourceManager.NetworkAnalytics;

// Generated from example definition: specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_Update_MaximumSet_Gen.json
// this example is just showing the usage of "DataProducts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataProductResource created on azure
// for more information of creating DataProductResource, please refer to the document of DataProductResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "aoiresourceGroupName";
string dataProductName = "dataproduct01";
ResourceIdentifier dataProductResourceId = DataProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataProductName);
DataProductResource dataProduct = client.GetDataProductResource(dataProductResourceId);

// invoke the operation
DataProductPatch patch = new DataProductPatch
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["userSpecifiedKeyName"] = "userSpecifiedKeyValue"
    },
    Owners = { "abc@micros.com", "def@micros.com" },
    PurviewAccount = "testpurview",
    PurviewCollection = "134567890",
    PrivateLinksEnabled = DataProductControlState.Disabled,
    CurrentMinorVersion = "1.0.1",
};
ArmOperation<DataProductResource> lro = await dataProduct.UpdateAsync(WaitUntil.Completed, patch);
DataProductResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataProductData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
