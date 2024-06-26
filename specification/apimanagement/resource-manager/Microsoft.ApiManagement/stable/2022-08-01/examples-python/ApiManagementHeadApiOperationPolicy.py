from typing import TYPE_CHECKING, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.apimanagement import ApiManagementClient

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models
"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apimanagement
# USAGE
    python api_management_head_api_operation_policy.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.api_operation_policy.get_entity_tag(
        resource_group_name="rg1",
        service_name="apimService1",
        api_id="5600b539c53f5b0062040001",
        operation_id="5600b53ac53f5b0062080006",
        policy_id="policy",
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiOperationPolicy.json
if __name__ == "__main__":
    main()
