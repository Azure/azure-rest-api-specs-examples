from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_office_data_connetor.py

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

    response = client.data_connectors.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        data_connector_id="73e01a99-5cd7-4139-a149-9f2736ff2ab5",
        data_connector={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "kind": "Office365",
            "properties": {
                "dataTypes": {
                    "exchange": {"state": "Enabled"},
                    "sharePoint": {"state": "Enabled"},
                    "teams": {"state": "Enabled"},
                },
                "tenantId": "2070ecc9-b4d5-4ae4-adaa-936fa1954fa8",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/CreateOfficeDataConnetor.json
if __name__ == "__main__":
    main()
