from azure.identity import DefaultAzureCredential

from azure.mgmt.datafactory import DataFactoryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datafactory
# USAGE
    python credentials_create.py

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

    response = client.credential_operations.create_or_update(
        resource_group_name="exampleResourceGroup",
        factory_name="exampleFactoryName",
        credential_name="exampleCredential",
        credential={
            "properties": {
                "type": "ManagedIdentity",
                "typeProperties": {
                    "resourceId": "/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Create.json
if __name__ == "__main__":
    main()
