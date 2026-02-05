from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.changes import ChangesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-changes
# USAGE
    python list_changes.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ChangesClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId1",
    )

    response = client.changes.list(
        resource_group_name="resourceGroup1",
        resource_provider_namespace="resourceProvider1",
        resource_type="resourceType1",
        resource_name="resourceName1",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/changes/stable/2022-05-01/examples/ListChanges.json
if __name__ == "__main__":
    main()
