from azure.identity import DefaultAzureCredential

from azure.mgmt.authorization import AuthorizationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-authorization
# USAGE
    python validate_role_assignment_schedule_request_by_name.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AuthorizationManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.role_assignment_schedule_requests.validate(
        scope="subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
        role_assignment_schedule_request_name="fea7a502-9a96-4806-a26f-eee560e52045",
        parameters={
            "properties": {
                "condition": "@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'",
                "conditionVersion": "1.0",
                "linkedRoleEligibilityScheduleId": "b1477448-2cc6-4ceb-93b4-54a202a89413",
                "principalId": "a3bb8764-cb92-4276-9d2a-ca1e895e55ea",
                "requestType": "SelfActivate",
                "roleDefinitionId": "/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608",
                "scheduleInfo": {
                    "expiration": {"duration": "PT8H", "endDateTime": None, "type": "AfterDuration"},
                    "startDateTime": "2020-09-09T21:35:27.91Z",
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2024-09-01-preview/examples/ValidateRoleAssignmentScheduleRequestByName.json
if __name__ == "__main__":
    main()
