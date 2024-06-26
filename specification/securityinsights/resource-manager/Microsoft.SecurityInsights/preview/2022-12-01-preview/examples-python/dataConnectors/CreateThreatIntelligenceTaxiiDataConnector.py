from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_threat_intelligence_taxii_data_connector.py

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
            "etag": "d12423f6-a60b-4ca5-88c0-feb1a182d0f0",
            "kind": "ThreatIntelligenceTaxii",
            "properties": {
                "collectionId": "135",
                "dataTypes": {"taxiiClient": {"state": "Enabled"}},
                "friendlyName": "testTaxii",
                "password": "--",
                "pollingFrequency": "OnceADay",
                "taxiiLookbackPeriod": "2020-01-01T13:00:30.123Z",
                "taxiiServer": "https://limo.anomali.com/api/v1/taxii2/feeds",
                "tenantId": "06b3ccb8-1384-4bcc-aec7-852f6d57161b",
                "userName": "--",
                "workspaceId": "dd124572-4962-4495-9bd2-9dade12314b4",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/CreateThreatIntelligenceTaxiiDataConnector.json
if __name__ == "__main__":
    main()
