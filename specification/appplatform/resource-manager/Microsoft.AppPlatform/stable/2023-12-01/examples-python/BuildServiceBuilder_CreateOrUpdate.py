from azure.identity import DefaultAzureCredential

from azure.mgmt.appplatform import AppPlatformManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appplatform
# USAGE
    python build_service_builder_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AppPlatformManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.build_service_builder.begin_create_or_update(
        resource_group_name="myResourceGroup",
        service_name="myservice",
        build_service_name="default",
        builder_name="mybuilder",
        builder_resource={
            "properties": {
                "buildpackGroups": [{"buildpacks": [{"id": "tanzu-buildpacks/java-azure"}], "name": "mix"}],
                "stack": {"id": "io.buildpacks.stacks.bionic", "version": "base"},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/BuildServiceBuilder_CreateOrUpdate.json
if __name__ == "__main__":
    main()
