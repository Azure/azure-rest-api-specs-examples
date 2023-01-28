using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-multiple-options.json
// this example is just showing the usage of "ContentKeyPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// get the collection of this ContentKeyPolicyResource
ContentKeyPolicyCollection collection = mediaServicesAccount.GetContentKeyPolicies();

// invoke the operation
string contentKeyPolicyName = "PolicyCreatedWithMultipleOptions";
ContentKeyPolicyData data = new ContentKeyPolicyData()
{
    Description = "ArmPolicyDescription",
    Options =
    {
    new ContentKeyPolicyOption(new ContentKeyPolicyClearKeyConfiguration(),new ContentKeyPolicyTokenRestriction("urn:issuer","urn:audience",new ContentKeyPolicySymmetricTokenKey(Convert.FromBase64String("AAAAAAAAAAAAAAAAAAAAAA==")),ContentKeyPolicyRestrictionTokenType.Swt))
    {
    Name = "ClearKeyOption",
    },new ContentKeyPolicyOption(new ContentKeyPolicyWidevineConfiguration("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),new ContentKeyPolicyOpenRestriction())
    {
    Name = "widevineoption",
    }
    },
};
ArmOperation<ContentKeyPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, contentKeyPolicyName, data);
ContentKeyPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContentKeyPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
