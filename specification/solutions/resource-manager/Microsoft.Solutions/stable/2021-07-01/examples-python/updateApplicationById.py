from azure.identity import DefaultAzureCredential
from azure.mgmt.managedapplications import ManagedApplicationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedapplications
# USAGE
    python update_application_by_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedApplicationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.applications.begin_update_by_id(
        application_id="subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication",
    ).result()
    print(response)


# x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplicationById.json
if __name__ == "__main__":
    main()
