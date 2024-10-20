using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/licenseProfile/LicenseProfile_Update.json
// this example is just showing the usage of "LicenseProfiles_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeLicenseProfileResource created on azure
// for more information of creating HybridComputeLicenseProfileResource, please refer to the document of HybridComputeLicenseProfileResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
ResourceIdentifier hybridComputeLicenseProfileResourceId = HybridComputeLicenseProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName);
HybridComputeLicenseProfileResource hybridComputeLicenseProfile = client.GetHybridComputeLicenseProfileResource(hybridComputeLicenseProfileResourceId);

// invoke the operation
HybridComputeLicenseProfilePatch patch = new HybridComputeLicenseProfilePatch()
{
    SubscriptionStatus = LicenseProfileSubscriptionStatusUpdate.Enable,
    ProductType = LicenseProfileProductType.WindowsServer,
    ProductFeatures =
    {
    new HybridComputeProductFeatureUpdate()
    {
    Name = "Hotpatch",
    SubscriptionStatus = LicenseProfileSubscriptionStatusUpdate.Enable,
    }
    },
    AssignedLicense = "{LicenseResourceId}",
    SoftwareAssuranceCustomer = true,
};
ArmOperation<HybridComputeLicenseProfileResource> lro = await hybridComputeLicenseProfile.UpdateAsync(WaitUntil.Completed, patch);
HybridComputeLicenseProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputeLicenseProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
