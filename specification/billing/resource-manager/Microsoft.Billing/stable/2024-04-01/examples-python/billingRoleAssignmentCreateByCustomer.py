from azure.identity import DefaultAzureCredential

from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python billing_role_assignment_create_by_customer.py

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

    response = client.billing_role_assignments.begin_create_by_customer(
        billing_account_name="00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30",
        billing_profile_name="BKM6-54VH-BG7-PGB",
        customer_name="703ab484-dda2-4402-827b-a74513b61e2d",
        parameters={
            "principalId": "00000000-0000-0000-0000-000000000000",
            "principalTenantId": "076915e7-de10-4323-bb34-a58c904068bb",
            "roleDefinitionId": "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfileName/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000",
            "userEmailAddress": "john@contoso.com",
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentCreateByCustomer.json
if __name__ == "__main__":
    main()
