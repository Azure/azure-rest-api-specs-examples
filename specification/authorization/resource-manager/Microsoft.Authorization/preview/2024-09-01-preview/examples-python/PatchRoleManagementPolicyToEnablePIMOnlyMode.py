from azure.identity import DefaultAzureCredential

from azure.mgmt.authorization import AuthorizationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-authorization
# USAGE
    python patch_role_management_policy_to_enable_pim_only_mode.py

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

    response = client.role_management_policies.update(
        scope="providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
        role_management_policy_name="570c3619-7688-4b34-b290-2b8bb3ccab2a",
        parameters={
            "properties": {
                "rules": [
                    {
                        "id": "PIMOnlyMode_Admin_Assignment",
                        "pimOnlyModeSettings": {
                            "excludedAssignmentTypes": ["ServicePrincipalsAsTarget"],
                            "excludes": [
                                {"id": "ec42a424-a0c0-4418-8788-d19bdeb03704", "type": "User"},
                                {"id": "00029dfb-0218-4e7a-9a85-c15dc0c880bc", "type": "Group"},
                                {"id": "0000103d-1fc2-4ac8-81de-71517765655c", "type": "ServicePrincipal"},
                            ],
                            "mode": "Enabled",
                        },
                        "ruleType": "RoleManagementPolicyPimOnlyModeRule",
                        "target": {
                            "caller": "Admin",
                            "enforcedSettings": ["all"],
                            "inheritableSettings": ["all"],
                            "level": "Assignment",
                            "operations": ["all"],
                            "targetObjects": [],
                        },
                    }
                ]
            }
        },
    )
    print(response)


# x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2024-09-01-preview/examples/PatchRoleManagementPolicyToEnablePIMOnlyMode.json
if __name__ == "__main__":
    main()
