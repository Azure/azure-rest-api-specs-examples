from azure.identity import DefaultAzureCredential

from azure.mgmt.authorization import AuthorizationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-authorization
# USAGE
    python get_role_management_policy_by_scope.py

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

    response = client.role_management_policies.list_for_scope(
        scope="providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2024-09-01-preview/examples/GetRoleManagementPolicyByScope.json
if __name__ == "__main__":
    main()
