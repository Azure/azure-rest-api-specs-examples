using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2023-01-01/examples/streaming-policies-create-secure-streaming.json
// this example is just showing the usage of "StreamingPolicies_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this StreamingPolicyResource
StreamingPolicyCollection collection = mediaServicesAccount.GetStreamingPolicies();

// invoke the operation
string streamingPolicyName = "UserCreatedSecureStreamingPolicy";
StreamingPolicyData data = new StreamingPolicyData()
{
    DefaultContentKeyPolicyName = "PolicyWithMultipleOptions",
    EnvelopeEncryption = new EnvelopeEncryption()
    {
        EnabledProtocols = new MediaEnabledProtocols(false, true, true, true),
        ContentKeys = new StreamingPolicyContentKeys()
        {
            DefaultKey = new EncryptionSchemeDefaultKey()
            {
                Label = "aesDefaultKey",
            },
        },
        CustomKeyAcquisitionUriTemplate = "https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}",
    },
    CommonEncryptionCenc = new CommonEncryptionCenc()
    {
        EnabledProtocols = new MediaEnabledProtocols(false, true, false, true),
        ClearTracks =
        {
        new MediaTrackSelection()
        {
        TrackSelections =
        {
        new TrackPropertyCondition(TrackPropertyType.FourCC,TrackPropertyCompareOperation.Equal)
        {
        Value = "hev1",
        }
        },
        }
        },
        ContentKeys = new StreamingPolicyContentKeys()
        {
            DefaultKey = new EncryptionSchemeDefaultKey()
            {
                Label = "cencDefaultKey",
            },
        },
        Drm = new CencDrmConfiguration()
        {
            PlayReady = new StreamingPolicyPlayReadyConfiguration()
            {
                CustomLicenseAcquisitionUriTemplate = "https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}",
                PlayReadyCustomAttributes = "PlayReady CustomAttributes",
            },
            WidevineCustomLicenseAcquisitionUriTemplate = "https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId",
        },
    },
    CommonEncryptionCbcs = new CommonEncryptionCbcs()
    {
        EnabledProtocols = new MediaEnabledProtocols(false, false, true, false),
        ContentKeys = new StreamingPolicyContentKeys()
        {
            DefaultKey = new EncryptionSchemeDefaultKey()
            {
                Label = "cbcsDefaultKey",
            },
        },
        Drm = new CbcsDrmConfiguration()
        {
            FairPlay = new StreamingPolicyFairPlayConfiguration(true)
            {
                CustomLicenseAcquisitionUriTemplate = "https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}",
            },
        },
    },
};
ArmOperation<StreamingPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, streamingPolicyName, data);
StreamingPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
