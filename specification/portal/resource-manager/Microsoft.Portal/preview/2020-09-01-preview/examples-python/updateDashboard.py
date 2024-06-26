from azure.identity import DefaultAzureCredential
from azure.mgmt.portal import Portal

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-portal
# USAGE
    python update_a_dashboard.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Portal(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.dashboards.update(
        resource_group_name="testRG",
        dashboard_name="testDashboard",
        dashboard={"tags": {"aKey": "bValue", "anotherKey": "anotherValue2"}},
    )
    print(response)


# x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/updateDashboard.json
if __name__ == "__main__":
    main()
