from azure.identity import DefaultAzureCredential

from azure.mgmt.appconfiguration import AppConfigurationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appconfiguration
# USAGE
    python configuration_stores_create_replica.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AppConfigurationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.replicas.begin_create(
        resource_group_name="myResourceGroup",
        config_store_name="contoso",
        replica_name="myReplicaEus",
        replica_creation_parameters={"location": "eastus"},
    ).result()
    print(response)


# x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateReplica.json
if __name__ == "__main__":
    main()
