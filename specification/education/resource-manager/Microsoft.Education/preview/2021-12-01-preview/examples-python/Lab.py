from azure.identity import DefaultAzureCredential
from azure.mgmt.education import EducationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-education
# USAGE
    python lab.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EducationManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.labs.get(
        billing_account_name="{billingAccountName}",
        billing_profile_name="{billingProfileName}",
        invoice_section_name="{invoiceSectionName}",
    )
    print(response)


# x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/Lab.json
if __name__ == "__main__":
    main()
