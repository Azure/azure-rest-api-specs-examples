using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Media.Models;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/streamingendpoint-create.json
// this example is just showing the usage of "StreamingEndpoints_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
string resourceGroupName = "mediaresources";
string accountName = "slitestmedia10";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// get the collection of this StreamingEndpointResource
StreamingEndpointCollection collection = mediaServicesAccount.GetStreamingEndpoints();

// invoke the operation
string streamingEndpointName = "myStreamingEndpoint1";
StreamingEndpointData data = new StreamingEndpointData(new AzureLocation("West US"))
{
    Description = "test event 1",
    ScaleUnits = 1,
    AvailabilitySetName = "availableset",
    AccessControl = new StreamingEndpointAccessControl
    {
        AkamaiSignatureHeaderAuthenticationKeyList = {new AkamaiSignatureHeaderAuthenticationKey
        {
        Identifier = "id1",
        Base64Key = "dGVzdGlkMQ==",
        ExpireOn = DateTimeOffset.Parse("2029-12-31T16:00:00-08:00"),
        }, new AkamaiSignatureHeaderAuthenticationKey
        {
        Identifier = "id2",
        Base64Key = "dGVzdGlkMQ==",
        ExpireOn = DateTimeOffset.Parse("2030-12-31T16:00:00-08:00"),
        }},
        AllowedIPs = {new IPRange
        {
        Name = "AllowedIp",
        Address = IPAddress.Parse("192.168.1.1"),
        }},
    },
    IsCdnEnabled = false,
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<StreamingEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, streamingEndpointName, data);
StreamingEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StreamingEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
