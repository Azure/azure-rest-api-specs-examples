from azure.identity import DefaultAzureCredential
from azure.mgmt.recoveryservicesdatareplication import RecoveryServicesDataReplicationMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicesdatareplication
# USAGE
    python policy_operation_status_get.py

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

    response = client.policy_operation_status.get(
        resource_group_name="rgrecoveryservicesdatareplication",
        vault_name="4",
        policy_name="xczxcwec",
        operation_id="wdqfsdxv",
    )
    print(response)


# x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/PolicyOperationStatus_Get.json
if __name__ == "__main__":
    main()
