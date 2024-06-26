from typing import Any, IO, TYPE_CHECKING, Union

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
    python api_management_update_tenant_access.py

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

    response = client.tenant_access.update(
        resource_group_name="rg1",
        service_name="apimService1",
        access_name="access",
        if_match="*",
        parameters={"properties": {"enabled": True}},
    )
    print(response)


# x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateTenantAccess.json
if __name__ == "__main__":
    main()
