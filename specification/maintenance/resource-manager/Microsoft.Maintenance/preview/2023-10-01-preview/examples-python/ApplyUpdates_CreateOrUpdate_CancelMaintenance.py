from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.maintenance import MaintenanceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-maintenance
# USAGE
    python apply_updates_create_or_update_cancel_maintenance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MaintenanceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="5b4b650e-28b9-4790-b3ab-ddbd88d727c4",
    )

    response = client.apply_updates.create_or_update_or_cancel(
        resource_group_name="examplerg",
        provider_name="Microsoft.Maintenance",
        resource_type="maintenanceConfigurations",
        resource_name="maintenanceConfig1",
        apply_update_name="20230901121200",
        apply_update={"properties": {"status": "Cancel"}},
    )
    print(response)


# x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ApplyUpdates_CreateOrUpdate_CancelMaintenance.json
if __name__ == "__main__":
    main()
