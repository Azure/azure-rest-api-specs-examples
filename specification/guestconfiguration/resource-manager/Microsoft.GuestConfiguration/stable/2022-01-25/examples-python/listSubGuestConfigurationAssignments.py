from azure.identity import DefaultAzureCredential
from azure.mgmt.guestconfig import GuestConfigurationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-guestconfig
# USAGE
    python list_all_guest_configuration_assignments_for_a_subscription.py

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

    response = client.guest_configuration_assignments.subscription_list()
    for item in response:
        print(item)


# x-ms-original-file: specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listSubGuestConfigurationAssignments.json
if __name__ == "__main__":
    main()
