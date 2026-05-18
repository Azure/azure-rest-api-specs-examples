from azure.identity import DefaultAzureCredential

from azure.mgmt.managementpartner import ACEProvisioningManagementPartnerAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managementpartner
# USAGE
    python get_partner_details_no_partner_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ACEProvisioningManagementPartnerAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.partners.get()
    print(response)


# x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/ManagementPartner/preview/2018-02-01/examples/GetPartnerDetailsNoPartnerId.json
if __name__ == "__main__":
    main()
