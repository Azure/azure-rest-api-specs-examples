from azure.identity import DefaultAzureCredential

from azure.mgmt.impactreporting import ImpactReportingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-impactreporting
# USAGE
    python insights_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ImpactReportingMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.insights.delete(
        workload_impact_name="impactid22",
        insight_name="insightId12",
    )


# x-ms-original-file: 2024-05-01-preview/Insights_Delete.json
if __name__ == "__main__":
    main()
