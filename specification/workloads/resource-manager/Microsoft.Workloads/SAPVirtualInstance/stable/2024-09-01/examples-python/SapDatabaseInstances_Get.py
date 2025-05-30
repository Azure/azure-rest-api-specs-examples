from azure.identity import DefaultAzureCredential

from azure.mgmt.workloadssapvirtualinstance import WorkloadsSapVirtualInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadssapvirtualinstance
# USAGE
    python sap_database_instances_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsSapVirtualInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.sap_database_instances.get(
        resource_group_name="test-rg",
        sap_virtual_instance_name="X00",
        database_instance_name="databaseServer",
    )
    print(response)


# x-ms-original-file: 2024-09-01/SapDatabaseInstances_Get.json
if __name__ == "__main__":
    main()
