from azure.identity import DefaultAzureCredential
from azure.mgmt.sqlvirtualmachine import SqlVirtualMachineManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sqlvirtualmachine
# USAGE
    python troubleshoot_sql_virtual_machine.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlVirtualMachineManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.sql_virtual_machine_troubleshoot.begin_troubleshoot(
        resource_group_name="testrg",
        sql_virtual_machine_name="testvm",
        parameters={
            "endTimeUtc": "2022-07-09T22:10:00Z",
            "properties": {"unhealthyReplicaInfo": {"availabilityGroupName": "AG1"}},
            "startTimeUtc": "2022-07-09T17:10:00Z",
            "troubleshootingScenario": "UnhealthyReplica",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/TroubleshootSqlVirtualMachine.json
if __name__ == "__main__":
    main()
