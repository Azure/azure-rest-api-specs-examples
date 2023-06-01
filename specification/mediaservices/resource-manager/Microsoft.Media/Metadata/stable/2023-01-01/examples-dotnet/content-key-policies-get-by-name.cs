using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2023-01-01/examples/content-key-policies-get-by-name.json
// this example is just showing the usage of "ContentKeyPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContentKeyPolicyResource created on azure
// for more information of creating ContentKeyPolicyResource, please refer to the document of ContentKeyPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosorg";
string accountName = "contosomedia";
string contentKeyPolicyName = "PolicyWithMultipleOptions";
ResourceIdentifier contentKeyPolicyResourceId = ContentKeyPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, contentKeyPolicyName);
ContentKeyPolicyResource contentKeyPolicy = client.GetContentKeyPolicyResource(contentKeyPolicyResourceId);

// invoke the operation
ContentKeyPolicyResource result = await contentKeyPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContentKeyPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
