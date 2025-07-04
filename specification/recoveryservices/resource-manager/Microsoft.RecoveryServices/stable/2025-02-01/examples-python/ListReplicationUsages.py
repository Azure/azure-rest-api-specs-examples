from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservices import RecoveryServicesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservices
# USAGE
    python list_replication_usages.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesClient(
        credential=DefaultAzureCredential(),
        subscription_id="6808dbbc-98c7-431f-a1b1-9580902423b7",
    )

    response = client.replication_usages.list(
        resource_group_name="avrai7517RG1",
        vault_name="avrai7517Vault1",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/ListReplicationUsages.json
if __name__ == "__main__":
    main()
