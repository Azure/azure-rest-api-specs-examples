from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.managedapplications import ApplicationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python create_or_update_application.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApplicationClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.applications.begin_create_or_update(
        resource_group_name="rg",
        application_name="myManagedApplication",
        parameters={
            "kind": "ServiceCatalog",
            "location": "East US 2",
            "properties": {
                "applicationDefinitionId": "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
                "managedResourceGroupId": "/subscriptions/subid/resourceGroups/myManagedRG",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2019-07-01/examples/createOrUpdateApplication.json
if __name__ == "__main__":
    main()
