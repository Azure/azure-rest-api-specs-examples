from azure.identity import DefaultAzureCredential
from azure.mgmt.managementpartner import ACEProvisioningManagementPartnerAPI

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managementpartner
# USAGE
    python patch_partner_details.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ACEProvisioningManagementPartnerAPI(
        credential=DefaultAzureCredential(),
    )

    response = client.partner.update(
        partner_id="123456",
    )
    print(response)


# x-ms-original-file: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/PatchPartnerDetails.json
if __name__ == "__main__":
    main()
