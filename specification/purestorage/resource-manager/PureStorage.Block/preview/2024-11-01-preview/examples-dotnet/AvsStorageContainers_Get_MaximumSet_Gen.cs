using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PureStorageBlock;

// Generated from example definition: 2024-11-01-preview/AvsStorageContainers_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AvsStorageContainer_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PureStorageAvsStorageContainerResource created on azure
// for more information of creating PureStorageAvsStorageContainerResource, please refer to the document of PureStorageAvsStorageContainerResource
string subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
string resourceGroupName = "rgpurestorage";
string storagePoolName = "storagePoolName";
string storageContainerName = "storageContainerName";
ResourceIdentifier pureStorageAvsStorageContainerResourceId = PureStorageAvsStorageContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storagePoolName, storageContainerName);
PureStorageAvsStorageContainerResource pureStorageAvsStorageContainer = client.GetPureStorageAvsStorageContainerResource(pureStorageAvsStorageContainerResourceId);

// invoke the operation
PureStorageAvsStorageContainerResource result = await pureStorageAvsStorageContainer.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PureStorageAvsStorageContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
