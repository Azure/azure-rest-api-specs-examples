from azure.identity import DefaultAzureCredential
from azure.mgmt.guestconfig import GuestConfigurationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-guestconfig
# USAGE
    python create_or_update_guest_configuration_assignment.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = GuestConfigurationClient(
        credential=DefaultAzureCredential(),
        subscription_id="mySubscriptionId",
    )

    response = client.guest_configuration_assignments.create_or_update(
        guest_configuration_assignment_name="NotInstalledApplicationForWindows",
        resource_group_name="myResourceGroupName",
        vm_name="myVMName",
        parameters={
            "location": "westcentralus",
            "name": "NotInstalledApplicationForWindows",
            "properties": {
                "context": "Azure policy",
                "guestConfiguration": {
                    "assignmentType": "ApplyAndAutoCorrect",
                    "configurationParameter": [
                        {"name": "[InstalledApplication]NotInstalledApplicationResource1;Name", "value": "NotePad,sql"}
                    ],
                    "contentHash": "123contenthash",
                    "contentUri": "https://thisisfake/pacakge",
                    "name": "NotInstalledApplicationForWindows",
                    "version": "1.*",
                },
            },
        },
    )
    print(response)


# x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/createOrUpdateGuestConfigurationAssignment.json
if __name__ == "__main__":
    main()
