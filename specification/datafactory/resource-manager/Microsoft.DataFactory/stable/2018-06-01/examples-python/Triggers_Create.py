from azure.identity import DefaultAzureCredential

from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python triggers_create.py

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

    response = client.triggers.create_or_update(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        trigger_name="exampleTrigger",
        trigger={
            "properties": {
                "pipelines": [
                    {
                        "parameters": {"OutputBlobNameList": ["exampleoutput.csv"]},
                        "pipelineReference": {"referenceName": "examplePipeline", "type": "PipelineReference"},
                    }
                ],
                "type": "ScheduleTrigger",
                "typeProperties": {
                    "recurrence": {
                        "endTime": "2018-06-16T00:55:13.8441801Z",
                        "frequency": "Minute",
                        "interval": 4,
                        "startTime": "2018-06-16T00:39:13.8441801Z",
                        "timeZone": "UTC",
                    }
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
if __name__ == "__main__":
    main()
