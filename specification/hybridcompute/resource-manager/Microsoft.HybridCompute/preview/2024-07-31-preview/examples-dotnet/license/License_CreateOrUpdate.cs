using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/license/License_CreateOrUpdate.json
// this example is just showing the usage of "Licenses_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HybridComputeLicenseResource
HybridComputeLicenseCollection collection = resourceGroupResource.GetHybridComputeLicenses();

// invoke the operation
string licenseName = "{licenseName}";
HybridComputeLicenseData data = new HybridComputeLicenseData(new AzureLocation("eastus2euap"))
{
    LicenseType = HybridComputeLicenseType.Esu,
    LicenseDetails = new HybridComputeLicenseDetails
    {
        State = HybridComputeLicenseState.Activated,
        Target = HybridComputeLicenseTarget.WindowsServer2012,
        Edition = HybridComputeLicenseEdition.DataCenter,
        LicenseCoreType = LicenseCoreType.PCore,
        Processors = 6,
    },
};
ArmOperation<HybridComputeLicenseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, licenseName, data);
HybridComputeLicenseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputeLicenseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
