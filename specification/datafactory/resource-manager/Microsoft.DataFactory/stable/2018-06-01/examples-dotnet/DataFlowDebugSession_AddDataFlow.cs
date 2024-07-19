using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Expressions.DataFactory;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataFactory;

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
DataFactoryDataFlowDebugPackageContent content = new DataFactoryDataFlowDebugPackageContent()
{
    SessionId = Guid.Parse("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
    DataFlow = new DataFactoryDataFlowDebugInfo(new DataFactoryMappingDataFlowProperties()
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
    new DataFactoryDatasetDebugInfo(new DelimitedTextDataset(new DataFactoryLinkedServiceReference("LinkedServiceReference","linkedService5"))
    {
    DataLocation = new AzureBlobStorageLocation()
    {
    Container = "dataflow-sample-data",
    FileName = "Ansiencoding.csv",
    },
    ColumnDelimiter = ",",
    QuoteChar = "\"",
    EscapeChar = "\\",
    FirstRowAsHeader = true,
    Schema = new DatasetSchemaDataElement[]
    {
    new DatasetSchemaDataElement()
    {
    SchemaColumnType = "String",
    }
    },
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
    new DataFactoryLinkedServiceDebugInfo(new AzureBlobStorageLinkedService()
    {
    ConnectionString = "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
    EncryptedCredential = "<credential>",
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
        ["sourcePath"] = BinaryData.FromString("\"Toy\""),
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
DataFactoryDataFlowStartDebugSessionResult result = await dataFactory.AddDataFlowToDebugSessionAsync(content);

Console.WriteLine($"Succeeded: {result}");
