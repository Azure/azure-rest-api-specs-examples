from azure.identity import DefaultAzureCredential

from azure.mgmt.authorization import AuthorizationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-authorization
# USAGE
    python delete_access_review_history_definition.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AuthorizationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    client.scope_access_review_history_definition.delete_by_id(
        scope="subscriptions/fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
        history_definition_id="fa73e90b-5bf1-45fd-a182-35ce5fc0674d",
    )


# x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/DeleteAccessReviewHistoryDefinition.json
if __name__ == "__main__":
    main()
