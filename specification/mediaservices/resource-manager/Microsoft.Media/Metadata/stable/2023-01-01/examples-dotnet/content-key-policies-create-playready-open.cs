using Azure;
using Azure.ResourceManager;
using System;
using System.Text;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Media.Models;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2023-01-01/examples/content-key-policies-create-playready-open.json
// this example is just showing the usage of "ContentKeyPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosorg";
string accountName = "contosomedia";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// get the collection of this ContentKeyPolicyResource
ContentKeyPolicyCollection collection = mediaServicesAccount.GetContentKeyPolicies();

// invoke the operation
string contentKeyPolicyName = "PolicyWithPlayReadyOptionAndOpenRestriction";
ContentKeyPolicyData data = new ContentKeyPolicyData
{
    Description = "ArmPolicyDescription",
    Options = {new ContentKeyPolicyOption(new ContentKeyPolicyPlayReadyConfiguration(new ContentKeyPolicyPlayReadyLicense[]
    {
    new ContentKeyPolicyPlayReadyLicense(true, ContentKeyPolicyPlayReadyLicenseType.Persistent, new ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader(), ContentKeyPolicyPlayReadyContentType.UltraVioletDownload)
    {
    SecurityLevel = PlayReadySecurityLevel.SL150,
    BeginOn = DateTimeOffset.Parse("2017-10-16T18:22:53.46Z"),
    PlayRight = new ContentKeyPolicyPlayReadyPlayRight(false, true, false, ContentKeyPolicyPlayReadyUnknownOutputPassingOption.NotAllowed)
    {
    ScmsRestriction = 2,
    },
    }
    }), new ContentKeyPolicyOpenRestriction())
    {
    Name = "ArmPolicyOptionName",
    }},
};
ArmOperation<ContentKeyPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, contentKeyPolicyName, data);
ContentKeyPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContentKeyPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
