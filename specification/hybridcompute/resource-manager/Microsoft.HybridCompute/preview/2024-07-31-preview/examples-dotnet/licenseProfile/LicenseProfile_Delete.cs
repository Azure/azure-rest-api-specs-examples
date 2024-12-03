using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/licenseProfile/LicenseProfile_Delete.json
// this example is just showing the usage of "LicenseProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

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
await hybridComputeLicenseProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
