from azure.identity import DefaultAzureCredential

from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python data_flow_debug_session_add_data_flow.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataFactoryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="12345678-1234-1234-1234-12345678abc",
    )

    response = client.data_flow_debug_session.add_data_flow(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        request={
            "dataFlow": {
                "name": "dataflow1",
                "properties": {
                    "type": "MappingDataFlow",
                    "typeProperties": {
                        "script": "\n\nsource(output(\n\t\tColumn_1 as string\n\t),\n\tallowSchemaDrift: true,\n\tvalidateSchema: false) ~> source1",
                        "sinks": [],
                        "sources": [
                            {
                                "dataset": {"referenceName": "DelimitedText2", "type": "DatasetReference"},
                                "name": "source1",
                            }
                        ],
                        "transformations": [],
                    },
                },
            },
            "datasets": [
                {
                    "name": "dataset1",
                    "properties": {
                        "annotations": [],
                        "linkedServiceName": {"referenceName": "linkedService5", "type": "LinkedServiceReference"},
                        "schema": [{"type": "String"}],
                        "type": "DelimitedText",
                        "typeProperties": {
                            "columnDelimiter": ",",
                            "escapeChar": "\\",
                            "firstRowAsHeader": True,
                            "location": {
                                "container": "dataflow-sample-data",
                                "fileName": "Ansiencoding.csv",
                                "type": "AzureBlobStorageLocation",
                            },
                            "quoteChar": '"',
                        },
                    },
                }
            ],
            "debugSettings": {
                "datasetParameters": {"Movies": {"path": "abc"}, "Output": {"time": "def"}},
                "parameters": {"sourcePath": "Toy"},
                "sourceSettings": [
                    {"rowLimit": 1000, "sourceName": "source1"},
                    {"rowLimit": 222, "sourceName": "source2"},
                ],
            },
            "linkedServices": [
                {
                    "name": "linkedService1",
                    "properties": {
                        "annotations": [],
                        "type": "AzureBlobStorage",
                        "typeProperties": {
                            "connectionString": "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
                            "encryptedCredential": "<credential>",
                        },
                    },
                }
            ],
            "sessionId": "f06ed247-9d07-49b2-b05e-2cb4a2fc871e",
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
if __name__ == "__main__":
    main()
