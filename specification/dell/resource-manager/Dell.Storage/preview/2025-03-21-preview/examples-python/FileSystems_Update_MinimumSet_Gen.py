from azure.identity import DefaultAzureCredential

from azure.mgmt.dellstorage import DellStorageMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dellstorage
# USAGE
    python file_systems_update_minimum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DellStorageMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.file_systems.update(
        resource_group_name="rgDell",
        filesystem_name="abcd",
        properties={"properties": {"capacity": {"current": "5"}, "delegatedSubnetId": "uqfvajvyltgmqvdnxhbrfqbpuey"}},
    )
    print(response)


# x-ms-original-file: 2025-03-21-preview/FileSystems_Update_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
