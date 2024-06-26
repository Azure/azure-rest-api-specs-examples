from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python create_or_update_a_partner.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.integration_account_partners.create_or_update(
        resource_group_name="testResourceGroup",
        integration_account_name="testIntegrationAccount",
        partner_name="testPartner",
        partner={
            "location": "westus",
            "properties": {
                "content": {"b2b": {"businessIdentities": [{"qualifier": "AA", "value": "ZZ"}]}},
                "metadata": {},
                "partnerType": "B2B",
            },
            "tags": {},
        },
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_CreateOrUpdate.json
if __name__ == "__main__":
    main()
