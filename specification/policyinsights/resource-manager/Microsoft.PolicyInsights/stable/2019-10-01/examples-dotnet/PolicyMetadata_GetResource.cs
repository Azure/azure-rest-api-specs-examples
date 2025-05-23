using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_GetResource.json
// this example is just showing the usage of "PolicyMetadata_GetResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PolicyMetadataResource created on azure
// for more information of creating PolicyMetadataResource, please refer to the document of PolicyMetadataResource
string resourceName = "NIST_SP_800-53_R4_AC-2";
ResourceIdentifier policyMetadataResourceId = PolicyMetadataResource.CreateResourceIdentifier(resourceName);
PolicyMetadataResource policyMetadata = client.GetPolicyMetadataResource(policyMetadataResourceId);

// invoke the operation
PolicyMetadataResource result = await policyMetadata.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyMetadataData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
