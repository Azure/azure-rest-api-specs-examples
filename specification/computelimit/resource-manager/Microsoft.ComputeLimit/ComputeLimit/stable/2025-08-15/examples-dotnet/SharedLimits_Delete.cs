using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeLimit.Models;
using Azure.ResourceManager.ComputeLimit;

// Generated from example definition: 2025-08-15/SharedLimits_Delete.json
// this example is just showing the usage of "SharedLimit_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ComputeLimitSharedLimitResource created on azure
// for more information of creating ComputeLimitSharedLimitResource, please refer to the document of ComputeLimitSharedLimitResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
AzureLocation location = new AzureLocation("eastus");
string name = "StandardDSv3Family";
ResourceIdentifier computeLimitSharedLimitResourceId = ComputeLimitSharedLimitResource.CreateResourceIdentifier(subscriptionId, location, name);
ComputeLimitSharedLimitResource computeLimitSharedLimit = client.GetComputeLimitSharedLimitResource(computeLimitSharedLimitResourceId);

// invoke the operation
await computeLimitSharedLimit.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
