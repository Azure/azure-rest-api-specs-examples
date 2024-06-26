from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python create_or_update_active_directory_connector.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureArcDataManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.active_directory_connectors.begin_create(
        resource_group_name="testrg",
        data_controller_name="testdataController",
        active_directory_connector_name="testADConnector",
        active_directory_connector_resource={
            "properties": {
                "spec": {
                    "activeDirectory": {
                        "domainControllers": {
                            "primaryDomainController": {"hostname": "dc1.contoso.local"},
                            "secondaryDomainControllers": [
                                {"hostname": "dc2.contoso.local"},
                                {"hostname": "dc3.contoso.local"},
                            ],
                        },
                        "realm": "CONTOSO.LOCAL",
                        "serviceAccountProvisioning": "manual",
                    },
                    "dns": {
                        "nameserverIPAddresses": ["11.11.111.111", "22.22.222.222"],
                        "preferK8sDnsForPtrLookups": False,
                        "replicas": 1,
                    },
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdateActiveDirectoryConnector.json
if __name__ == "__main__":
    main()
