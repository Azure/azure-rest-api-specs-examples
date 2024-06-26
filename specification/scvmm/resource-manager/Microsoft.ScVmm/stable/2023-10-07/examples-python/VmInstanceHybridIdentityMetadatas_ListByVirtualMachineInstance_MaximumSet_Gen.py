from azure.identity import DefaultAzureCredential

from azure.mgmt.scvmm import ScVmmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python vm_instance_hybrid_identity_metadatas_list_by_virtual_machine_instance_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ScVmmMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.vm_instance_hybrid_identity_metadatas.list_by_virtual_machine_instance(
        resource_uri="gtgclehcbsyave",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VmInstanceHybridIdentityMetadatas_ListByVirtualMachineInstance_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
