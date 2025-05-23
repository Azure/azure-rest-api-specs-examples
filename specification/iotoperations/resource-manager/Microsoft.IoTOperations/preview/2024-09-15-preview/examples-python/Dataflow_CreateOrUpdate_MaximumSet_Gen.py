from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python dataflow_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTOperationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.dataflow.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="resource-name123",
        dataflow_profile_name="resource-name123",
        dataflow_name="resource-name123",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "properties": {
                "mode": "Enabled",
                "operations": [
                    {
                        "builtInTransformationSettings": {
                            "datasets": [
                                {
                                    "description": "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                                    "expression": "aatbwomvflemsxialv",
                                    "inputs": ["mosffpsslifkq"],
                                    "key": "qsfqcgxaxnhfumrsdsokwyv",
                                    "schemaRef": "n",
                                }
                            ],
                            "filter": [
                                {
                                    "description": "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                                    "expression": "n",
                                    "inputs": ["sxmjkbntgb"],
                                    "type": "Filter",
                                }
                            ],
                            "map": [
                                {
                                    "description": "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                                    "expression": "txoiltogsarwkzalsphvlmt",
                                    "inputs": ["xsbxuk"],
                                    "output": "nvgtmkfl",
                                    "type": "NewProperties",
                                }
                            ],
                            "schemaRef": "mcdc",
                            "serializationFormat": "Delta",
                        },
                        "destinationSettings": {
                            "dataDestination": "cbrh",
                            "endpointRef": "kybkchnzimerguekuvqlqiqdvvrt",
                        },
                        "name": "knnafvkwoeakm",
                        "operationType": "Source",
                        "sourceSettings": {
                            "assetRef": "zayyykwmckaocywdkohmu",
                            "dataSources": ["chkkpymxhp"],
                            "endpointRef": "iixotodhvhkkfcfyrkoveslqig",
                            "schemaRef": "pknmdzqll",
                            "serializationFormat": "Json",
                        },
                    }
                ],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
