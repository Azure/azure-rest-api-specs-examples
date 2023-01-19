using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
// this example is just showing the usage of "DataFlowDebugSession_AddDataFlow" operation, for the dependent resources, they will have to be created separately.

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
FactoryDataFlowDebugPackageContent content = new FactoryDataFlowDebugPackageContent()
{
    SessionId = Guid.Parse("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
    DataFlow = new FactoryDataFlowDebugInfo(new FactoryMappingDataFlowDefinition()
    {
        Sources =
        {
        new DataFlowSource("source1")
        {
        Dataset = new DatasetReference(DatasetReferenceType.DatasetReference,"DelimitedText2"),
        }
        },
        Sinks =
        {
        },
        Transformations =
        {
        },
        Script = "\n\nsource(output(\n\t\tColumn_1 as string\n\t),\n\tallowSchemaDrift: true,\n\tvalidateSchema: false) ~> source1",
    })
    {
        Name = "dataflow1",
    },
    Datasets =
    {
    new FactoryDatasetDebugInfo(new DelimitedTextDataset(new FactoryLinkedServiceReference(FactoryLinkedServiceReferenceType.LinkedServiceReference,"linkedService5"))
    {
    DataLocation = new AzureBlobStorageLocation()
    {
    Container = BinaryData.FromString("dataflow-sample-data"),
    FileName = BinaryData.FromString("Ansiencoding.csv"),
    },
    ColumnDelimiter = BinaryData.FromString(","),
    QuoteChar = BinaryData.FromString("\""),
    EscapeChar = BinaryData.FromString("\\"),
    FirstRowAsHeader = BinaryData.FromString("true"),
    Schema = BinaryData.FromObjectAsJson(new object[] { new Dictionary<string, object>()
    {
    ["type"] = "String"} }),
    Annotations =
    {
    },
    })
    {
    Name = "dataset1",
    }
    },
    LinkedServices =
    {
    new FactoryLinkedServiceDebugInfo(new AzureBlobStorageLinkedService()
    {
    ConnectionString = BinaryData.FromString("DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;"),
    EncryptedCredential = BinaryData.FromString("<credential>"),
    Annotations =
    {
    },
    })
    {
    Name = "linkedService1",
    }
    },
    DebugSettings = new DataFlowDebugPackageDebugSettings()
    {
        SourceSettings =
        {
        new DataFlowSourceSetting()
        {
        SourceName = "source1",
        RowLimit = 1000,
        },new DataFlowSourceSetting()
        {
        SourceName = "source2",
        RowLimit = 222,
        }
        },
        Parameters =
        {
        ["sourcePath"] = BinaryData.FromString("Toy"),
        },
        DatasetParameters = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
            ["Movies"] = new Dictionary<string, object>()
            {
                ["path"] = "abc"
            },
            ["Output"] = new Dictionary<string, object>()
            {
                ["time"] = "def"
            }
        }),
    },
};
FactoryDataFlowStartDebugSessionResult result = await dataFactory.AddDataFlowToDebugSessionAsync(content);

Console.WriteLine($"Succeeded: {result}");
