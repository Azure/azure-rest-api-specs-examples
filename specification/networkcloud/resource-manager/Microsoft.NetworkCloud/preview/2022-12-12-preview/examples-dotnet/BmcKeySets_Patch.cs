using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_Patch.json
// this example is just showing the usage of "BmcKeySets_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BmcKeySetResource created on azure
// for more information of creating BmcKeySetResource, please refer to the document of BmcKeySetResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
string bmcKeySetName = "bmcKeySetName";
ResourceIdentifier bmcKeySetResourceId = BmcKeySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, bmcKeySetName);
BmcKeySetResource bmcKeySet = client.GetBmcKeySetResource(bmcKeySetResourceId);

// invoke the operation
BmcKeySetPatch patch = new BmcKeySetPatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
    Expiration = DateTimeOffset.Parse("2022-12-31T23:59:59.008Z"),
    UserList =
    {
    new KeySetUser("userABC",new SshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"))
    {
    Description = "Needs access for troubleshooting as a part of the support team",
    },new KeySetUser("userXYZ",new SshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"))
    {
    Description = "Needs access for troubleshooting as a part of the support team",
    }
    },
};
ArmOperation<BmcKeySetResource> lro = await bmcKeySet.UpdateAsync(WaitUntil.Completed, patch);
BmcKeySetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BmcKeySetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
