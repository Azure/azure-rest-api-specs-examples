from azure.identity import DefaultAzureCredential
from azure.mgmt.guestconfig import GuestConfigurationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-guestconfig
# USAGE
    python get_a_guest_configuration_assignment_report_by_id_for_a_virtual_machine.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = GuestConfigurationClient(
        credential=DefaultAzureCredential(),
        subscription_id="mySubscriptionid",
    )

    response = client.guest_configuration_assignment_reports.get(
        resource_group_name="myResourceGroupName",
        guest_configuration_assignment_name="AuditSecureProtocol",
        report_id="7367cbb8-ae99-47d0-a33b-a283564d2cb1",
        vm_name="myvm",
    )
    print(response)


# x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationAssignmentReportById.json
if __name__ == "__main__":
    main()
