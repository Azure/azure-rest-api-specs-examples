using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Get.json
// this example is just showing the usage of "ChangeDataCapture_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryChangeDataCaptureResource created on azure
// for more information of creating DataFactoryChangeDataCaptureResource, please refer to the document of DataFactoryChangeDataCaptureResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string changeDataCaptureName = "exampleChangeDataCapture";
ResourceIdentifier dataFactoryChangeDataCaptureResourceId = DataFactoryChangeDataCaptureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, changeDataCaptureName);
DataFactoryChangeDataCaptureResource dataFactoryChangeDataCapture = client.GetDataFactoryChangeDataCaptureResource(dataFactoryChangeDataCaptureResourceId);

// invoke the operation
DataFactoryChangeDataCaptureResource result = await dataFactoryChangeDataCapture.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryChangeDataCaptureData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
