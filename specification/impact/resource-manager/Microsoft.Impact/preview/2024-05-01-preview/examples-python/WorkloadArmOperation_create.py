from azure.identity import DefaultAzureCredential

from azure.mgmt.impactreporting import ImpactReportingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-impactreporting
# USAGE
    python workload_arm_operation_create.py

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

    response = client.workload_impacts.begin_create(
        workload_impact_name="impact-002",
        resource={
            "properties": {
                "armCorrelationIds": ["00000000-0000-0000-0000-000000000000"],
                "clientIncidentDetails": {"clientIncidentId": "AA123", "clientIncidentSource": "Jira"},
                "endDateTime": None,
                "impactCategory": "ArmOperation",
                "impactDescription": "deletion of resource failed",
                "impactedResourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext",
                "startDateTime": "2022-06-15T05:59:46.6517821Z",
                "workload": {"context": "webapp/scenario1", "toolset": "Other"},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-05-01-preview/WorkloadArmOperation_create.json
if __name__ == "__main__":
    main()
