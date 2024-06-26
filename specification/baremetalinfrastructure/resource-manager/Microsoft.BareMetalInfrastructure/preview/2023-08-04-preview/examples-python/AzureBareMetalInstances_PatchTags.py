from azure.identity import DefaultAzureCredential
from azure.mgmt.baremetalinfrastructure import BareMetalInfrastructureClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-baremetalinfrastructure
# USAGE
    python azure_bare_metal_instances_patch_tags.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BareMetalInfrastructureClient(
        credential=DefaultAzureCredential(),
        subscription_id="f0f4887f-d13c-4943-a8ba-d7da28d2a3fd",
    )

    response = client.azure_bare_metal_instances.update(
        resource_group_name="myResourceGroup",
        azure_bare_metal_instance_name="myABMInstance",
        tags_parameter={"tags": {"testkey": "testvalue"}},
    )
    print(response)


# x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_PatchTags.json
if __name__ == "__main__":
    main()
