from azure.identity import DefaultAzureCredential
from azure.mgmt.recoveryservicesdatareplication import RecoveryServicesDataReplicationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesdatareplication
# USAGE
    python policy_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesDataReplicationMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="930CEC23-4430-4513-B855-DBA237E2F3BF",
    )

    client.policy.begin_delete(
        resource_group_name="rgrecoveryservicesdatareplication",
        vault_name="4",
        policy_name="wqfscsdv",
    ).result()


# x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Policy_Delete.json
if __name__ == "__main__":
    main()
