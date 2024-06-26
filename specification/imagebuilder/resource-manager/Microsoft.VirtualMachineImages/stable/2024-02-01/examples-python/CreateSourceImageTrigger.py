from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.imagebuilder import ImageBuilderClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-imagebuilder
# USAGE
    python create_source_image_trigger.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ImageBuilderClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.triggers.begin_create_or_update(
        resource_group_name="myResourceGroup",
        image_template_name="myImageTemplate",
        trigger_name="source",
        parameters={"properties": {"kind": "SourceImage"}},
    ).result()
    print(response)


# x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/CreateSourceImageTrigger.json
if __name__ == "__main__":
    main()
