from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.devcenter import DevCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devcenter
# USAGE
    python check_scoped_name_availability_dev_center_catalog.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevCenterMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="0ac520ee-14c0-480f-b6c9-0a90c58ffff",
    )

    response = client.check_scoped_name_availability.execute(
        name_availability_request={
            "name": "name1",
            "scope": "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso",
            "type": "Microsoft.DevCenter/devcenters/catalogs",
        },
    )
    print(response)


# x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/CheckScopedNameAvailability_DevCenterCatalog.json
if __name__ == "__main__":
    main()
