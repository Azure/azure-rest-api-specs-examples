from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.maintenance import MaintenanceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-maintenance
# USAGE
    python configuration_assignments_for_resource_group_update_for_resource.py

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

    response = client.configuration_assignments_for_resource_group.update(
        resource_group_name="examplerg",
        configuration_assignment_name="workervmConfiguration",
        configuration_assignment={
            "properties": {
                "filter": {
                    "locations": ["Japan East", "UK South"],
                    "resourceTypes": ["Microsoft.HybridCompute/machines", "Microsoft.Compute/virtualMachines"],
                    "tagSettings": {
                        "filterOperator": "Any",
                        "tags": {
                            "tag1": ["tag1Value1", "tag1Value2", "tag1Value3"],
                            "tag2": ["tag2Value1", "tag2Value2", "tag2Value3"],
                        },
                    },
                },
                "maintenanceConfigurationId": "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/ConfigurationAssignmentsForResourceGroup_UpdateForResource.json
if __name__ == "__main__":
    main()
