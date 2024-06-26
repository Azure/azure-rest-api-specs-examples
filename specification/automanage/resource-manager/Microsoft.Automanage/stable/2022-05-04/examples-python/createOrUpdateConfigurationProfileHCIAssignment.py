from azure.identity import DefaultAzureCredential
from azure.mgmt.automanage import AutomanageClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automanage
# USAGE
    python create_or_update_configuration_profile_hci_assignment.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomanageClient(
        credential=DefaultAzureCredential(),
        subscription_id="mySubscriptionId",
    )

    response = client.configuration_profile_hci_assignments.create_or_update(
        resource_group_name="myResourceGroupName",
        cluster_name="myClusterName",
        configuration_profile_assignment_name="default",
        parameters={
            "properties": {
                "configurationProfile": "/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"
            }
        },
    )
    print(response)


# x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileHCIAssignment.json
if __name__ == "__main__":
    main()
