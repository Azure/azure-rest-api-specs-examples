using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/Datasets_Create.json
// this example is just showing the usage of "Datasets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryDatasetResource created on azure
// for more information of creating DataFactoryDatasetResource, please refer to the document of DataFactoryDatasetResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string datasetName = "exampleDataset";
ResourceIdentifier dataFactoryDatasetResourceId = DataFactoryDatasetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, datasetName);
DataFactoryDatasetResource dataFactoryDataset = client.GetDataFactoryDatasetResource(dataFactoryDatasetResourceId);

// invoke the operation
DataFactoryDatasetData data = new DataFactoryDatasetData(new AzureBlobDataset(new DataFactoryLinkedServiceReference(default, "exampleLinkedService"))
{
    FolderPath = null,
    FileName = null,
    Format = new DatasetTextFormat(),
    Parameters =
    {
    ["MyFileName"] = new EntityParameterSpecification(EntityParameterType.String),
    ["MyFolderPath"] = new EntityParameterSpecification(EntityParameterType.String)
    },
});
ArmOperation<DataFactoryDatasetResource> lro = await dataFactoryDataset.UpdateAsync(WaitUntil.Completed, data);
DataFactoryDatasetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryDatasetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
