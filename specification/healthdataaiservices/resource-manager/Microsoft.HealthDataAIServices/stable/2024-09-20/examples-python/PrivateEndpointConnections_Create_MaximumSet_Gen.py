from azure.identity import DefaultAzureCredential

from azure.mgmt.healthdataaiservices import HealthDataAIServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthdataaiservices
# USAGE
    python private_endpoint_connections_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HealthDataAIServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.private_endpoint_connections.begin_create(
        resource_group_name="rgopenapi",
        deid_service_name="deidTest",
        private_endpoint_connection_name="kgwgrrpabvrsrrvpcgcnfmyfgyrl",
        resource={
            "properties": {
                "privateEndpoint": {},
                "privateLinkServiceConnectionState": {
                    "actionsRequired": "ulb",
                    "description": "xr",
                    "status": "Pending",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-09-20/PrivateEndpointConnections_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
