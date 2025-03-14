using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/ContainerRefresh.json
// this example is just showing the usage of "Containers_Refresh" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeStorageContainerResource created on azure
// for more information of creating DataBoxEdgeStorageContainerResource, please refer to the document of DataBoxEdgeStorageContainerResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string storageAccountName = "storageaccount1";
string containerName = "blobcontainer1";
ResourceIdentifier dataBoxEdgeStorageContainerResourceId = DataBoxEdgeStorageContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, storageAccountName, containerName);
DataBoxEdgeStorageContainerResource dataBoxEdgeStorageContainer = client.GetDataBoxEdgeStorageContainerResource(dataBoxEdgeStorageContainerResourceId);

// invoke the operation
await dataBoxEdgeStorageContainer.RefreshAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
