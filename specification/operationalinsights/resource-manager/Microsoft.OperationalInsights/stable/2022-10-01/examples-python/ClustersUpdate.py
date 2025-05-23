from azure.identity import DefaultAzureCredential

from azure.mgmt.loganalytics import LogAnalyticsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-loganalytics
# USAGE
    python clusters_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogAnalyticsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="53bc36c5-91e1-4d09-92c9-63b89e571926",
    )

    response = client.clusters.begin_update(
        resource_group_name="oiautorest6685",
        cluster_name="oiautorest6685",
        parameters={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/53bc36c5-91e1-4d09-92c9-63b89e571926/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {}
                },
            },
            "properties": {
                "keyVaultProperties": {
                    "keyName": "aztest2170cert",
                    "keyRsaSize": 1024,
                    "keyVaultUri": "https://aztest2170.vault.azure.net",
                    "keyVersion": "",
                }
            },
            "sku": {"capacity": 1000, "name": "CapacityReservation"},
            "tags": {"tag1": "val1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/ClustersUpdate.json
if __name__ == "__main__":
    main()
