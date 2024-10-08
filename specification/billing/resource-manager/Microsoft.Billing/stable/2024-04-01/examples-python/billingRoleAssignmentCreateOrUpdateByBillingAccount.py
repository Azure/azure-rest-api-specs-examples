from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_role_assignment_create_or_update_by_billing_account.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.billing_role_assignments.begin_create_or_update_by_billing_account(
        billing_account_name="7898901",
        billing_role_assignment_name="9dfd08c2-62a3-4d47-85bd-1cdba1408402",
        parameters={
            "properties": {
                "principalId": "00000000-0000-0000-0000-000000000000",
                "principalTenantId": "076915e7-de10-4323-bb34-a58c904068bb",
                "roleDefinitionId": "/providers/Microsoft.Billing/billingAccounts/7898901/billingRoleDefinitions/9f1983cb-2574-400c-87e9-34cf8e2280db",
                "userEmailAddress": "john@contoso.com",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateOrUpdateByBillingAccount.json
if __name__ == "__main__":
    main()
