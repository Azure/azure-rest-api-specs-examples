from azure.identity import DefaultAzureCredential

from azure.mgmt.fabric import FabricMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-fabric
# USAGE
    python fabric_capacities_list_skus.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FabricMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.fabric_capacities.list_skus()
    for item in response:
        print(item)


# x-ms-original-file: 2023-11-01/FabricCapacities_ListSkus.json
if __name__ == "__main__":
    main()
