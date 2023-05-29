using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Datastores_Get.json
// this example is just showing the usage of "Datastores_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudClusterResource created on azure
// for more information of creating AvsPrivateCloudClusterResource, please refer to the document of AvsPrivateCloudClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string clusterName = "cluster1";
ResourceIdentifier avsPrivateCloudClusterResourceId = AvsPrivateCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, clusterName);
AvsPrivateCloudClusterResource avsPrivateCloudCluster = client.GetAvsPrivateCloudClusterResource(avsPrivateCloudClusterResourceId);

// get the collection of this AvsPrivateCloudDatastoreResource
AvsPrivateCloudDatastoreCollection collection = avsPrivateCloudCluster.GetAvsPrivateCloudDatastores();

// invoke the operation
string datastoreName = "datastore1";
bool result = await collection.ExistsAsync(datastoreName);

Console.WriteLine($"Succeeded: {result}");
