from azure.identity import DefaultAzureCredential

from azure.mgmt.search import SearchManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-search
# USAGE
    python search_update_service_with_data_exfiltration.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SearchManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.services.update(
        resource_group_name="rg1",
        search_service_name="mysearchservice",
        service={
            "properties": {"dataExfiltrationProtections": ["BlockAll"], "replicaCount": 2},
            "tags": {"app-name": "My e-commerce app", "new-tag": "Adding a new tag"},
        },
    )
    print(response)


# x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/SearchUpdateServiceWithDataExfiltration.json
if __name__ == "__main__":
    main()
