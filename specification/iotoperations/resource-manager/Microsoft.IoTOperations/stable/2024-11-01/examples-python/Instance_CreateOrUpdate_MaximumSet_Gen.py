from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python instance_create_or_update_maximum_set_gen.py

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

    response = client.instance.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="aio-instance",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "identity": {"type": "None", "userAssignedIdentities": {}},
            "location": "xvewadyhycrjpu",
            "properties": {
                "description": "kpqtgocs",
                "schemaRegistryRef": {
                    "resourceId": "/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-11-01/Instance_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
