from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python get_io_tby_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.data_connectors.get(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        data_connector_id="d2e5dc7a-f3a2-429d-954b-939fa8c2932e",
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/GetIoTById.json
if __name__ == "__main__":
    main()
