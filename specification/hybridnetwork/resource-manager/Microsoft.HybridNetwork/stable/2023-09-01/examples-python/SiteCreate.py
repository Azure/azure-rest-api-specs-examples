from azure.identity import DefaultAzureCredential
from azure.mgmt.hybridnetwork import HybridNetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridnetwork
# USAGE
    python site_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridNetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.sites.begin_create_or_update(
        resource_group_name="rg1",
        site_name="testSite",
        parameters={
            "location": "westUs2",
            "properties": {
                "nfvis": [
                    {"location": "westUs2", "name": "nfvi1", "nfviType": "AzureCore"},
                    {
                        "customLocationReference": {
                            "id": "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation1"
                        },
                        "name": "nfvi2",
                        "nfviType": "AzureArcKubernetes",
                    },
                    {
                        "customLocationReference": {
                            "id": "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation2"
                        },
                        "name": "nfvi3",
                        "nfviType": "AzureOperatorNexus",
                    },
                ]
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteCreate.json
if __name__ == "__main__":
    main()
