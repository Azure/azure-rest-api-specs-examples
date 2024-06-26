from azure.identity import DefaultAzureCredential
from azure.mgmt.billingbenefits import BillingBenefitsRP

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billingbenefits
# USAGE
    python savings_plan_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingBenefitsRP(
        credential=DefaultAzureCredential(),
    )

    response = client.savings_plan.update(
        savings_plan_order_id="20000000-0000-0000-0000-000000000000",
        savings_plan_id="30000000-0000-0000-0000-000000000000",
        body={
            "properties": {
                "appliedScopeProperties": {
                    "resourceGroupId": "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"
                },
                "appliedScopeType": "Single",
                "displayName": "TestDisplayName",
                "renew": True,
                "renewProperties": {
                    "purchaseProperties": {
                        "properties": {
                            "appliedScopeProperties": {
                                "resourceGroupId": "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"
                            },
                            "appliedScopeType": "Single",
                            "billingPlan": "P1M",
                            "billingScopeId": "/subscriptions/10000000-0000-0000-0000-000000000000",
                            "commitment": {"amount": 15.23, "currencyCode": "USD", "grain": "Hourly"},
                            "displayName": "TestDisplayName_renewed",
                            "renew": False,
                            "term": "P1Y",
                        },
                        "sku": {"name": "Compute_Savings_Plan"},
                    }
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanUpdate.json
if __name__ == "__main__":
    main()
