from azure.identity import DefaultAzureCredential
from azure.mgmt.managedapplications import ManagedApplicationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedapplications
# USAGE
    python delete_application.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedApplicationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.applications.begin_delete(
        resource_group_name="rg",
        application_name="myManagedApplication",
    ).result()


# x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplication.json
if __name__ == "__main__":
    main()
