from azure.identity import DefaultAzureCredential

from azure.mgmt.impactreporting import ImpactReportingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-impactreporting
# USAGE
    python insights_create.py

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

    response = client.insights.create(
        workload_impact_name="impactid22",
        insight_name="insightId12",
        resource={
            "properties": {
                "category": "repair",
                "content": {
                    "description": 'At 2018-11-08T00:00:00Z UTC, your services dependent on these resources :code:`<link href=”…”>VM1</link>` may have experienced an issue. :code:`<br/>`\\ :code:`<div>`We have identified an outage that affected these resources(s). You can look at outage information on :code:`<link href="https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ">NL2W-VCZ</link>` link.\\ :code:`<div>`',
                    "title": "Impact Has been correlated to an outage",
                },
                "eventTime": "2023-06-15T04:00:00.009223Z",
                "impact": {
                    "impactId": "/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22",
                    "impactedResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername",
                    "startTime": "2023-06-15T01:00:00.009223Z",
                },
                "insightUniqueId": "00000000-0000-0000-0000-000000000000",
                "status": "resolved",
            }
        },
    )
    print(response)


# x-ms-original-file: 2024-05-01-preview/Insights_Create.json
if __name__ == "__main__":
    main()
