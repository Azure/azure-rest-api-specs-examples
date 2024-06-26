from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.support import MicrosoftSupport

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-support
# USAGE
    python update_advanced_diagnostic_consent_of_support_ticket_for_subscription.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftSupport(
        credential=DefaultAzureCredential(),
        subscription_id="132d901f-189d-4381-9214-fe68e27e05a1",
    )

    response = client.support_tickets.update(
        support_ticket_name="testticket",
        update_support_ticket={"advancedDiagnosticConsent": "Yes"},
    )
    print(response)


# x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/UpdateAdvancedDiagnosticConsentOfSupportTicketForSubscription.json
if __name__ == "__main__":
    main()
