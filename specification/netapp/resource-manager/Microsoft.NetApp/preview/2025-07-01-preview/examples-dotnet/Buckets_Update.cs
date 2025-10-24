using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Buckets_Update.json
// this example is just showing the usage of "Buckets_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBucketResource created on azure
// for more information of creating NetAppBucketResource, please refer to the document of NetAppBucketResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string bucketName = "bucket1";
ResourceIdentifier netAppBucketResourceId = NetAppBucketResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, bucketName);
NetAppBucketResource netAppBucket = client.GetNetAppBucketResource(netAppBucketResourceId);

// invoke the operation
NetAppBucketPatch patch = new NetAppBucketPatch
{
    Server = new NetAppBucketServerPatchProperties
    {
        Fqdn = "fullyqualified.domainname.com",
        CertificateObject = "<REDACTED>",
    },
    Permissions = NetAppBucketPatchPermission.ReadWrite,
};
ArmOperation<NetAppBucketResource> lro = await netAppBucket.UpdateAsync(WaitUntil.Completed, patch);
NetAppBucketResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppBucketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
