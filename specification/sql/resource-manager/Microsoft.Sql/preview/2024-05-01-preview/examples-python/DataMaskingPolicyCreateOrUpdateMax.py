from azure.identity import DefaultAzureCredential

from azure.mgmt.sql import SqlManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-sql
# USAGE
    python data_masking_policy_create_or_update_max.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SqlManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.data_masking_policies.create_or_update(
        resource_group_name="sqlcrudtest-6852",
        server_name="sqlcrudtest-2080",
        database_name="sqlcrudtest-331",
        data_masking_policy_name="Default",
        parameters={"properties": {"dataMaskingState": "Enabled", "exemptPrincipals": "testuser;"}},
    )
    print(response)


# x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/DataMaskingPolicyCreateOrUpdateMax.json
if __name__ == "__main__":
    main()
