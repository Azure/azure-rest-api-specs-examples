from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python cloud_service_role_instance_get_remote_desktop_file.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.cloud_service_role_instances.get_remote_desktop_file(
        role_instance_name="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        resource_group_name="rgcloudService",
        cloud_service_name="aaaa",
    )
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Get_RemoteDesktopFile.json
if __name__ == "__main__":
    main()
