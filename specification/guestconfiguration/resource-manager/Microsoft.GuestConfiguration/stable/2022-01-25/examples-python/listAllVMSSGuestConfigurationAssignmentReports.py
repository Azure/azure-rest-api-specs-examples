from azure.identity import DefaultAzureCredential
from azure.mgmt.guestconfig import GuestConfigurationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-guestconfig
# USAGE
    python list_all_reports_for_the_vmss_guest_configuration_assignment_with_latest_report_first.py

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

    response = client.guest_configuration_assignment_reports_vmss.list(
        resource_group_name="myResourceGroupName",
        vmss_name="myVMSSName",
        name="AuditSecureProtocol",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllVMSSGuestConfigurationAssignmentReports.json
if __name__ == "__main__":
    main()
