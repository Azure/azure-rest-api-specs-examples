from azure.identity import DefaultAzureCredential
from azure.mgmt.billingbenefits import BillingBenefitsRP

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billingbenefits
# USAGE
    python savings_plans_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BillingBenefitsRP(
        credential=DefaultAzureCredential(),
    )

    response = client.savings_plan.list_all()
    for item in response:
        print(item)


# x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
if __name__ == "__main__":
    main()
