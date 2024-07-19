using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/license/License_Delete.json
// this example is just showing the usage of "Licenses_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeLicenseResource created on azure
// for more information of creating HybridComputeLicenseResource, please refer to the document of HybridComputeLicenseResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string licenseName = "{licenseName}";
ResourceIdentifier hybridComputeLicenseResourceId = HybridComputeLicenseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, licenseName);
HybridComputeLicenseResource hybridComputeLicense = client.GetHybridComputeLicenseResource(hybridComputeLicenseResourceId);

// invoke the operation
await hybridComputeLicense.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
