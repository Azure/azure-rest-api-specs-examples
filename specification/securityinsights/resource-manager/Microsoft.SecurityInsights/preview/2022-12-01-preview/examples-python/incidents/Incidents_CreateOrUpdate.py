from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python incidents_create_or_update.py

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

    response = client.incidents.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        incident_id="73e01a99-5cd7-4139-a149-9f2736ff2ab5",
        incident={
            "etag": '"0300bf09-0000-0000-0000-5c37296e0000"',
            "properties": {
                "classification": "FalsePositive",
                "classificationComment": "Not a malicious activity",
                "classificationReason": "InaccurateData",
                "description": "This is a demo incident",
                "firstActivityTimeUtc": "2019-01-01T13:00:30Z",
                "lastActivityTimeUtc": "2019-01-01T13:05:30Z",
                "owner": {
                    "assignedTo": None,
                    "email": None,
                    "objectId": "2046feea-040d-4a46-9e2b-91c2941bfa70",
                    "ownerType": None,
                    "userPrincipalName": None,
                },
                "severity": "High",
                "status": "Closed",
                "title": "My incident",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/incidents/Incidents_CreateOrUpdate.json
if __name__ == "__main__":
    main()
