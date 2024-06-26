from azure.identity import DefaultAzureCredential
from azure.mgmt.billingbenefits import BillingBenefitsRP

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billingbenefits
# USAGE
    python savings_plan_order_alias_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingBenefitsRP(
        credential=DefaultAzureCredential(),
    )

    response = client.savings_plan_order_alias.begin_create(
        savings_plan_order_alias_name="spAlias123",
        body={
            "properties": {
                "appliedScopeProperties": None,
                "appliedScopeType": "Shared",
                "billingPlan": "P1M",
                "billingScopeId": "/subscriptions/30000000-0000-0000-0000-000000000000",
                "commitment": {"amount": 0.001, "currencyCode": "USD", "grain": "Hourly"},
                "displayName": "Compute_SavingsPlan_10-28-2022_16-38",
                "term": "P3Y",
            },
            "sku": {"name": "Compute_Savings_Plan"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderAliasCreate.json
if __name__ == "__main__":
    main()
