from azure.identity import DefaultAzureCredential
from azure.mgmt.billing import BillingManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-billing
# USAGE
    python put_instruction.py

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

    response = client.instructions.put(
        billing_account_name="{billingAccountName}",
        billing_profile_name="{billingProfileName}",
        instruction_name="{instructionName}",
        parameters={
            "properties": {
                "amount": 5000,
                "endDate": "2020-12-30T21:26:47.997Z",
                "startDate": "2019-12-30T21:26:47.997Z",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInstruction.json
if __name__ == "__main__":
    main()
