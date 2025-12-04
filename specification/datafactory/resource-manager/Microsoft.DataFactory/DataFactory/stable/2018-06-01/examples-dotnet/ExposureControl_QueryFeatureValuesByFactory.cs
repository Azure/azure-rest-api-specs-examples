using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/ExposureControl_QueryFeatureValuesByFactory.json
// this example is just showing the usage of "ExposureControl_QueryFeatureValuesByFactory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryResource created on azure
// for more information of creating DataFactoryResource, please refer to the document of DataFactoryResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
ResourceIdentifier dataFactoryResourceId = DataFactoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName);
DataFactoryResource dataFactory = client.GetDataFactoryResource(dataFactoryResourceId);

// invoke the operation
ExposureControlBatchContent content = new ExposureControlBatchContent(new ExposureControlContent[]
{
new ExposureControlContent
{
FeatureName = "ADFIntegrationRuntimeSharingRbac",
FeatureType = "Feature",
},
new ExposureControlContent
{
FeatureName = "ADFSampleFeature",
FeatureType = "Feature",
}
});
ExposureControlBatchResult result = await dataFactory.GetExposureControlFeaturesAsync(content);

Console.WriteLine($"Succeeded: {result}");
