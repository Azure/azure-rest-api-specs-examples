from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python create_or_update_a_map_larger_than_4_mb.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="<Azure-subscription-ID>",
    )

    response = client.integration_account_maps.create_or_update(
        resource_group_name="testResourceGroup",
        integration_account_name="testIntegrationAccount",
        map_name="testMap",
        map={
            "location": "westus",
            "properties": {
                "contentLink": {"uri": "<blob-SAS-URL-for-map>"},
                "contentType": "application/xml",
                "mapType": "Xslt",
                "metadata": {},
            },
        },
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountLargeMaps_CreateOrUpdate.json
if __name__ == "__main__":
    main()
