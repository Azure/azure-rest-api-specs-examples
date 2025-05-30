using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/SACPut.json
// this example is just showing the usage of "StorageAccountCredentials_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeDeviceResource created on azure
// for more information of creating DataBoxEdgeDeviceResource, please refer to the document of DataBoxEdgeDeviceResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
ResourceIdentifier dataBoxEdgeDeviceResourceId = DataBoxEdgeDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName);
DataBoxEdgeDeviceResource dataBoxEdgeDevice = client.GetDataBoxEdgeDeviceResource(dataBoxEdgeDeviceResourceId);

// get the collection of this DataBoxEdgeStorageAccountCredentialResource
DataBoxEdgeStorageAccountCredentialCollection collection = dataBoxEdgeDevice.GetDataBoxEdgeStorageAccountCredentials();

// invoke the operation
string name = "sac1";
DataBoxEdgeStorageAccountCredentialData data = new DataBoxEdgeStorageAccountCredentialData("sac1", DataBoxEdgeStorageAccountSslStatus.Disabled, DataBoxEdgeStorageAccountType.BlobStorage)
{
    UserName = "cisbvt",
    AccountKey = new AsymmetricEncryptedSecret("lAeZEYi6rNP1/EyNaVUYmTSZEYyaIaWmwUsGwek0+xiZj54GM9Ue9/UA2ed/ClC03wuSit2XzM/cLRU5eYiFBwks23rGwiQOr3sruEL2a74EjPD050xYjA6M1I2hu/w2yjVHhn5j+DbXS4Xzi+rHHNZK3DgfDO3PkbECjPck+PbpSBjy9+6Mrjcld5DIZhUAeMlMHrFlg+WKRKB14o/og56u5/xX6WKlrMLEQ+y6E18dUwvWs2elTNoVO8PBE8SM/CfooX4AMNvaNdSObNBPdP+F6Lzc556nFNWXrBLRt0vC7s9qTiVRO4x/qCNaK/B4y7IqXMllwQFf4Np9UQ2ECA==", DataBoxEdgeEncryptionAlgorithm.Aes256)
    {
        EncryptionCertThumbprint = "2A9D8D6BE51574B5461230AEF02F162C5F01AD31",
    },
};
ArmOperation<DataBoxEdgeStorageAccountCredentialResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DataBoxEdgeStorageAccountCredentialResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeStorageAccountCredentialData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
