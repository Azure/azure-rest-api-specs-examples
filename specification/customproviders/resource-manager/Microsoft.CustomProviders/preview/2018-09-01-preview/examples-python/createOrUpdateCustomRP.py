from azure.identity import DefaultAzureCredential
from azure.mgmt.customproviders import Customproviders

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-customproviders
# USAGE
    python create_or_update_the_custom_resource_provider.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = Customproviders(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.custom_resource_provider.begin_create_or_update(
        resource_group_name="testRG",
        resource_provider_name="newrp",
        resource_provider={
            "location": "eastus",
            "properties": {
                "actions": [{"endpoint": "https://mytestendpoint/", "name": "TestAction", "routingType": "Proxy"}],
                "resourceTypes": [
                    {"endpoint": "https://mytestendpoint2/", "name": "TestResource", "routingType": "Proxy,Cache"}
                ],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/createOrUpdateCustomRP.json
if __name__ == "__main__":
    main()
