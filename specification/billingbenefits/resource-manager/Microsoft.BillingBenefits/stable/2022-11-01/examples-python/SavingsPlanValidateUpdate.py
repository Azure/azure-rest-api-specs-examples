from azure.identity import DefaultAzureCredential
from azure.mgmt.billingbenefits import BillingBenefitsRP

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billingbenefits
# USAGE
    python savings_plan_validate_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingBenefitsRP(
        credential=DefaultAzureCredential(),
    )

    response = client.savings_plan.validate_update(
        savings_plan_order_id="20000000-0000-0000-0000-000000000000",
        savings_plan_id="30000000-0000-0000-0000-000000000000",
        body={
            "benefits": [
                {
                    "appliedScopeProperties": {
                        "managementGroupId": "/providers/Microsoft.Management/managementGroups/30000000-0000-0000-0000-000000000100",
                        "tenantId": "30000000-0000-0000-0000-000000000100",
                    },
                    "appliedScopeType": "ManagementGroup",
                },
                {
                    "appliedScopeProperties": {
                        "managementGroupId": "/providers/Microsoft.Management/managementGroups/MockMG",
                        "tenantId": "30000000-0000-0000-0000-000000000100",
                    },
                    "appliedScopeType": "ManagementGroup",
                },
            ]
        },
    )
    print(response)


# x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanValidateUpdate.json
if __name__ == "__main__":
    main()
