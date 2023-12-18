from azure.identity import DefaultAzureCredential
from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python dedicated_host_update_resize.py

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

    response = client.dedicated_hosts.begin_update(
        resource_group_name="rgcompute",
        host_group_name="aaaaaaaaa",
        host_name="aaaaaaaaaaaaaaaaaaaaa",
        parameters={"sku": {"name": "DSv3-Type1"}},
    ).result()
    print(response)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/dedicatedHostExamples/DedicatedHost_Update_Resize.json
if __name__ == "__main__":
    main()
