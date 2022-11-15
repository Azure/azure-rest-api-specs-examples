from azure.identity import DefaultAzureCredential
from azure.mgmt.maintenance import MaintenanceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-maintenance
# USAGE
    python public_maintenance_configurations_list.py

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

    response = client.public_maintenance_configurations.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/examples/PublicMaintenanceConfigurations_List.json
if __name__ == "__main__":
    main()
