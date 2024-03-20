from azure.identity import DefaultAzureCredential
from azure.mgmt.servicelinker import ServiceLinkerManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicelinker
# USAGE
    python patch_dryrun.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceLinkerManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.linkers.begin_update_dryrun(
        resource_uri="subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
        dryrun_name="dryrunName",
        parameters={
            "properties": {
                "parameters": {
                    "actionName": "createOrUpdate",
                    "authInfo": {
                        "authType": "secret",
                        "name": "name",
                        "secretInfo": {"secretType": "rawValue", "value": "secret"},
                    },
                    "targetService": {
                        "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
                        "type": "AzureResource",
                    },
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2023-04-01-preview/examples/PatchDryrun.json
if __name__ == "__main__":
    main()
