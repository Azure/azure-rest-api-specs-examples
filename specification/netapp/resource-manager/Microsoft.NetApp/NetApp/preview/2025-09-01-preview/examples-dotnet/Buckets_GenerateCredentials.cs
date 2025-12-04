using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/Buckets_GenerateCredentials.json
// this example is just showing the usage of "Buckets_GenerateCredentials" operation, for the dependent resources, they will have to be created separately.

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
NetAppBucketCredentialsExpiry body = new NetAppBucketCredentialsExpiry
{
    KeyPairExpiryDays = 3,
};
NetAppBucketGenerateCredentials result = await netAppBucket.GenerateCredentialsAsync(body);

Console.WriteLine($"Succeeded: {result}");
