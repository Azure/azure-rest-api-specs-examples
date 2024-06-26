from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.support import MicrosoftSupport

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-support
# USAGE
    python create_machine_learning_quota_ticket_for_low_priority_cores.py

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

    response = client.support_tickets.begin_create(
        support_ticket_name="testticket",
        create_support_ticket_parameters={
            "properties": {
                "advancedDiagnosticConsent": "Yes",
                "contactDetails": {
                    "country": "usa",
                    "firstName": "abc",
                    "lastName": "xyz",
                    "preferredContactMethod": "email",
                    "preferredSupportLanguage": "en-US",
                    "preferredTimeZone": "Pacific Standard Time",
                    "primaryEmailAddress": "abc@contoso.com",
                },
                "description": "my description",
                "problemClassificationId": "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/machine_learning_service_problemClassification_guid",
                "quotaTicketDetails": {
                    "quotaChangeRequestSubType": "BatchAml",
                    "quotaChangeRequestVersion": "1.0",
                    "quotaChangeRequests": [{"payload": '{"NewLimit":200,"Type":"LowPriority"}', "region": "EastUS"}],
                },
                "serviceId": "/providers/Microsoft.Support/services/quota_service_guid",
                "severity": "moderate",
                "supportPlanId": "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=",
                "title": "my title",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateMachineLearningQuotaTicketForLowPriorityCores.json
if __name__ == "__main__":
    main()
