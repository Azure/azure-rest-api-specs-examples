from azure.identity import DefaultAzureCredential

from azure.mgmt.batch import BatchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-batch
# USAGE
    python nsp_configuration_reconcile.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BatchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.network_security_perimeter.begin_reconcile_configuration(
        resource_group_name="default-azurebatch-japaneast",
        account_name="sampleacct",
        network_security_perimeter_configuration_name="00000000-0000-0000-0000-000000000000.sampleassociation",
    ).result()


# x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/NspConfigurationReconcile.json
if __name__ == "__main__":
    main()
