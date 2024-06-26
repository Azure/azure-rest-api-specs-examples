from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python post_get_insights.py

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

    response = client.entities.get_insights(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        entity_id="e1d3d618-e11f-478b-98e3-bb381539a8e1",
        parameters={
            "addDefaultExtendedTimeRange": False,
            "endTime": "2021-10-01T00:00:00.000Z",
            "insightQueryIds": ["cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4"],
            "startTime": "2021-09-01T00:00:00.000Z",
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/entities/insights/PostGetInsights.json
if __name__ == "__main__":
    main()
