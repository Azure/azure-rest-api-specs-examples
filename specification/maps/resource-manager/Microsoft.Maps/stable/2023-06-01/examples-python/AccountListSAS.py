from azure.identity import DefaultAzureCredential
from azure.mgmt.maps import AzureMapsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-maps
# USAGE
    python account_list_sas.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureMapsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="21a9967a-e8a9-4656-a70b-96ff1c4d05a0",
    )

    response = client.accounts.list_sas(
        resource_group_name="myResourceGroup",
        account_name="myMapsAccount",
        maps_account_sas_parameters={
            "expiry": "2017-05-24T11:42:03.1567373Z",
            "maxRatePerSecond": 500,
            "principalId": "e917f87b-324d-4728-98ed-e31d311a7d65",
            "regions": ["eastus"],
            "signingKey": "primaryKey",
            "start": "2017-05-24T10:42:03.1567373Z",
        },
    )
    print(response)


# x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/AccountListSAS.json
if __name__ == "__main__":
    main()
