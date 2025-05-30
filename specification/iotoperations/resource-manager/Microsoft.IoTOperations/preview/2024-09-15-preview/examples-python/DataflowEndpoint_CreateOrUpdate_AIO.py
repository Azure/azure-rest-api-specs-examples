from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python dataflow_endpoint_create_or_update_aio.py

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

    response = client.dataflow_endpoint.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="resource-name123",
        dataflow_endpoint_name="aio-builtin-broker-endpoint",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "properties": {
                "endpointType": "Mqtt",
                "mqttSettings": {
                    "authentication": {
                        "method": "Kubernetes",
                        "serviceAccountTokenSettings": {"audience": "aio-internal"},
                    },
                    "host": "aio-broker:18883",
                    "tls": {"mode": "Enabled", "trustedCaCertificateConfigMapRef": "aio-ca-trust-bundle-test-only"},
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_AIO.json
if __name__ == "__main__":
    main()
